package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

// Fields that start with lower case characters are package internal and not exposed,
// If you want to reference the field from another package it needs to start with an upper case character,
// eg.

// package yelk

// type PhoneOptions struct {
// 	Phone string
// 	Cc    string
// 	Lang  string
// }
var customerlist []account

var registrationlist = make([]newcustomer, 5)

type account struct {
	//id                                      int64
	Firstname, Lastname, Password, Username string
	Balance                                 float64
}

type newcustomer struct {
	username, password string
}

func (a account) printCustomers() string {
	return ("Firstname " + a.Firstname +
		"Lastname " + a.Lastname +
		"Password " + a.Password +
		"Username " + a.Username +
		"Balance " + strconv.FormatFloat(a.Balance, 'f', 6, 64)) + "\n"

}
func (nc newcustomer) Register() {
	var usrname string
	var password string
	fmt.Println("Please enter a username")
	fmt.Scanln(&usrname)
	nc.username = usrname
	fmt.Println("Please enter a password")
	fmt.Scanln(&password)
	nc.password = password

	fl := searchForUsername(usrname)
	if fl == true {
		fmt.Println("This customer is already in the list")
	} else {
		registrationlist = append(registrationlist, nc)
		fmt.Println("This customer was succesfully added")
	}
	fmt.Println("The people who are registare are ", registrationlist)
}

func (a account) addCustomer() {
	var frname string
	var lsname string
	var pswd string
	var usrnm string
	var bal float64
	fmt.Println("Please enter your first name: ")
	fmt.Scanln(&frname)
	a.Firstname = frname
	fmt.Println("Please enter your last name: ")
	fmt.Scanln(&lsname)
	a.Lastname = lsname
	fmt.Println("Please enter your password: ")
	fmt.Scanln(&pswd)
	a.Password = pswd
	fmt.Println("Please enter your username: ")
	fmt.Scanln(&usrnm)
	a.Username = usrnm
	fmt.Println("Please Deposit an initial payment to set up your balance")
	fmt.Scanln(&bal)
	a.Balance = bal

	fl := searchForUsername(usrnm)

	if fl == true {
		fmt.Println("This username has aleady signed up ")
	} else {
		customerlist = append(customerlist, a)
		fmt.Println("Customer was successfully added")
	}
	fmt.Println("The customers who've signed up are", customerlist)
}
func searchForUsername(usrname string) bool {
	for i := 0; i < len(customerlist); i++ {
		if customerlist[i].Username == usrname {
			return true
		}
	}
	return false
}

func (a *account) add(acc account) {
	customerlist = append(customerlist, acc)
	fmt.Println("Added", acc.Firstname, " ", acc.Lastname)
}

func createCustomer(fname string, lname string, pw string, usrnm string, balance float64) account {
	a := account{fname, lname, pw, usrnm, balance}
	return a
}

func (a *account) getAmount() float64 {
	return a.Balance
}

func (a *account) Withdraw(money float64) {
	fmt.Println("Your current balance is: ", a.Balance)
	if a.Balance < money {
		fmt.Println("Sorry you're out of cash!")
	} else {
		a.Balance -= money
	}
}

func (a *account) Deposit(money float64) {
	a.Balance += money
	fmt.Println("Your new balance is", a.Balance)
}

//might not be needed
func WriteToFile(c map[int]account) {
	file, err := os.Create("customers.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 1; i < len(c)+1; i++ {
		file.WriteString((c[i]).printCustomers())
	}
	fmt.Println("write file was successful")
	err = file.Close()

}

//NEED TO EDIT?
func OpenDB() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	db.Exec("INSERT INTO accounts VALUES ('Dio', 'Brando', 'DIO','ZAWARUDO',5000.75)")
	getAll(db)
}

//NEED TO EDIT
func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

//NEED TO EDIT
func getAll(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM ACCOUNTS")
	for rows.Next() {
		var Firstname string
		var Lastname string
		var Username string
		var Password string
		var Balance float64
		rows.Scan(&Firstname, &Lastname, &Username, &Password, &Balance)
		fmt.Println(Firstname, Lastname, Username, Password, Balance)
	}
}

//NEED TO EDIT
func searchByName(db *sql.DB, searchvalue string) {
	row := db.QueryRow("SELECT * FROM accounts WHERE Firstname = $1", searchvalue)
	var Firstname string
	var Lastname string
	var Username string
	var Password string
	var Balance float64
	row.Scan(&Firstname, &Lastname, &Username, &Password, &Balance)
	fmt.Println(Firstname, Lastname, Username, Password, Balance)
}
