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
	fmt.Println("Welcome!")
	fmt.Println("1) Login Returning Customer")
	fmt.Println("2) Register New Customer")
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
		fmt.Println("> Goodbye")
		goto Exit
	default:
		menu()
	}
	menu()
Exit:
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
	fmt.Printf("%-30v", "Login ID:")
	fmt.Printf("%-20v", "Password:")
	fmt.Printf("%-20v", "First Name:")
	fmt.Printf("%-20v", "Last Name:")
	fmt.Println()
	fmt.Print("================================================")
	fmt.Println("==============================================")
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
		//fmt.Println(email, pass, first, last)

		fmt.Printf("%-30v", email)
		fmt.Printf("%-20v", pass)
		fmt.Printf("%-20v", first)
		fmt.Printf("%-20v", last)
		fmt.Println()
	}

	if count == 0 {
		fmt.Println("No Data in Table")
	}

	fmt.Print("================================================")
	fmt.Println("==============================================")
	fmt.Println()
}

// Authenticate login and password input
// param1 = identify either "customer" or "employee"
func authenticate(who string) {
	var email string
	var pass string

	fmt.Print("Login: ")
	fmt.Scan(&email)
	fmt.Print("Password: ")
	fmt.Scan(&pass)

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
		fmt.Println("> Login Successful")
		fmt.Println()
		if who == "customer" {
			customerMenu(email)
		} else {
			employeeMenu(email)
		}
	} else {
		fmt.Println("> Login ID or Password do not match.")
		fmt.Println()
		menu()
	}
}

// Menu for Customers
// param1 = customer login id
func customerMenu(login string) {
	fmt.Println("Customer:", login)
	var input string
	fmt.Println("1) View Accounts")
	fmt.Println("2) Open New Account")
	fmt.Println("3) Join Account")
	fmt.Println("4) Exit")
	fmt.Print(": ")
	fmt.Scan(&input)
	fmt.Println()

	switch input {
	case "1":
		printAccounts(login)
	case "2":
		openAccount(login)
	case "3":
		fmt.Println("weeee")
	case "4":
		fmt.Println("> Goodbye")
		fmt.Println()
		goto End
	default:
		customerMenu(login)
	}
	customerMenu(login)
End:
}

// Menu for Employees
// param1 = employee login id
func employeeMenu(login string) {
	fmt.Println("Employee:", login)
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
		fmt.Println("> Goodbye")
		fmt.Println()
		goto Exit
	default:
		employeeMenu(login)
	}
	employeeMenu(login)
Exit:
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

// Open new account input name and balance
// param1 = customer login id
func openAccount(login string) {
	var name string
	var balance float32
	fmt.Print("Insert Account Name: ")
	fmt.Scan(&name)
	fmt.Print("Insert Account Balance: $")
	fmt.Scan(&balance)
	fmt.Println()

	sqlStatement := `
	insert into account (email, acc_type, acc_balance)
	values ($1, $2, $3)`

	_, err := db.Exec(sqlStatement, login, name, balance)
	if err != nil {
		panic(err)
	}
}

// Print accounts associated with login id
// param1 = customer login id
func printAccounts(login string) {
	fmt.Printf("%-30v", "Login ID:")
	fmt.Printf("%-20v", "Account Type:")
	fmt.Printf("%-20v", "Account Balance:")
	fmt.Printf("%-20v", "Account Number:")
	fmt.Println()
	fmt.Print("================================================")
	fmt.Println("==============================================")
	sqlStatement := `select * from account`

	rows, _ := db.Query(sqlStatement)
	var count int

	for rows.Next() {
		count++
		var email string
		var name string
		var balance float32
		var number int
		rows.Scan(&email, &name, &balance, &number)
		//fmt.Println(email, name, balance, number)

		fmt.Printf("%-30v", email)
		fmt.Printf("%-20v$", name)
		fmt.Printf("%-20v", balance)
		fmt.Printf("%-20v", number)
		fmt.Println()
	}

	if count == 0 {
		fmt.Println("No Data in Table")
	}

	fmt.Print("================================================")
	fmt.Println("==============================================")
	fmt.Println()
}
