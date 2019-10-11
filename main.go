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

	// Menu
	var input string
	fmt.Println("Welcome!")
	fmt.Println("1) Login Returning Customer")
	fmt.Println("2) Sign Up New Customer")
	fmt.Println("3) Employee Only")
	fmt.Print(": ")
	fmt.Scan(&input)
	fmt.Println()

	if input == "1" {
		var email string
		var pass string
		fmt.Print("Login: ")
		fmt.Scan(&email)
		fmt.Print("Password: ")
		fmt.Scan(&pass)

		if authenticate(db, "customer", email, pass) == true {
			fmt.Println("Successfully logged in!")
		} else {
			fmt.Println("Incorrect password")
		}
	} else if input == "2" {
		addTable(db, "customer")

	} else if input == "3" {

	} else {
		fmt.Println("Sorry invalid input")
	}

	printTable(db, "customer")
	fmt.Println()
	//printTable(db, "employee")
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

// Check if email matches the password in database and return bool
// param1 = database
// param2 = identify either "customer" or "employee"
// param3 = supposed email
// param4 = supposed password
func authenticate(db *sql.DB, who string, email string, pass string) bool {
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
		return true
	}
	return false
}
