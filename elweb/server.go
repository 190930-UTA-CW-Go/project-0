package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/artytheparty/project-0/acc"
	"github.com/artytheparty/project-0/bank"
	"github.com/artytheparty/project-0/emp"
	_ "github.com/lib/pq"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/login", login)
	http.HandleFunc("/depwit", depwit)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/employee", employee)
	http.ListenAndServe(":6060", nil)
}

//Userholder Struct holds all the information which should be passed to the template
type Userholder struct {
	ID, Username, Password, Fname, Lname string
	Accounts                             []acc.Account
	IncorrectUnOrPass                    bool
	Loggedin                             bool
	Notsignedin                          bool
}

//Application holds information on potential users
type Application struct {
	Username string
	pass     string
	Fname    string
	lname    string
	Acctype  string
	Bal      float64
}

//Employeeholder holds info so that info can be passed to the browser for employees
type Employeeholder struct {
	Employee     emp.Employee
	Applications []Application
	Signedin     bool
	WrongPorU    bool
}

var globalUser Userholder = Userholder{}
var globalEmp Employeeholder = Employeeholder{}

func employee(response http.ResponseWriter, request *http.Request) {
	connecDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dataB, err := sql.Open("postgres", connecDB)
	defer dataB.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to DB")
	temp, _ := template.ParseFiles("template/emplog.html")
	uhold := request.FormValue("username")
	phold := request.FormValue("pw")
	fmt.Println("userinput:", uhold, "userpass:", phold)
	v := Employeeholder{}
	//navigation Bools
	v.Signedin = false
	v.WrongPorU = false
	empHolder := bank.GetEmployeeInfo(uhold, dataB)
	if phold == empHolder.GetEmployeePass() {
		v.Signedin = true
		v.Employee.EmpID = empHolder.EmpID
		v.Employee.EmpUsername = uhold
		v.Employee.EmpPass = phold
		v.Employee.EmpFName = empHolder.EmpFName
		v.Employee.EmpLName = empHolder.EmpLName
		rows, _ := dataB.Query("SELECT * FROM applications")
		for rows.Next() {
			var username string
			var pass string
			var fname string
			var lname string
			var types string
			var bal float64
			rows.Scan(&username, &pass, &fname, &lname, &types, &bal)
			v.Applications = append(v.Applications, Application{username, pass, fname, lname, types, bal})
		}
		fmt.Println(v)
		globalEmp = v

	} else {
		v.Signedin = false
		v.WrongPorU = true
	}

	//executes the template last thing
	fmt.Println(temp.Execute(response, v))
}

func logout(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("template/logout.html")
	v := globalUser
	fmt.Println(temp.Execute(response, v))
	globalUser = Userholder{}
	globalUser.IncorrectUnOrPass = false
	globalUser.Loggedin = false
	globalUser.Notsignedin = true

}

func depwit(response http.ResponseWriter, request *http.Request) {
	connecDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dataB, err := sql.Open("postgres", connecDB)
	defer dataB.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to DB")
	v := globalUser
	//holde for user

	var balanceholder float64
	temp, _ := template.ParseFiles("template/depwit.html")
	option := request.FormValue("depwith")
	accoption := request.FormValue("accselect")
	amount := request.FormValue("value")
	fmt.Println("User Inputs: ", option, accoption, amount)
	if option == "Deposit" {
		row := dataB.QueryRow("SELECT accbal FROM accounts WHERE accnum=$1", accoption)
		row.Scan(&balanceholder)
		depositholder, _ := strconv.ParseFloat(amount, 64)
		balanceholder = balanceholder + depositholder
		_, err2 := dataB.Exec("UPDATE accounts SET accbal=$1 WHERE accnum=$2", balanceholder, accoption)
		if err2 != nil {
			fmt.Println("Deposit Failed")
		}
		usrhold := bank.GetUsrInfo(v.Username, dataB)
		v.Accounts = usrhold.GetAccounts()
	}
	if option == "Withdraw" {
		row := dataB.QueryRow("SELECT accbal FROM accounts WHERE accnum=$1", accoption)
		row.Scan(&balanceholder)
		withdrawholder, _ := strconv.ParseFloat(amount, 64)
		balanceholder = balanceholder - withdrawholder
		_, err2 := dataB.Exec("UPDATE accounts SET accbal=$1 WHERE accnum=$2", balanceholder, accoption)
		if err2 != nil {
			fmt.Println("Withdraw Failed")
		}
		usrhold := bank.GetUsrInfo(v.Username, dataB)
		v.Accounts = usrhold.GetAccounts()
	}
	globalUser = v
	//executes the template last thing
	fmt.Println(temp.Execute(response, v))
}

func login(response http.ResponseWriter, request *http.Request) {
	connecDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dataB, err := sql.Open("postgres", connecDB)
	defer dataB.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to DB")
	v := Userholder{}
	//navigation variables
	v.IncorrectUnOrPass = false
	v.Loggedin = false
	v.Notsignedin = true
	temp, _ := template.ParseFiles("template/login.html")
	uhold := request.FormValue("username")
	phold := request.FormValue("pw")
	fmt.Println("userinput:", uhold, "userpass:", phold)
	usrhold := bank.GetUsrInfo(uhold, dataB)
	if phold == usrhold.GetUsrPassword() {
		v.Loggedin = true
		v.IncorrectUnOrPass = false
		v.Notsignedin = false
		v.ID = usrhold.GetUsrID()
		v.Username = uhold
		v.Password = phold
		v.Fname = usrhold.GetUsrFName()
		v.Lname = usrhold.GetUsrLName()
		v.Accounts = usrhold.GetAccounts()
		globalUser = v
	} else {
		v.Loggedin = false
		v.IncorrectUnOrPass = true
		v.Notsignedin = false
	}

	//executes the template last thing
	fmt.Println(temp.Execute(response, v))
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)
