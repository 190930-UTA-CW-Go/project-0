package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
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
	Firstname, Lastname, Password, Username string
	Balance                                 float64
}

type newcustomer struct {
	username, password string
}

type employee struct {
	Firstname, Lastname, Password, Username string
	// something to access customer
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
func OpenDB() *sql.DB {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	//db.Exec("INSERT INTO accounts VALUES ('Dio', 'Brando', 'DIO','ZAWARUDO',5000.75)")
	addAccounts(db, "Khang", "Tran", "Deathberry", "Pika", 75.87)
	searchByName(db, "Khang")
	//deposit(db, 50.25, 75.87, "Khang")
	withdraw(db, 50.42, 75.87, "Khang")
	//searchByUsrNm(db, "Deathberry")
	searchByUsrNm(db, "Deathberry")
	register(db, "Deathberry", "I eat birds")
	getAll(db)
	return db
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
	rows, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var firstname string
		var lastname string
		var username string
		var password string
		var balance float64
		rows.Scan(&firstname, &lastname, &username, &password, &balance)
		fmt.Println(firstname, lastname, username, password, balance)
	}
}

//NEED TO EDIT
func searchByName(db *sql.DB, searchvalue string) {
	row := db.QueryRow("SELECT * FROM accounts WHERE firstname = $1", searchvalue)
	var firstname string
	var lastname string
	var username string
	var password string
	var balance float64
	row.Scan(&firstname, &lastname, &username, &password, &balance)
	fmt.Println(firstname, lastname, username, password, balance)
}

func searchByUsrNm(db *sql.DB, searchusr string) bool {
	row := db.QueryRow("SELECT username FROM newcustomers WHERE username = $1", searchusr)
	var username string
	row.Scan(&username)
	if searchusr == username {
		fmt.Println("The username ", searchusr, " was found")
		return true
	} else {
		fmt.Println("The username ", searchusr, "was not found")
		return false
	}
}
func addAccounts(db *sql.DB, firstname string, lastname string, username string, password string, balance float64) {
	db.Exec("INSERT INTO accounts (firstname, lastname, username, password, balance)"+
		"VALUES ($1, $2, $3, $4, $5)", firstname, lastname, username, password, balance)
}

func deposit(db *sql.DB, money float64, balance float64, fname string) {
	db.Exec("UPDATE accounts SET balance = $1 WHERE firstname = $2", money+balance, fname)
	fmt.Println("Updated new balance is", money+balance, fname)
}

func withdraw(db *sql.DB, money float64, balance float64, fname string) {
	if balance-money < 0 {
		fmt.Println("Sorry you can't withdraw that much!")
	} else {
		db.Exec("UPDATE accounts SET balance = $1 WHERE firstname = $2", balance-money, fname)
		fmt.Println("Updated balance is now", balance-money, fname)
	}
}

func register(db *sql.DB, usrname string, pw string) {
	if searchByUsrNm(db, usrname) == true {
		fmt.Println("Sorry this user is already in the list!")
	} else {
		db.Exec("INSERT INTO newcustomers username, password)"+"VALUES ($1, $2)", usrname, pw)
		fmt.Println("User has now been added!")
	}
}
