package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

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
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	ping(db)
	startup(db)
	//showUserTable(db)
	//showAcctsTable(db)
	//printjointaccts(db)
	//userauth(db)
	//createbankacct(db)
	//deposit(db)
	//createuseracct(db)
	//createjoint(db)
	//jointdeposit(db)
	//employeeauth(db)
}

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

//first ask user or employee?
//if employee, log in
// if user, exisiting or create new?
//have users log in, then do user whattodo

func startup(db *sql.DB) {
	fmt.Println("Hi there! Are you a user or an employee?")
	//fmt.Println("Enter 1 for employee and 2 for customer")

	fmt.Println("Enter your choice.")
	fmt.Println("[1] Employee")
	fmt.Println("[2] Customer")

	var response1 int
	var response2 int
	fmt.Scan(&response1)

	if response1 == 1 {
		hold := employeeauth(db)
		if hold == 1 {
			fmt.Println("Logged in successfully")
			empwhattodo(db)
		} else {
			fmt.Println("Your login info is incorrect.")
		}
	} else if response1 == 2 {
		fmt.Println("Enter 1 to create a new account or enter 2 to log in.")

		fmt.Scan(&response2)
		if response2 == 1 {
			createuseracct(db)
		} else if response2 == 2 {
			var hold int
			hold = userauth(db)
			if hold == 1 {
				fmt.Println("Logged in successfully")
				userwhattodo(db)
			} else {
				fmt.Println("Your login info is incorrect.")
			}
		}
	}

}

func userwhattodo(db *sql.DB) {

	fmt.Println("What would you like to do today?")
	fmt.Println("[1] Check your account balance")
	fmt.Println("[2] Make a deposit")
	fmt.Println("[3] Make a withdrawal")
	fmt.Println("[4] Open a new account")
	fmt.Println("[5] Create a joint account")
	fmt.Println("Or enter any other number to exit.")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		getacctbalance(db)

	case 2:
		deposit(db)

	case 3:
		withdrawal(db)

	case 4:
		createbankacct(db)

	case 5:
		createjoint(db)

	default:
		fmt.Println("Thank you for using Piggy Bank. Goodbye!")
	}

}

func empwhattodo(db *sql.DB) {
	fmt.Println("What would you like to do today?")
	fmt.Println("[1] To check an account balance")
	fmt.Println("[2] Make a deposit")
	fmt.Println("[3] Make a withdrawal")
	fmt.Println("[4] Create a new bank account")
	fmt.Println("[5] Create a new user account")
	fmt.Println("[6] Create a joint account")
	fmt.Println("[7] View the user data table")
	fmt.Println("[8] To view all joint accounts")
	fmt.Println("Or enter any other number to quit.")
	var response3 int
	fmt.Scan(&response3)

	switch response3 {
	case 1:
		getacctbalance(db)

	case 2:
		deposit(db)

	case 3:
		withdrawal(db)

	case 4:
		createbankacct(db)

	case 5:
		createuseracct(db)

	case 6:
		createjoint(db)

	case 7:
		showUserTable(db)

	case 8:
		showAcctsTable(db)

	default:
		fmt.Println("Thank you for using Piggy Bank.")
	}
}

//this function prints out all the values in the user accounts table except password
func showUserTable(db *sql.DB) {
	rows, _ := db.Query(`SELECT * FROM user_accounts`)

	fmt.Println("+-------------------------------------------------+")
	//fmt.Printf("|  uniqname  | FirstName  | LastName  \n")
	fmt.Printf("%-20v", "Uniqname")
	fmt.Printf("%-20v", "First Name")
	fmt.Printf("%-20v", "Last Name")
	fmt.Println()
	fmt.Println("+-------------------------------------------------+")
	for rows.Next() {
		var uniqname string
		var userfirst string
		var userlast string
		var password string

		rows.Scan(&uniqname, &userfirst, &userlast, &password)
		//	fmt.Println(uniqname, userfirst, userlast)
		//fmt.Printf("  %s  | %s  |  %s \n", uniqname, userfirst, userlast)
		fmt.Printf("%-20v", uniqname)
		fmt.Printf("%-20v", userfirst)
		fmt.Printf("%-20v", userlast)
		fmt.Println()

	}
	fmt.Println("+-------------------------------------------------+")

}

//This is a function that prints out all the values in the bank accounts table
func showAcctsTable(db *sql.DB) {
	rows, _ := db.Query(`SELECT * FROM bank_accounts`)
	for rows.Next() {
		var acctnumber int
		var uniqname string
		var acctbalance float64
		var accttype string

		rows.Scan(&acctnumber, &uniqname, &acctbalance, &accttype)
		fmt.Println("AccountNumber   Uniqname   Balance   Savings")
		fmt.Println(acctnumber, uniqname, acctbalance, accttype)
	}
}

//this is a function that creates user accounts
func createuseracct(db *sql.DB) {
	var uniqname string
	var userfirst string
	var userlast string
	var password string

	fmt.Println("Please choose a unique username. Note that no one else can have the same username as you.")
	fmt.Scan(&uniqname)

	fmt.Println("Please enter your first name.")
	fmt.Scan(&userfirst)

	fmt.Println("Please enter your last name.")
	fmt.Scan(&userlast)

	fmt.Println("Please choose a password")
	fmt.Scan(&password)

	sqlStatement := `INSERT INTO user_accounts values ($1,$2,$3,$4)`
	db.Exec(sqlStatement, uniqname, userfirst, userlast, password)
}

// this function creates a new bank account
func createbankacct(db *sql.DB) {
	var givenuniqname string
	fmt.Println("Please enter your uniqname")
	fmt.Scan(&givenuniqname)

	acctnumber := getacctnum()
	acctbalance := 0

	var accttype string
	fmt.Println("Let's name this piggy! What is this account for? ex. 'checking'")
	fmt.Scan(&accttype)

	sqlStatement := `INSERT INTO bank_accounts values($1,$2,$3,$4)`
	_, err := db.Exec(sqlStatement, acctnumber, givenuniqname, acctbalance, accttype)

	if err != nil {
		fmt.Print("An error occurred creating your account.", err)
	}
}

// generates a pseudo-random number
func getacctnum() string {
	rand.Seed(time.Now().UnixNano())
	var num int
	num = rand.Intn(10000)
	fmt.Println(num)

	acctnumstr := strconv.Itoa(num)
	return acctnumstr
}

//deposit is a function that will ask for an account number, get its balance, and then add to it
func deposit(db *sql.DB) {
	fmt.Println("Howdy! Let's make a deposit! Please enter an account number")
	var searchvalue string
	fmt.Scanln(&searchvalue)

	row := db.QueryRow("SELECT acctbalance FROM bank_accounts WHERE acctnumber = $1", searchvalue)

	var acctbalance float64
	err := row.Scan(&acctbalance)
	fmt.Println("Your previous balance was ", acctbalance)
	if err != nil {
		fmt.Println(err)
	}

	var depositamt float64
	fmt.Println("How much would you like to deposit in this account?")
	fmt.Scan(&depositamt)

	newbalance := depositamt + acctbalance
	db.Exec("UPDATE bank_accounts SET acctbalance = $1 WHERE acctnumber = $2", newbalance, searchvalue)
	fmt.Println("Way to go! Your new account balance is ", newbalance)
}

//This function asks for an account number, returns its balance, and then takes some money out of it
func withdrawal(db *sql.DB) {
	fmt.Println("Howdy! Please enter your account number to make a withdrawal.")
	var searchvalue string
	fmt.Scanln(&searchvalue)

	row := db.QueryRow("SELECT acctbalance FROM bank_accounts WHERE acctnumber = $1", searchvalue)

	var acctbalance float64
	err := row.Scan(&acctbalance)
	fmt.Println("Your previous balance was ", acctbalance)
	if err != nil {
		fmt.Println(err)
	}

	var withdrawamt float64
	fmt.Println("How much would you like to withdraw from this account?")
	fmt.Scan(&withdrawamt)

	if withdrawamt > acctbalance {
		fmt.Println(" Oink Oink! You don't have enough funds.")
	} else {
		newbalance := acctbalance - withdrawamt
		db.Exec("UPDATE bank_accounts SET acctbalance = $1 WHERE acctnumber = $2", newbalance, searchvalue)
		fmt.Println("Your remaining account balance is", newbalance)
	}

}

//this function looks up an account number and returns an account balance
func getacctbalance(db *sql.DB) {
	fmt.Println("Please enter an account number")
	var searchvalue string
	fmt.Scanln(&searchvalue)

	row := db.QueryRow("SELECT acctbalance FROM bank_accounts WHERE acctnumber = $1", searchvalue)

	var acctbalance float64
	err := row.Scan(&acctbalance)
	fmt.Println("Your account balance is ", acctbalance)
	if err != nil {
		fmt.Println(err)
	}
}

//this is a function to check your login info
func userauth(db *sql.DB) int {
	var collectuniqname string
	var collectpassword string

	fmt.Print("uniqname:")
	fmt.Scan(&collectuniqname)

	fmt.Print("password:")
	fmt.Scan(&collectpassword)

	rows, err := db.Query("SELECT * FROM user_accounts WHERE uniqname = $1 AND password = $2", collectuniqname, collectpassword)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var uniqname string
		var userfirst string
		var userlast string
		var password string

		err = rows.Scan(&uniqname, &userfirst, &userlast, &password)
		if err != nil {
			fmt.Println("Error scanning row", err)
			continue
		}
		fmt.Println("Welcome,", uniqname)
		return 1
	}
	return -1
}

//createjoint creates a new joint account for two users
func createjoint(db *sql.DB) {
	var holduniqname1 string
	var holduniqname2 string
	var holdacctname string

	fmt.Println("To create a joint account, please enter the uniqnames of each account holder.")
	fmt.Println("Uniqname of user 1: ")
	fmt.Scan(&holduniqname1)
	fmt.Println("Uniqname of user 2: ")
	fmt.Scan(&holduniqname2)

	acctnumber := getacctnum()
	acctbalance := 0

	fmt.Println("Let's name this account! What is this account for? ex 'savings'")
	fmt.Scan(&holdacctname)

	sqlStatement := `INSERT INTO joint_accounts values($1,$2,$3,$4,$5)`
	_, err := db.Exec(sqlStatement, acctnumber, holduniqname1, holduniqname2, acctbalance, holdacctname)

	if err != nil {
		fmt.Print("An error occurred creating your account.", err)
	}
}

//this function prints the joint accounts table
func printjointaccts(db *sql.DB) {
	rows, _ := db.Query(`SELECT * FROM joint_accounts`)
	for rows.Next() {
		var acctnumber string
		var uniqname1 string
		var uniqname2 string
		var acctbalance float64
		var acctname string

		rows.Scan(&acctnumber, &uniqname1, &uniqname2, &acctbalance, &acctname)
		fmt.Println(acctnumber, uniqname1, uniqname2, acctbalance, acctname)
	}
}

//jointdeposit is a function that will ask for an account number, get its balance, and then add to it
func jointdeposit(db *sql.DB) {
	fmt.Println("Howdy, Partners! Let's make a deposit! Please enter your account number")
	var searchvalue string
	fmt.Scanln(&searchvalue)

	row := db.QueryRow("SELECT acctbalance FROM joint_accounts WHERE acctnumber = $1", searchvalue)

	var acctbalance float64
	err := row.Scan(&acctbalance)
	fmt.Println("Your previous balance was ", acctbalance)
	if err != nil {
		fmt.Println(err)
	}

	var depositamt float64
	fmt.Println("How much would you like to deposit in this account?")
	fmt.Scan(&depositamt)

	newbalance := depositamt + acctbalance
	db.Exec("UPDATE joint_accounts SET acctbalance = $1 WHERE acctnumber = $2", newbalance, searchvalue)
	fmt.Println("Way to go! Your new account balance is ", newbalance)
}

//this is a function to check employee login info
func employeeauth(db *sql.DB) int {
	var collectempnumber string
	var collectpassword string

	fmt.Print("Employee Number:")
	fmt.Scan(&collectempnumber)

	fmt.Print("Employee password:")
	fmt.Scan(&collectpassword)

	rows, err := db.Query("SELECT * FROM employee_info WHERE emp_number = $1 AND emp_password = $2", collectempnumber, collectpassword)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var emp_number string
		var emp_first string
		var emp_last string
		var emp_password string
		var manager bool

		err = rows.Scan(&emp_number, &emp_first, &emp_last, &emp_password, &manager)
		if err != nil {
			fmt.Println("Error scanning row", err)
			continue
		}
		fmt.Println("Welcome,", emp_first)
		return 1
	}
	return -1
}

//This function asks for a joint account number, returns its balance, and then takes some money out of it
func jointwithdraw(db *sql.DB) {
	fmt.Println("Howdy Partners! Please enter your account number to make a withdrawal.")
	var searchvalue string
	fmt.Scanln(&searchvalue)

	row := db.QueryRow("SELECT acctbalance FROM joint_accounts WHERE acctnumber = $1", searchvalue)

	var acctbalance float64
	err := row.Scan(&acctbalance)
	fmt.Println("Your current balance is ", acctbalance)
	if err != nil {
		fmt.Println(err)
	}

	var withdrawamt float64
	fmt.Println("How much would you like to withdraw from this account?")
	fmt.Scan(&withdrawamt)

	if withdrawamt > acctbalance {
		fmt.Println(" Oink oink! You don't seem to have enough funds for this transaction.")
	} else {
		newbalance := withdrawamt + acctbalance
		db.Exec("UPDATE joint_accounts SET acctbalance = $1 WHERE acctnumber = $2", newbalance, searchvalue)
		fmt.Println("Your remaining account balance is", newbalance)
	}
}
