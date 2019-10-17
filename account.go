package main

import (
	"database/sql"
	"fmt"
	"os"

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

// type account struct {
// 	Firstname, Lastname, Password, Username string
// 	Balance                                 float64
// }

// type newcustomer struct {
// 	username, password string
// }

// type employee struct {
// 	Firstname, Lastname, Password, Username string
// 	// something to access customer
// }

//might not be needed
// func WriteToFile(c map[int]account) {
// 	file, err := os.Create("customers.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for i := 1; i < len(c)+1; i++ {
// 		file.WriteString((c[i]).printCustomers())
// 	}
// 	fmt.Println("write file was successful")
// 	err = file.Close()

// }

func OpenDB() *sql.DB {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	//db.Exec("INSERT INTO accounts VALUES ('Dio', 'Brando', 'DIO','ZAWARUDO',5000.75)")
	options(db)
	//addAccounts(db, "Khang", "Tran", "Deathberry", "Pika", 75.87)
	//searchByName(db, "Khang")
	//deposit(db, 50.25, 75.87, "Khang")
	//withdraw(db, 50.42, 75.87, "Khang")
	//searchByUsrNm(db, "Deathberry")
	//deleteCust(db, "Deathberry")
	//register(db, "C9Sneaky", "cosplayfnc")
	//getAll(db)
	//getAllCusts(db)
	//getAll(db)
	return db
}

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func getAllCusts(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM customers")
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

func getAllAccounts(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var username string
		var password string
		var fname string
		var lname string
		rows.Scan(&username, &password, &fname, &lname)
		fmt.Println(username, password, fname, lname)
	}
}

func searchByName(db *sql.DB, searchvalue string) {
	row := db.QueryRow("SELECT * FROM customers WHERE firstname = $1", searchvalue)
	var firstname string
	var lastname string
	var username string
	var password string
	var balance float64
	row.Scan(&firstname, &lastname, &username, &password, &balance)
	fmt.Println(firstname, lastname, username, password, balance)
}

func searchByUsrNm(db *sql.DB, searchusr string) bool {
	row := db.QueryRow("SELECT username FROM accounts WHERE username = $1", searchusr)
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

func searchByEmplUsrNm(db *sql.DB, searchusr string) bool {
	row := db.QueryRow("SELECT username FROM employees WHERE username = $1", searchusr)
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

func searchByCustNm(db *sql.DB, searchusr string) bool {
	row := db.QueryRow("SELECT username FROM customers WHERE username = $1", searchusr)
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
func addCustomers(db *sql.DB, firstname string, lastname string, username string, password string, balance float64) {
	db.Exec("INSERT INTO customers (firstname, lastname, username, password, balance)"+
		"VALUES ($1, $2, $3, $4, $5)", firstname, lastname, username, password, balance)
}

func addAccounts(db *sql.DB, firstname string, lastname string, username string, password string) {
	db.Exec("INSERT INTO accounts (firstname, lastname, username, password)"+
		"VALUES ($1, $2, $3, $4)", firstname, lastname, username, password)
}
func addCustomer(db *sql.DB, username string, password string) {
	db.Exec("INSERT INTO customers (username, password) VALUES ($1, $2)", username, password)
	fmt.Println("Customer has been added")
}
func deposit(db *sql.DB, money float64, balance float64, usrname string) {
	db.Exec("UPDATE customers SET balance = $1 WHERE username = $2", money+balance, usrname)
	fmt.Println("Updated new balance is", money+balance, usrname)
}

func withdraw(db *sql.DB, money float64, balance float64, usrname string) {
	if balance-money < 0 {
		fmt.Println("Sorry you can't withdraw that much!")
	} else {
		db.Exec("UPDATE customers SET balance = $1 WHERE username = $2", balance-money, usrname)
		fmt.Println("Updated balance is now", balance-money, usrname)
	}
}

func register(db *sql.DB, usrname string, pw string) {
	if searchByCustNm(db, usrname) == true {
		fmt.Println("Sorry this user is already in the list!")
	} else {
		db.Exec("INSERT INTO customers (username, password) VALUES ($1, $2)", usrname, pw)
		fmt.Println("User has now been added!")
	}
}
func deleteAcc(db *sql.DB, username string) {
	db.Exec("DELETE FROM accounts where username = $1", username)
	fmt.Println("Deleted ", username)
}

func deleteCust(db *sql.DB, username string) {
	db.Exec("DELETE FROM customers where username = $1", username)
	fmt.Println("Deleted ", username)
}

func login(db *sql.DB, username string, password string) bool {
	row := db.QueryRow("SELECT username, password FROM customers WHERE username = $1 AND password = $2", username, password)
	var usrname string
	var pword string
	row.Scan(&usrname, &pword)
	fmt.Println(usrname, pword)
	if (username == usrname) && (password == pword) {
		fmt.Println("The username ", username, " was found")
		fmt.Println("The password was found")
		return true
	} else {
		fmt.Println("The username ", username, "was not found")
		fmt.Println("The password was not found")
		return false
	}
}
func options(db *sql.DB) {
	fmt.Println("Are you a customer or employee? ")
	var role string
	fmt.Scanln(&role)
	if role == "customer" {
		var uname string
		var pword string
		fmt.Print("Please enter in your username: ")
		fmt.Scanln(&uname)
		fmt.Print("Please enter in your password: ")
		fmt.Scanln(&pword)
		if login(db, uname, pword) == true {
			fmt.Println("Please select an option \n1. Add customer \n2. Delete customer \n3. Register customer  \n4. Deposit \n5. Withdraw \n6. Exit")
			var input int
			fmt.Scanln(&input)
			switch input {
			case 1:
				var firstname string
				var lastname string
				var username string
				var password string
				var balance float64
				fmt.Print("Enter your first name: ")
				fmt.Scanln(&firstname)
				fmt.Print("Enter your last name: ")
				fmt.Scanln(&lastname)
				fmt.Print("Enter your username: ")
				fmt.Scanln(&username)
				fmt.Print("Enter your password: ")
				fmt.Scanln(&password)
				fmt.Print("Enter your initial balance: ")
				fmt.Scanln(&balance)
				addCustomers(db, firstname, lastname, username, password, balance)
				//break
				options(db)
			case 2:
				var username string
				fmt.Print("Enter in the username that you wish to delete")
				fmt.Scanln(&username)
				deleteCust(db, username)
				options(db)
			case 3:
				var username string
				var password string
				fmt.Print("Enter username: \n")
				fmt.Scanln(&username)
				fmt.Print("Enter password: \n")
				fmt.Scanln(&password)
				fmt.Println("You entered in", username, password)
				register(db, username, password)
				options(db)
			case 4:
				var money float64
				var balance float64
				var username string
				fmt.Print("Please enter your username: ")
				fmt.Scanln(&username)
				fmt.Print("Enter your balance: ")
				fmt.Scanln(&balance)
				fmt.Print("Enter the amount you'd like to deposit: ")
				fmt.Scanln(&money)
				deposit(db, money, balance, username)
				options(db)
			case 5:
				var money float64
				var balance float64
				var username string
				fmt.Print("Please enter your username: ")
				fmt.Scanln(&username)
				fmt.Print("Enter your balance: ")
				fmt.Scanln(&balance)
				fmt.Print("Enter the amount you'd like to withdraw: ")
				fmt.Scanln(&money)
				withdraw(db, money, balance, username)
				options(db)
			case 6:
				os.Exit(0)
			default:
				fmt.Println("Sorry that option isn't on the list")
			}

		}
	} else {
		fmt.Println("Please select an option \n1. View Customer information \n2. Approve Customer  \n3. Deny Customer")
		var input int
		fmt.Scanln(&input)
		switch input {
		case 1:
			fmt.Println("Option 1 in employee")
			getAllCusts(db)
			options(db)
		case 2:
			fmt.Println("Option 2 in employee")
			options(db)
		case 3:
			fmt.Println("Option 3 in employee")
			options(db)
		default:
			fmt.Println("Sorry there's no option like that!")
		}
	}
}
