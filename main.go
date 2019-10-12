package main

import (
    "html/template"
    "net/http"
    "fmt"
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
    Username string
    Password string
    Approved bool
    Denied bool
    Pending bool
    Notapplied bool
}
type Applications struct {
    Username string
    Firstname string
    Lastname string
    Address string
    Phone  string
}
type Accountholders struct {
    Username string
    Checking int
    Savings int
}
type ViewInfo struct {
    Ap []Applications
    Usr []Users
    Ac []Accountholders
    Singleuser  Users
    Singleaccount  Accountholders
    Singleapp Applications
}
var currentUser string = "-1"
var employee bool = false;
func index(response http.ResponseWriter, request *http.Request){
     temp, _ := template.ParseFiles("templates/index.html")
     temp.Execute(response,nil)	
}
func register(response http.ResponseWriter, request *http.Request){
     temp, _ := template.ParseFiles("templates/register.html")
     temp.Execute(response,nil)      
}
func confirm(response http.ResponseWriter, request *http.Request){
	db := connect()
	temp, _ := template.ParseFiles("templates/confirm.html")
	var query string
        user := Users{}
        user.Username =  request.FormValue("name")
        user.Password = request.FormValue("pw")
	if uniqueName(db, user.Username) == false || len(user.Username)<3{
		db.Close()
		if len(user.Username)<3{
			temp, _ := template.ParseFiles("templates/nametooshort.html")
                        temp.Execute(response,nil)

		}else {
			temp, _ = template.ParseFiles("templates/notunique.html")
			temp.Execute(response,nil)
		}
		return
	}
	if len(user.Password) < 3 {
		db.Close()
		temp, _ := template.ParseFiles("templates/pwtooshort.html")
		temp.Execute(response,nil)
		return
	}
	query = "INSERT INTO users (username, password, status)"
        query += " VALUES ($1, $2, $3)"
	db.QueryRow(query, user.Username, user.Password, "notapplied")
	defer db.Close()
	temp.Execute(response,user)
}
func login(response http.ResponseWriter, request *http.Request){
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
	if currentUser == "-1"{
        	user.Username =  request.FormValue("name")
        	user.Password = request.FormValue("pw")
		if uniqueName(db, user.Username)==true{
			db.Close()
                	temp, _ := template.ParseFiles("templates/namenotfound.html")
                	temp.Execute(response,nil)
			return
		}
		if passwordMatches(db,user.Username, user.Password)==false {
			db.Close()
                	temp, _ := template.ParseFiles("templates/pwnotmatch.html")
                	temp.Execute(response,nil)
                	return
		}
		currentUser = user.Username
		employee = false
	}else{
		user.Username = currentUser
	}
	status = getStatus(db,user.Username)
	if status == "notapplied"{
		user.Notapplied = true
	}else if status == "approved" {
		user.Approved = true
	}else if status == "denied" {
		user.Denied = true
	}else {
		user.Pending = true
	}
	defer db.Close()
	view.Singleuser = user
	ac.Checking, ac.Savings = getBalance(db,user.Username)
	view.Singleaccount = ac
        temp.Execute(response,view)
}
func deposit(response http.ResponseWriter, request *http.Request){
	db := connect()
	var current int
	var query,statement,status string
        ac := Accountholders{}
        user := Users{}
        view := ViewInfo{}
	user.Username = currentUser
        user.Approved = false
        user.Denied = false
        user.Pending = false
        user.Notapplied = false
	temp, _ := template.ParseFiles("templates/login.html")
	amount,_ := strconv.Atoi(request.FormValue("amount"))
	choice := request.FormValue("account")
	if choice == "checking" {
		 query = "SELECT checking FROM accountholders WHERE username=$1" 
	}else{
		 query = "SELECT savings FROM accountholders WHERE username=$1" 
	}
	row := db.QueryRow(query,currentUser)
        row.Scan(&current)
	amount += current
	if choice == "checking" {
                 statement = "UPDATE accountholders SET checking=$1 WHERE username=$2" 
        }else{
                 statement = "UPDATE accountholders SET savings=$1 WHERE username=$2" 
        }
        db.Exec(statement,amount,currentUser)
	status = getStatus(db,user.Username)
        if status == "notapplied"{
                user.Notapplied = true
        }else if status == "approved" {
                user.Approved = true
        }else if status == "denied" {
                user.Denied = true
        }else {
                user.Pending = true
        }
        defer db.Close()
        view.Singleuser = user
        ac.Checking, ac.Savings = getBalance(db,user.Username)
        view.Singleaccount = ac
	temp.Execute(response,view)
}
func employeelogin(response http.ResponseWriter, request *http.Request){
	db := connect()
	temp, _ := template.ParseFiles("templates/employeelogin.html")
	user := Users{}
        user.Username = request.FormValue("name")
        user.Password = request.FormValue("pw")
	if !employee {
		if user.Username =="" {
			db.Close()
			temp, _ := template.ParseFiles("templates/notauthorized.html")
			temp.Execute(response,nil)
			return
		}
		if uniqueEmployeeName(db, user.Username)==true{
                	db.Close()
                	temp, _ := template.ParseFiles("templates/employeenotfound.html")
                	temp.Execute(response,nil)
                	return
        	}
        	if employeePasswordMatches(db,user.Username, user.Password)==false {
                	db.Close()
                	temp, _ := template.ParseFiles("templates/employeepwnotmatch.html")
                	temp.Execute(response,nil)
                	return
        	}
	}	 
	currentUser = user.Username
	employee = true 
	view := ViewInfo{}
	rows, _ := db.Query("select * from applications")
        for rows.Next() {
                var username,firstname,lastname,address,phone string
		var ap = Applications{}
                rows.Scan(&username,&firstname,&lastname,&address,&phone)
                ap.Username = username
                ap.Firstname = firstname
                ap.Lastname = lastname
                ap.Address = address
                ap.Phone = phone
		view.Ap = append(view.Ap,ap)
        }
        defer db.Close()
	temp.Execute(response, view)
}
func process(response http.ResponseWriter, request *http.Request){
	db := connect()
	temp, _ := template.ParseFiles("templates/employeelogin.html")
	var query,statement,choice,action string
	choice = request.FormValue("choice")
	action = request.FormValue("action")
	statement = "delete from applications where username = $1"
        db.Exec(statement,choice)
	if action == "approve"{
	    query = "INSERT INTO accountholders (username, checking, savings)"
            query += " VALUES ($1, $2, $3)"
            db.QueryRow(query,choice, 0, 0)
	    statement = "UPDATE users SET status=$1 WHERE username=$2"
            db.Exec(statement,"approved",choice)

	}else{
	     statement = "UPDATE users SET status=$1 WHERE username=$2"
            db.Exec(statement,"denied",choice)

	}
	view := ViewInfo{}
        rows, _ := db.Query("select * from applications")
        for rows.Next() {
                var username,firstname,lastname,address,phone string
                var ap = Applications{}
                rows.Scan(&username,&firstname,&lastname,&address,&phone)
                ap.Username = username
                ap.Firstname = firstname
                ap.Lastname = lastname
                ap.Address = address
                ap.Phone = phone
                view.Ap = append(view.Ap,ap)
        }
        defer db.Close()

	temp.Execute(response,view)
}
func viewAccounts(response http.ResponseWriter, request *http.Request){
	db := connect()
	temp, _ := template.ParseFiles("templates/viewaccounts.html")
	if !employee {
               db.Close()
               temp, _ := template.ParseFiles("templates/notauthorized.html")
               temp.Execute(response,nil)
               return
        }
	view := ViewInfo{}
	rows, _ := db.Query("select * from accountholders")
        for rows.Next() {
                var username string
		var checking, savings int
		var ac = Accountholders{}
                rows.Scan(&username,&checking,&savings)
                ac.Username = username
		ac.Checking = checking
		ac.Savings = savings
		view.Ac = append(view.Ac,ac)
	}
        defer db.Close()
        temp.Execute(response, view)

}
func apply(response http.ResponseWriter, request *http.Request){
        db := connect()
	var query string
        temp, _ := template.ParseFiles("templates/apply.html")
	statement := "UPDATE users SET status=$1 WHERE username=$2"
	db.Exec(statement,"pending",currentUser)
	ap := Applications{}
	ap.Firstname =  request.FormValue("first")
        ap.Lastname = request.FormValue("last")
	ap.Address =  request.FormValue("address")
        ap.Phone = request.FormValue("phone")
	query = "INSERT INTO applications (username, firstname, lastname, address, phone)"
        query += " VALUES ($1, $2, $3, $4, $5)"
        db.QueryRow(query, currentUser, ap.Firstname, ap.Lastname, ap.Address, ap.Phone)	
	defer db.Close()
	temp.Execute(response,ap)
}
func uniqueName(db *sql.DB, name string) bool {
	rows, _ := db.Query("select username from users")
        for rows.Next() {
                var username string
                rows.Scan(&username)
		if(name == username){
			return false
		}
        }
	return true
}
func passwordMatches(db *sql.DB, name string, password string) bool {
	var pw string
	row := db.QueryRow("select password from users where username = $1",name)
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
                if(name == username){
                        return false
                }
        }
        return true
}
func employeePasswordMatches(db *sql.DB, name string, password string) bool {
        var pw string
        row := db.QueryRow("select password from employees where username = $1",name)
        row.Scan(&pw)
        if password == pw {
                return true
        }
        return false
}

func getStatus(db *sql.DB, name string) string{
	var status string
	row := db.QueryRow("select status from users where username = $1",name)
	row.Scan(&status)
	return status
}
func getBalance(db *sql.DB, name string) (int,int){
	var checking, savings int
	row := db.QueryRow("select checking from accountholders where username = $1",name)
	row.Scan(&checking)
	row = db.QueryRow("select savings from accountholders where username = $1",name)
        row.Scan(&savings)
	return checking,savings
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
     return db;
}
func main() {
     http.HandleFunc("/",index)
     http.HandleFunc("/register",register)
     http.HandleFunc("/confirm",confirm)
     http.HandleFunc("/login",login)
     http.HandleFunc("/apply",apply)
     http.HandleFunc("/employeelogin", employeelogin)
     http.HandleFunc("/process", process)
     http.HandleFunc("/viewaccounts", viewAccounts)
     http.HandleFunc("/deposit", deposit)
     http.ListenAndServe(":7000",nil)
}

