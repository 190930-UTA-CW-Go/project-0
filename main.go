package main

import (
	"fmt"
	"html/template"
	"net/http"
	"database/sql"
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
	Username   string
	Password   string
	Approved   bool
	Denied     bool
	Pending    bool
	Notapplied bool
}
type Applications struct {
	Username  string
	Firstname string
	Lastname  string
	Address   string
	Phone     string
}
type Accountholders struct {
	Username string
	Checking int
	Savings  int
}
type ViewInfo struct {
	Ap            []Applications
	Usr           []Users
	Ac            []Accountholders
	Singleuser    Users
	Singleaccount Accountholders
	Singleapp     Applications
	Insufficient  bool
}
type LoginInfo struct {
	CurrentUser string
	Employee    bool
	Loggedin    bool
}

var Signin = LoginInfo{}

func index(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("templates/index.html")
	temp.Execute(response, Signin)
}
func register(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("templates/register.html")
	temp.Execute(response, nil)
}
func confirm(response http.ResponseWriter, request *http.Request) {
	db := connect()
	temp, _ := template.ParseFiles("templates/confirm.html")
	var query string
	user := Users{}
	user.Username = request.FormValue("name")
	user.Password = request.FormValue("pw")
	if uniqueName(db, user.Username) == false || len(user.Username) < 3 {
		db.Close()
		if len(user.Username) < 3 {
			temp, _ := template.ParseFiles("templates/nametooshort.html")
			temp.Execute(response, nil)

		} else {
			temp, _ = template.ParseFiles("templates/notunique.html")
			temp.Execute(response, nil)
		}
		return
	}
	if len(user.Password) < 3 {
		db.Close()
		temp, _ := template.ParseFiles("templates/pwtooshort.html")
		temp.Execute(response, nil)
		return
	}
	query = "INSERT INTO users (username, password, status)"
	query += " VALUES ($1, $2, $3)"
	db.QueryRow(query, user.Username, user.Password, "notapplied")
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
	user.Approved = false
	user.Denied = false
	user.Pending = false
	user.Notapplied = false
	if !Signin.Loggedin {
		user.Username = request.FormValue("name")
		user.Password = request.FormValue("pw")
		if uniqueName(db, user.Username) == true {
			db.Close()
			temp, _ := template.ParseFiles("templates/namenotfound.html")
			temp.Execute(response, nil)
			return
		}
		if passwordMatches(db, user.Username, user.Password) == false {
			db.Close()
			temp, _ := template.ParseFiles("templates/pwnotmatch.html")
			temp.Execute(response, nil)
			return
		}
		Signin.CurrentUser = user.Username
		Signin.Loggedin = true
		Signin.Employee = false
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
		query = "SELECT checking FROM accountholders WHERE username=$1"
	} else {
		query = "SELECT savings FROM accountholders WHERE username=$1"
	}
	row := db.QueryRow(query, Signin.CurrentUser)
	row.Scan(&current)
	amount += current
	if choice == "checking" {
		statement = "UPDATE accountholders SET checking=$1 WHERE username=$2"
	} else {
		statement = "UPDATE accountholders SET savings=$1 WHERE username=$2"
	}
	db.Exec(statement, amount, Signin.CurrentUser)
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
		query = "SELECT checking FROM accountholders WHERE username=$1"
	} else {
		query = "SELECT savings FROM accountholders WHERE username=$1"
	}
	row := db.QueryRow(query, Signin.CurrentUser)
	row.Scan(&current)
	current -= amount
	if choice == "checking" {
		statement = "UPDATE accountholders SET checking=$1 WHERE username=$2"
	} else {
		statement = "UPDATE accountholders SET savings=$1 WHERE username=$2"
	}
	if current >= 0 {
		db.Exec(statement, current, Signin.CurrentUser)
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
		query = "SELECT checking FROM accountholders WHERE username=$1"
	} else {
		query = "SELECT savings FROM accountholders WHERE username=$1"
	}
	row := db.QueryRow(query, Signin.CurrentUser)
	row.Scan(&fromAmount)
	// query amount in account transferring to
	if toAccount == "checking" {
		query = "SELECT checking FROM accountholders WHERE username=$1"
	} else {
		query = "SELECT savings FROM accountholders WHERE username=$1"
	}
	row = db.QueryRow(query, Signin.CurrentUser)
	row.Scan(&toAmount)
	// create statements to transfer money
	if fromAccount == "checking" {
		statement1 = "UPDATE accountholders SET checking=$1 WHERE username=$2"
		statement2 = "UPDATE accountholders SET savings=$1 WHERE username=$2"
	} else {
		statement1 = "UPDATE accountholders SET savings=$1 WHERE username=$2"
		statement2 = "UPDATE accountholders SET checking=$1 WHERE username=$2"
	}
	fromAmount -= transferAmount
	toAmount += transferAmount
	// perform transaction if sufficient funds
	if fromAccount == toAccount {
		sameAccount = true
	}
	if fromAmount >= 0 && !sameAccount {
		db.Exec(statement1, fromAmount, Signin.CurrentUser)
		db.Exec(statement2, toAmount, Signin.CurrentUser)
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
	user.Username = request.FormValue("name")
	user.Password = request.FormValue("pw")
	if !Signin.Employee {
		if user.Username == "" {
			db.Close()
			temp, _ := template.ParseFiles("templates/notauthorized.html")
			temp.Execute(response, nil)
			return
		}
		if uniqueEmployeeName(db, user.Username) == true {
			db.Close()
			temp, _ := template.ParseFiles("templates/employeenotfound.html")
			temp.Execute(response, nil)
			return
		}
		if employeePasswordMatches(db, user.Username, user.Password) == false {
			db.Close()
			temp, _ := template.ParseFiles("templates/employeepwnotmatch.html")
			temp.Execute(response, nil)
			return
		}
	}
	Signin.CurrentUser = user.Username
	Signin.Employee = true
	view := ViewInfo{}
	rows, _ := db.Query("select * from applications")
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
	defer db.Close()
	temp.Execute(response, view)
}
func process(response http.ResponseWriter, request *http.Request) {
	db := connect()
	temp, _ := template.ParseFiles("templates/employeelogin.html")
	var query, statement, choice, action string
	choice = request.FormValue("choice")
	action = request.FormValue("action")
	statement = "delete from applications where username = $1"
	db.Exec(statement, choice)
	if action == "approve" {
		query = "INSERT INTO accountholders (username, checking, savings)"
		query += " VALUES ($1, $2, $3)"
		db.QueryRow(query, choice, 0, 0)
		statement = "UPDATE users SET status=$1 WHERE username=$2"
		db.Exec(statement, "approved", choice)

	} else {
		statement = "UPDATE users SET status=$1 WHERE username=$2"
		db.Exec(statement, "denied", choice)
	}
	view := ViewInfo{}
	rows, _ := db.Query("select * from applications")
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
	defer db.Close()
	temp.Execute(response, view)
}
func viewAccounts(response http.ResponseWriter, request *http.Request) {
	db := connect()
	temp, _ := template.ParseFiles("templates/viewaccounts.html")
	if !Signin.Employee {
		db.Close()
		temp, _ := template.ParseFiles("templates/notauthorized.html")
		temp.Execute(response, nil)
		return
	}
	view := ViewInfo{}
	rows, _ := db.Query("select * from accountholders")
	for rows.Next() {
		var username string
		var checking, savings int
		var ac = Accountholders{}
		rows.Scan(&username, &checking, &savings)
		ac.Username = username
		ac.Checking = checking
		ac.Savings = savings
		view.Ac = append(view.Ac, ac)
	}
	defer db.Close()
	temp.Execute(response, view)

}
func apply(response http.ResponseWriter, request *http.Request) {
	db := connect()
	var query string
	temp, _ := template.ParseFiles("templates/apply.html")
	statement := "UPDATE users SET status=$1 WHERE username=$2"
	db.Exec(statement, "pending", Signin.CurrentUser)
	ap := Applications{}
	ap.Firstname = request.FormValue("first")
	ap.Lastname = request.FormValue("last")
	ap.Address = request.FormValue("address")
	ap.Phone = request.FormValue("phone")
	query = "INSERT INTO applications (username, firstname, lastname, address, phone)"
	query += " VALUES ($1, $2, $3, $4, $5)"
	db.QueryRow(query, Signin.CurrentUser, ap.Firstname, ap.Lastname, ap.Address, ap.Phone)
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
	rows, _ := db.Query("select username from users")
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
	row := db.QueryRow("select password from users where username = $1", name)
	row.Scan(&pw)
	if password == pw {
		return true
	}
	return false
}
func uniqueEmployeeName(db *sql.DB, name string) bool {
	rows, _ := db.Query("select username from employees")
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
	row := db.QueryRow("select password from employees where username = $1", name)
	row.Scan(&pw)
	if password == pw {
		return true
	}
	return false
}

func getStatus(db *sql.DB, name string) string {
	var status string
	row := db.QueryRow("select status from users where username = $1", name)
	row.Scan(&status)
	return status
}
func getBalance(db *sql.DB, name string) (int, int) {
	var checking, savings int
	row := db.QueryRow("select checking from accountholders where username = $1", name)
	row.Scan(&checking)
	row = db.QueryRow("select savings from accountholders where username = $1", name)
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
	http.HandleFunc("/register", register)
	http.HandleFunc("/confirm", confirm)
	http.HandleFunc("/login", login)
	http.HandleFunc("/apply", apply)
	http.HandleFunc("/employeelogin", employeelogin)
	http.HandleFunc("/process", process)
	http.HandleFunc("/viewaccounts", viewAccounts)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
	http.HandleFunc("/logout", logout)
	Signin.Loggedin = false
	Signin.Employee = false
	http.ListenAndServe(":7000", nil)
}
