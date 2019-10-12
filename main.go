package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Global variable for database
var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	// Connect to database
	var err error
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	//Main Menu
	menu()
}

// Interactive Text Menu
func menu() {
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
		authenticate("customer")
	case "2":
		addRecord("customer")
	case "3":
		authenticate("employee")
	case "4":
		fmt.Println("Bye")
	default:
		goto Top
	}
}

// Insert new data to table
// param1 = identify either "customer" or "employee"
func addRecord(who string) {
	var email, pass, first, last string
	fmt.Print("Insert Email: ")
	fmt.Scan(&email)
	fmt.Print("Insert Password: ")
	fmt.Scan(&pass)
	fmt.Print("Insert First Name: ")
	fmt.Scan(&first)
	fmt.Print("Insert Last Name: ")
	fmt.Scan(&last)
	fmt.Println()

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
// param1 = identify either "customer" or "employee"
func printTable(who string) {
	fmt.Println("==============================")
	sqlStatement := ``
	if who == "customer" {
		sqlStatement = `select * from customer`
	} else {
		sqlStatement = `select * from employee`
	}

	rows, _ := db.Query(sqlStatement)
	var count int

	for rows.Next() {
		count++
		var email string
		var pass string
		var first string
		var last string
		rows.Scan(&email, &pass, &first, &last)
		fmt.Println(email, pass, first, last)
	}

	if count == 0 {
		fmt.Println("No Data in Table")
	}

	fmt.Println("==============================")
	fmt.Println()
}

// Authenticate login and password input
// param1 = identify either "customer" or "employee"
func authenticate(who string) {
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
		fmt.Println("Username or Password do not match.")
		fmt.Println("1) Retry")
		fmt.Println("2) Go to Menu")
		fmt.Print(": ")
		fmt.Scan(&input)
		fmt.Println()

		switch input {
		case "1":
			goto Top
		case "2":
			menu()
		case "3":
			fmt.Println("Invalid input going to Menu")
			menu()
		}
	}
}

func customerMenu() {
	var input string
	fmt.Println("1) View Accounts")
	fmt.Println("2) Open New Account")
	fmt.Println("3) Join Account")
	fmt.Print(": ")
	fmt.Scan(&input)
	fmt.Println()

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
	var input string
	fmt.Println("1) Print Customer Table")
	fmt.Println("2) Print Employee Table")
	fmt.Println("3) Delete Customer Record")
	fmt.Println("4) Delete Employee Record")
	fmt.Println("5) Approve/Deny Customer Applications")
	fmt.Println("6) Add New Employee")
	fmt.Println("7) Exit")
	fmt.Print(": ")
	fmt.Scan(&input)
	fmt.Println()

	switch input {
	case "1":
		printTable("customer")
	case "2":
		printTable("employee")
	case "3":
		deleteRecord("customer")
	case "4":
		deleteRecord("employee")
	case "5":
		//
	case "6":
		addRecord("employee")
	case "7":
		fmt.Println("Bye")
		goto End
	default:
		employeeMenu()
	}

	employeeMenu()
End:
}

// Delete record
// param1 = identify either "customer" or "employee"
func deleteRecord(who string) {
	var email string
	fmt.Print("Login ID: ")
	fmt.Scan(&email)

	sqlStatement := ``
	if who == "customer" {
		sqlStatement = `delete from customer where email = $1`
	} else {
		sqlStatement = `delete from employee where email = $1`
	}
	res, err := db.Exec(sqlStatement, email)
	if err == nil {
		count, err := res.RowsAffected()
		if err == nil {
			if count == 0 {
				fmt.Println("> Invalid Login ID")
			} else {
				fmt.Println("> Successfully Deleted")
			}
			fmt.Println()
		}
	}
}
