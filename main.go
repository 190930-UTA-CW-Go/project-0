package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	// Connect to database
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	//Main Menu
	menu(db)
}

// Interactive text menu
//param1 = database
func menu(db *sql.DB) {
	var input string

Top:
	fmt.Println("Welcome!")
	fmt.Println("1) Login Returning Customer")
	fmt.Println("2) Sign Up New Customer")
	fmt.Println("3) Employee Only")
	fmt.Println("4) Exit")
	fmt.Print(": ")
	fmt.Scan(&input)
	fmt.Println()

	switch input {
	case "1":
		authenticate(db, "customer")
	case "2":
		addTable(db, "customer")
	case "3":
		authenticate(db, "employee")
	case "4":
		fmt.Println("Bye")
	default:
		goto Top
	}

	//printTable(db, "customer")
	//fmt.Println()
}

// Insert new data to table
// param1 = database
// param2 = identify either "customer" or "employee"
func addTable(db *sql.DB, who string) {
	var email, pass, first, last string
	fmt.Print("Insert Email: ")
	fmt.Scan(&email)
	fmt.Print("Insert Password: ")
	fmt.Scan(&pass)
	fmt.Print("Insert First Name: ")
	fmt.Scan(&first)
	fmt.Print("Insert Last Name: ")
	fmt.Scan(&last)

	sqlStatement := ``
	if who == "customer" {
		sqlStatement = `
		insert into customer (email, pass, first_name, last_name)
		values ($1, $2, $3, $4)`
	} else {
		sqlStatement = `
		insert into employee (email, pass, first_name, last_name)
		values ($1, $2, $3, $4)`
	}

	_, err := db.Exec(sqlStatement, email, pass, first, last)
	if err != nil {
		panic(err)
	}
}

// Prints table
// param1 = database
// param2 = identify either "customer" or "employee"
func printTable(db *sql.DB, who string) {
	sqlStatement := ``
	if who == "customer" {
		sqlStatement = `select * from customer`
	} else {
		sqlStatement = `select * from employee`
	}

	rows, _ := db.Query(sqlStatement)
	for rows.Next() {
		var email string
		var pass string
		var first string
		var last string
		rows.Scan(&email, &pass, &first, &last)
		fmt.Println(email, pass, first, last)
	}
}

// Authenticate login and password input
// param1 = database
// param2 = identify either "customer" or "employee"
func authenticate(db *sql.DB, who string) {
	var email string
	var pass string
Top:
	fmt.Print("Login: ")
	fmt.Scan(&email)
	fmt.Print("Password: ")
	fmt.Scan(&pass)
	fmt.Println()

	sqlStatement := ``
	if who == "customer" {
		sqlStatement = `select pass from customer where email=$1`
	} else {
		sqlStatement = `select pass from employee where email=$1`
	}

	row := db.QueryRow(sqlStatement, email)
	var hold string
	row.Scan(&hold)

	if pass == hold {
		fmt.Println("Login Successful")
		if who == "customer" {
			customerMenu()
		} else {
			employeeMenu()
		}
	} else {
		var input string
		fmt.Println("Incorrect Credentials")
		fmt.Println("1) Retry")
		fmt.Println("2) Go to Menu")
		fmt.Print(": ")
		fmt.Scan(&input)
		fmt.Println()

		switch input {
		case "1":
			goto Top
		case "2":
			menu(db)
		case "3":
			fmt.Println("Invalid input going to Menu")
			menu(db)
		}
	}
}

func customerMenu() {
	var input string
	fmt.Println("1) View Accounts")
	fmt.Println("2) Open New Account")
	fmt.Println("3) Join Accoun")
	fmt.Print(": ")
	fmt.Scan(&input)

	switch input {
	case "1":
		fmt.Println("ahhh")
	case "2":
		fmt.Println("ooooo")
	case "3":
		fmt.Println("weeee")
	default:
		fmt.Println("dead")
	}
}

func employeeMenu() {
	fmt.Println("Inside employeeMenu()")
}
