package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "accounts"
)

type Users struct {
	Username      string
	Password      string
	Approved      bool
	Denied        bool
	Pending       bool
	Notapplied    bool
	Nametooshort  bool
	Namenotunique bool
	Pwtooshort    bool
}
type Applications struct {
	Username  string
	Firstname string
	Lastname  string
	Address   string
	Phone     string
}
type Accountholders struct {
	Username      string
	Firstname     string
	Lastname      string
	Address       string
	Phone         string
	Accountnumber int
	Checking      int
	Savings       int
}
type ViewInfo struct {
	Ap            []Applications
	Usr           []Users
	Ac            []Accountholders
	Singleuser    Users
	Singleaccount Accountholders
	Singleapp     Applications
	Insufficient  bool
	Login         LoginInfo
}
type LoginInfo struct {
	CurrentUser   string
	Employee      bool
	Loggedin      bool
	Invalidname   bool
	Invalidpw     bool
	Notauthorized bool
}

var Signin = LoginInfo{}

func index(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("templates/index.html")
	temp.Execute(response, Signin)
}
func registrationForm(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("templates/registrationform.html")
	temp.Execute(response, nil)
}
func register(response http.ResponseWriter, request *http.Request) {
	db := connect()
	temp, _ := template.ParseFiles("templates/register.html")
	user := Users{}
	user.Username = request.FormValue("name")
	user.Password = request.FormValue("pw")
	if len(user.Username) < 3 {
		user.Nametooshort = true

	} else if uniqueName(db, user.Username) == false {
		user.Namenotunique = true
	} else if len(user.Password) < 3 {
		user.Pwtooshort = true
	} else {
		statement := "INSERT INTO users (username, password, status)"
		statement += " VALUES ($1, $2, $3);"
		_, err := db.Exec(statement, user.Username, user.Password, "notapplied")
		if err != nil {
			panic(err)
		}
	}
	defer db.Close()
	temp.Execute(response, user)
}
func login(response http.ResponseWriter, request *http.Request) {
	db := connect()
	temp, _ := template.ParseFiles("templates/login.html")
	var status string
	ac := Accountholders{}
	user := Users{}
	view := ViewInfo{}
	login := LoginInfo{}
	user.Approved = false
	user.Denied = false
	user.Pending = false
	user.Notapplied = false
	if !Signin.Loggedin {
		user.Username = request.FormValue("name")
		user.Password = request.FormValue("pw")
		if uniqueName(db, user.Username) == true {
			login.Invalidname = true
		} else if passwordMatches(db, user.Username, user.Password) == false {
			login.Invalidpw = true
		} else {
			Signin.CurrentUser = user.Username
			Signin.Loggedin = true
			Signin.Employee = false
		}
	} else {

		user.Username = Signin.CurrentUser
	}
	status = getStatus(db, user.Username)
	if status == "notapplied" {
		user.Notapplied = true
	} else if status == "approved" {
		user.Approved = true
	} else if status == "denied" {
		user.Denied = true
	} else {
		user.Pending = true
	}
	defer db.Close()
	view.Singleuser = user
	ac.Checking, ac.Savings = getBalance(db, user.Username)
	view.Singleaccount = ac
	view.Login = login
	temp.Execute(response, view)
}
func deposit(response http.ResponseWriter, request *http.Request) {
	db := connect()
	var current int
	var query, statement, status string
	ac := Accountholders{}
	user := Users{}
	view := ViewInfo{}
	user.Username = Signin.CurrentUser
	user.Approved = false
	user.Denied = false
	user.Pending = false
	user.Notapplied = false
	temp, _ := template.ParseFiles("templates/login.html")
	amount, _ := strconv.Atoi(request.FormValue("amount"))
	choice := request.FormValue("account")
	if choice == "checking" {
		query = "SELECT checking FROM accountholders WHERE username=$1;"
	} else {
		query = "SELECT savings FROM accountholders WHERE username=$1;"
	}
	row := db.QueryRow(query, Signin.CurrentUser)
	row.Scan(&current)
	amount += current
	if choice == "checking" {
		statement = "UPDATE accountholders SET checking=$1 WHERE username=$2;"
	} else {
		statement = "UPDATE accountholders SET savings=$1 WHERE username=$2;"
	}
	_, err := db.Exec(statement, amount, Signin.CurrentUser)
	if err != nil {
		panic(err)
	}
	status = getStatus(db, user.Username)
	if status == "notapplied" {
		user.Notapplied = true
	} else if status == "approved" {
		user.Approved = true
	} else if status == "denied" {
		user.Denied = true
	} else {
		user.Pending = true
	}
	defer db.Close()
	view.Singleuser = user
	ac.Checking, ac.Savings = getBalance(db, user.Username)
	view.Singleaccount = ac
	temp.Execute(response, view)
}
func withdraw(response http.ResponseWriter, request *http.Request) {
	db := connect()
	var current int
	var query, statement, status string
	ac := Accountholders{}
	user := Users{}
	view := ViewInfo{}
	user.Username = Signin.CurrentUser
	user.Approved = false
	user.Denied = false
	user.Pending = false
	user.Notapplied = false
	temp, _ := template.ParseFiles("templates/login.html")
	amount, _ := strconv.Atoi(request.FormValue("amount"))
	choice := request.FormValue("account")
	if choice == "checking" {
		query = "SELECT checking FROM accountholders WHERE username=$1;"
	} else {
		query = "SELECT savings FROM accountholders WHERE username=$1;"
	}
	row := db.QueryRow(query, Signin.CurrentUser)
	row.Scan(&current)
	current -= amount
	if choice == "checking" {
		statement = "UPDATE accountholders SET checking=$1 WHERE username=$2;"
	} else {
		statement = "UPDATE accountholders SET savings=$1 WHERE username=$2;"
	}
	if current >= 0 {
		_, err := db.Exec(statement, current, Signin.CurrentUser)
		if err != nil {
			panic(err)
		}
	}
	status = getStatus(db, user.Username)
	if status == "notapplied" {
		user.Notapplied = true
	} else if status == "approved" {
		user.Approved = true
	} else if status == "denied" {
		user.Denied = true
	} else {
		user.Pending = true
	}
	defer db.Close()
	view.Singleuser = user
	view.Insufficient = false
	ac.Checking, ac.Savings = getBalance(db, user.Username)
	view.Singleaccount = ac
	if current < 0 {
		view.Insufficient = true
	}
	temp.Execute(response, view)
}
func transfer(response http.ResponseWriter, request *http.Request) {
	db := connect()
	var fromAmount, toAmount int
	var sameAccount bool = false
	var query, statement1, statement2, status string
	ac := Accountholders{}
	user := Users{}
	view := ViewInfo{}
	user.Username = Signin.CurrentUser
	user.Approved = false
	user.Denied = false
	user.Pending = false
	user.Notapplied = false
	temp, _ := template.ParseFiles("templates/login.html")
	transferAmount, _ := strconv.Atoi(request.FormValue("amount"))
	fromAccount := request.FormValue("fromaccount")
	toAccount := request.FormValue("toaccount")
	// query amount in account transferring from
	if fromAccount == "checking" {
		query = "SELECT checking FROM accountholders WHERE username=$1;"
	} else {
		query = "SELECT savings FROM accountholders WHERE username=$1;"
	}
	row := db.QueryRow(query, Signin.CurrentUser)
	row.Scan(&fromAmount)
	// query amount in account transferring to
	if toAccount == "checking" {
		query = "SELECT checking FROM accountholders WHERE username=$1;"
	} else {
		query = "SELECT savings FROM accountholders WHERE username=$1;"
	}
	row = db.QueryRow(query, Signin.CurrentUser)
	row.Scan(&toAmount)
	// create statements to transfer money
	if fromAccount == "checking" {
		statement1 = "UPDATE accountholders SET checking=$1 WHERE username=$2;"
		statement2 = "UPDATE accountholders SET savings=$1 WHERE username=$2;"
	} else {
		statement1 = "UPDATE accountholders SET savings=$1 WHERE username=$2;"
		statement2 = "UPDATE accountholders SET checking=$1 WHERE username=$2;"
	}
	fromAmount -= transferAmount
	toAmount += transferAmount
	// perform transaction if sufficient funds
	if fromAccount == toAccount {
		sameAccount = true
	}
	if fromAmount >= 0 && !sameAccount {
		_, err := db.Exec(statement1, fromAmount, Signin.CurrentUser)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(statement2, toAmount, Signin.CurrentUser)
		if err != nil {
			panic(err)
		}
	}
	status = getStatus(db, user.Username)
	if status == "notapplied" {
		user.Notapplied = true
	} else if status == "approved" {
		user.Approved = true
	} else if status == "denied" {
		user.Denied = true
	} else {
		user.Pending = true
	}
	defer db.Close()
	view.Singleuser = user
	view.Insufficient = false
	ac.Checking, ac.Savings = getBalance(db, user.Username)
	view.Singleaccount = ac
	if fromAmount < 0 {
		view.Insufficient = true
	}
	temp.Execute(response, view)
}
func employeelogin(response http.ResponseWriter, request *http.Request) {
	db := connect()
	temp, _ := template.ParseFiles("templates/employeelogin.html")
	user := Users{}
	login := LoginInfo{}
	view := ViewInfo{}
	if !Signin.Employee {
		user.Username = request.FormValue("name")
		user.Password = request.FormValue("pw")
		if len(user.Username) == 0 {
			login.Notauthorized = true
		} else if uniqueEmployeeName(db, user.Username) {
			login.Invalidname = true
		} else if employeePasswordMatches(db, user.Username, user.Password) == false {
			login.Invalidpw = true
		} else {
			Signin.CurrentUser = user.Username
			Signin.Employee = true
		}
	}
	view.Login = login
	if Signin.Employee {
		rows, _ := db.Query("SELECT * FROM applications;")
		for rows.Next() {
			var username, firstname, lastname, address, phone string
			var ap = Applications{}
			rows.Scan(&username, &firstname, &lastname, &address, &phone)
			ap.Username = username
			ap.Firstname = firstname
			ap.Lastname = lastname
			ap.Address = address
			ap.Phone = phone
			view.Ap = append(view.Ap, ap)
		}
		rows, _ = db.Query("SELECT * FROM accountholders")
		for rows.Next() {
			var username, fn, ln, ph, add string
			var accountnumber, checking, savings int
			var ac = Accountholders{}
			rows.Scan(&username, &accountnumber, &fn, &ln, &add, &ph, &checking, &savings)
			ac.Username = username
			ac.Firstname = fn
			ac.Lastname = ln
			ac.Address = add
			ac.Phone = ph
			ac.Accountnumber = accountnumber
			ac.Checking = checking
			ac.Savings = savings
			view.Ac = append(view.Ac, ac)
		}
	}
	defer db.Close()
	temp.Execute(response, view)
}
func process(response http.ResponseWriter, request *http.Request) {
	db := connect()
	temp, _ := template.ParseFiles("templates/employeelogin.html")
	var statement, query, choice, action, fn, ln, add, ph string
	choice = request.FormValue("choice")
	action = request.FormValue("action")
	query = "SELECT firstname,lastname,address,phone FROM applications WHERE username = $1;"
	row := db.QueryRow(query, choice)
	row.Scan(&fn, &ln, &add, &ph)
	statement = "DELETE FROM applications WHERE username = $1;"
	_, err := db.Exec(statement, choice)
	if err != nil {
		panic(err)
	}
	if action == "approve" {
		statement = "INSERT INTO accountholders (username,firstname,lastname,address,phone,"
		statement += "checking,savings)"
		statement += " VALUES ($1, $2, $3, $4, $5, $6, $7);"
		_, err = db.Exec(statement, choice, fn, ln, add, ph, 0, 0)
		if err != nil {
			panic(err)
		}
		statement = "UPDATE users SET status=$1 WHERE username=$2;"
		_, err = db.Exec(statement, "approved", choice)
		if err != nil {
			panic(err)
		}
	} else {
		statement = "UPDATE users SET status=$1 WHERE username=$2;"
		_, err = db.Exec(statement, "denied", choice)
		if err != nil {
			panic(err)
		}
	}
	view := ViewInfo{}
	rows, _ := db.Query("SELECT * FROM applications")
	for rows.Next() {
		var username, firstname, lastname, address, phone string
		var ap = Applications{}
		rows.Scan(&username, &firstname, &lastname, &address, &phone)
		ap.Username = username
		ap.Firstname = firstname
		ap.Lastname = lastname
		ap.Address = address
		ap.Phone = phone
		view.Ap = append(view.Ap, ap)
	}
	rows, _ = db.Query("SELECT * FROM accountholders")
	for rows.Next() {
		var username, fn, ln, ph, add string
		var accountnumber, checking, savings int
		var ac = Accountholders{}
		rows.Scan(&username, &accountnumber, &fn, &ln, &add, &ph, &checking, &savings)
		ac.Username = username
		ac.Firstname = fn
		ac.Lastname = ln
		ac.Address = add
		ac.Phone = ph
		ac.Accountnumber = accountnumber
		ac.Checking = checking
		ac.Savings = savings
		view.Ac = append(view.Ac, ac)
	}
	defer db.Close()
	temp.Execute(response, view)
}
func apply(response http.ResponseWriter, request *http.Request) {
	db := connect()
	var statement string
	temp, _ := template.ParseFiles("templates/apply.html")
	statement = "UPDATE users SET status=$1 WHERE username=$2;"
	_, err := db.Exec(statement, "pending", Signin.CurrentUser)
	if err != nil {
		panic(err)
	}
	ap := Applications{}
	ap.Firstname = request.FormValue("first")
	ap.Lastname = request.FormValue("last")
	ap.Address = request.FormValue("address")
	ap.Phone = request.FormValue("phone")
	statement = "INSERT INTO applications (username, firstname, lastname, address, phone)"
	statement += " VALUES ($1, $2, $3, $4, $5);"
	_, err = db.Exec(statement, Signin.CurrentUser, ap.Firstname, ap.Lastname, ap.Address, ap.Phone)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	temp.Execute(response, ap)
}
func logout(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("templates/index.html")
	Signin.Loggedin = false
	Signin.Employee = false
	temp.Execute(response, Signin)
}
func uniqueName(db *sql.DB, name string) bool {
	rows, _ := db.Query("SELECT username FROM users")
	for rows.Next() {
		var username string
		rows.Scan(&username)
		if name == username {
			return false
		}
	}
	return true
}
func passwordMatches(db *sql.DB, name string, password string) bool {
	var pw string
	row := db.QueryRow("SELECT password FROM users WHERE username = $1", name)
	row.Scan(&pw)
	if password == pw {
		return true
	}
	return false
}
func uniqueEmployeeName(db *sql.DB, name string) bool {
	rows, _ := db.Query("SELECT username FROM employees")
	for rows.Next() {
		var username string
		rows.Scan(&username)
		if name == username {
			return false
		}
	}
	return true
}
func employeePasswordMatches(db *sql.DB, name string, password string) bool {
	var pw string
	row := db.QueryRow("SELECT password FROM employees WHERE username = $1", name)
	row.Scan(&pw)
	if password == pw {
		return true
	}
	return false
}

func getStatus(db *sql.DB, name string) string {
	var status string
	row := db.QueryRow("SELECT status FROM users WHERE username = $1", name)
	row.Scan(&status)
	return status
}
func getBalance(db *sql.DB, name string) (int, int) {
	var checking, savings int
	row := db.QueryRow("SELECT checking FROM accountholders where username = $1", name)
	row.Scan(&checking)
	row = db.QueryRow("SELECT savings FROM accountholders where username = $1", name)
	row.Scan(&savings)
	return checking, savings
}
func connect() *sql.DB {
	var conn string
	conn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected to database")
	return db
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/registrationform", registrationForm)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/apply", apply)
	http.HandleFunc("/employeelogin", employeelogin)
	http.HandleFunc("/process", process)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
	http.HandleFunc("/logout", logout)
	Signin.Loggedin = false
	Signin.Employee = false
	http.ListenAndServe(":7000", nil)
}
