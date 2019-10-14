package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

// Global variable for database
var db *sql.DB

// Global constant for length of account id
const idLength = 3

// Connection string information
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	// Connecting to database
	var err error
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	//Set Seed
	rand.Seed(time.Now().UTC().UnixNano())

	//Main Menu
	menu()

}

// Randomly generate id that doesn't start with 0
// Double check it doesn't already exist in account table
// return = the generated id
func generateID() (s string) {
Top:
	var x int
	for i := 0; i < idLength; i++ {
		x = rand.Intn(9)
		for i == 0 && x == 0 {
			x = rand.Intn(9)
		}
		s += strconv.Itoa(x)
	}

	var hold string
	sqlStatement := `select acc_id from account where acc_id = $1`
	result := db.QueryRow(sqlStatement, s)
	result.Scan(&hold)

	if hold == "" {
		return
	} else {
		goto Top
	}
}

// Main Menu
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
		fmt.Println()
		goto Exit
	default:
		menu()
	}
	menu()
Exit:
}

// Insert new record to table
// param1 = identify which table "customer" or "employee"
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
// param1 = identify which table "customer" or "employee"
func printTable(who string) {
	var count int
	var email string
	var pass string
	var first string
	var last string
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

	for rows.Next() {
		// count variable used as empty table error checker
		count++
		rows.Scan(&email, &pass, &first, &last)
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
// param1 = identify which table "customer" or "employee"
func authenticate(who string) {
	var email string
	var pass string
	var hold string

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
	var input string
	fmt.Println("Customer:", login)
	fmt.Println("1) View Accounts")
	fmt.Println("2) Open New Account")
	fmt.Println("3) Apply for Joint Account")
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
		applyJoint(login)
	case "4":
		fmt.Println("> Goodbye")
		fmt.Println()
		goto End
	default:
		invalidPrint()
	}
	customerMenu(login)
End:
}

// Menu for Employees
// param1 = employee login id
func employeeMenu(login string) {
	var input string
	fmt.Println("Employee:", login)
	fmt.Println("1) Print Customer Table")
	fmt.Println("2) Print Employee Table")
	fmt.Println("3) Print Account Table")
	fmt.Println("4) Delete Customer Record")
	fmt.Println("5) Delete Employee Record")
	fmt.Println("6) Verify Joint Accounts")
	fmt.Println("7) Add New Employee")
	fmt.Println("8) Exit")
	fmt.Print(": ")
	fmt.Scan(&input)
	fmt.Println()

	switch input {
	case "1":
		printTable("customer")
	case "2":
		printTable("employee")
	case "3":
		printAccounts("")
	case "4":
		deleteRecord("customer")
	case "5":
		deleteRecord("employee")
	case "6":
		verifyJoint()
	case "7":
		addRecord("employee")
	case "8":
		fmt.Println("> Goodbye")
		fmt.Println()
		goto Exit
	default:
		invalidPrint()
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
				invalidPrint()
			} else {
				fmt.Println("> Successfully Deleted")
				fmt.Println()
			}
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
	insert into account (email, acc_type, acc_balance, acc_id)
	values ($1, $2, $3, $4)`

	_, err := db.Exec(sqlStatement, login, name, balance, generateID())
	if err != nil {
		panic(err)
	}
}

// Print accounts associated with login id
// param1 = customer login id
func printAccounts(login string) {
	var count int
	var email string
	var name string
	var balance float32
	var number int
	fmt.Printf("%-30v", "Login ID:")
	fmt.Printf("%-20v", "Account Type:")
	fmt.Printf("%-20v", "Account Balance:")
	fmt.Printf("%-20v", "Account Number:")
	fmt.Println()
	fmt.Print("================================================")
	fmt.Println("==============================================")

	sqlStatement := ""
	if login == "" {
		sqlStatement = `select * from account order by email`
		rows, _ := db.Query(sqlStatement)
		for rows.Next() {
			// count variable used as empty table error checker
			count++
			rows.Scan(&email, &name, &balance, &number)
			fmt.Printf("%-30v", email)
			fmt.Printf("%-20v$", name)
			fmt.Printf("%-20v", balance)
			fmt.Printf("%-20v", number)
			fmt.Println()
		}
	} else {
		sqlStatement = `select * from account where email = $1`
		rows, _ := db.Query(sqlStatement, login)
		for rows.Next() {
			// count variable used as empty table error checker
			count++
			rows.Scan(&email, &name, &balance, &number)
			fmt.Printf("%-30v", email)
			fmt.Printf("%-20v$", name)
			fmt.Printf("%-20v", balance)
			fmt.Printf("%-20v", number)
			fmt.Println()
		}
	}

	if count == 0 {
		fmt.Println("No Data in Table")
	}

	fmt.Print("================================================")
	fmt.Println("==============================================")
	fmt.Println()
}

func applyJoint(login string) {
	var oneNumber string
	var twoNumber string
	var hold1 string
	var hold2 string

	fmt.Print("Input Your Account Number: ")
	fmt.Scan(&oneNumber)
	fmt.Print("Input Joint Account Number: ")
	fmt.Scan(&twoNumber)
	fmt.Println()

	if oneNumber == twoNumber {
		invalidPrint()
	} else {
		sqlStatement := `select email from account where acc_id = $1`

		result1 := db.QueryRow(sqlStatement, oneNumber)
		result1.Scan(&hold1)

		result2 := db.QueryRow(sqlStatement, twoNumber)
		result2.Scan(&hold2)

		if hold1 == "" || hold2 == "" || hold1 != login || hold1 == hold2 {
			invalidPrint()
		} else {
			fmt.Println("Submitted Joint Account Request")
			fmt.Println()
			sqlStatement = `
			insert into joint (email1, email2, num1, num2)
			values ($1, $2, $3, $4)`

			_, err := db.Exec(sqlStatement, hold1, hold2, oneNumber, twoNumber)
			if err != nil {
				panic(err)
			}
		}
	}
}

// Approve/Deny Customer Applications
func verifyJoint() {
	var count = printJoints()
	var input string
	var hold string

	if count != 0 {
		fmt.Print("Input: ")
		fmt.Scan(&input)

		sqlStatement := `select index from joint where index = $1`
		result := db.QueryRow(sqlStatement, input)
		result.Scan(&hold)

		if hold == "" {
			invalidPrint()
		} else {
			var choice string
			fmt.Println()
			fmt.Println("1) Approve")
			fmt.Println("2) Deny")
			fmt.Print(": ")
			fmt.Scan(&choice)

			switch choice {
			case "1":
				// Get acc_id values
				var idOne, idTwo string
				sqlOne := `select num1 from joint where index = $1`
				sqlTwo := `select num2 from joint where index = $1`
				resOne := db.QueryRow(sqlOne, input)
				resOne.Scan(&idOne)
				resTwo := db.QueryRow(sqlTwo, input)
				resTwo.Scan(&idTwo)

				// Use acc_id values to get acc_balance
				var balOne, balTwo float32
				sqlThree := `select acc_balance from account where acc_id = $1`
				resThree := db.QueryRow(sqlThree, idOne)
				resThree.Scan(&balOne)
				resFour := db.QueryRow(sqlThree, idTwo)
				resFour.Scan(&balTwo)

				// Update the affect records
				var newID string = generateID()
				sqlUpdate := `
				update account
				set acc_type = $1, acc_balance = $2, acc_id = $3
				where acc_id = $4`
				_, err := db.Exec(sqlUpdate, "JOINT", balOne+balTwo, newID, idOne)
				if err != nil {
					panic(err)
				}

				_, err = db.Exec(sqlUpdate, "JOINT", balOne+balTwo, newID, idTwo)
				if err != nil {
					panic(err)
				}

				// Delete the joint record now that it's been approved
				deleteJoint(input, "Joint Application Approved")

			case "2":
				deleteJoint(input, "Joint Application Denied")
			default:
			}
		}
		fmt.Println()
	}
}

// Deletes a record from the joint table
// param1 = index primary key to delete record
// param2 = string message to output
func deleteJoint(input string, print string) {
	sqlStatement := `delete from joint where index = $1`
	res, err := db.Exec(sqlStatement, input)
	if err == nil {
		count, err := res.RowsAffected()
		if err == nil {
			if count == 0 {
				invalidPrint()
			} else {
				fmt.Println(">", print)
				fmt.Println()
			}
		}
	}
}

// Prints joint table
func printJoints() (count int) {
	count = 0
	var index string
	var email1 string
	var email2 string
	var num1 int
	var num2 int

	fmt.Print("   ")
	fmt.Printf("%-25v", "#1 Login:")
	fmt.Printf("%-25v", "#2 Login:")
	fmt.Printf("%-20v", "#1 Account ID:")
	fmt.Printf("%-20v", "#2 Account ID:")
	fmt.Println()
	fmt.Print("================================================")
	fmt.Println("==============================================")

	sqlStatement := "select * from joint"
	rows, _ := db.Query(sqlStatement)
	for rows.Next() {
		count++

		rows.Scan(&index, &email1, &email2, &num1, &num2)
		fmt.Print(index + ") ")
		fmt.Printf("%-25v", email1)
		fmt.Printf("%-25v", email2)
		fmt.Printf("%-20v", num1)
		fmt.Printf("%-20v", num2)
		fmt.Println()
	}

	if count == 0 {
		fmt.Println("No Data in Table")
	}

	fmt.Print("================================================")
	fmt.Println("==============================================")
	fmt.Println()
	return
}

func invalidPrint() {
	fmt.Println("> Invalid Input")
	fmt.Println()
}
