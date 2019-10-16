package menu

import (
	"fmt"

	"github.com/gittingdavid/project-0/database"
	"github.com/gittingdavid/project-0/method"
	"github.com/gittingdavid/project-0/print"
)

// Global constant for length of account id
const idLength = 3

// Menu = Main Menu
func Menu() {
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
		Authenticate("customer")
	case "2":
		method.AddRecord("customer")
	case "3":
		Authenticate("employee")
	case "4":
		fmt.Println("> Goodbye")
		fmt.Println()
		goto Exit
	default:
		Menu()
	}
	Menu()
Exit:
}

// CustomerMenu = for Customers
// param1 = customer login id
func CustomerMenu(login string) {
	var input string
	fmt.Println("Customer:", login)
	fmt.Println("1) View Accounts")
	fmt.Println("2) Withdraw")
	fmt.Println("3) Deposit")
	fmt.Println("4) Transfer")
	fmt.Println("5) Open New Account")
	fmt.Println("6) Apply for Joint Account")
	fmt.Println("7) Exit")
	fmt.Print(": ")
	fmt.Scan(&input)
	fmt.Println()

	switch input {
	case "1":
		print.Accounts(login)
	case "2":
		method.Money(login, "withdraw")
	case "3":
		method.Money(login, "deposit")
	case "4":
		method.Transfer(login)
	case "5":
		method.OpenAccount(login)
	case "6":
		method.ApplyJoint(login)
	case "7":
		fmt.Println("> Goodbye")
		fmt.Println()
		goto End
	default:
		print.Invalid()
	}
	CustomerMenu(login)
End:
}

// EmployeeMenu for Employees
// param1 = employee login id
func EmployeeMenu(login string) {
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
		print.Table("customer")
	case "2":
		print.Table("employee")
	case "3":
		print.Accounts("")
	case "4":
		method.DeleteRecord("customer")
	case "5":
		method.DeleteRecord("employee")
	case "6":
		method.VerifyJoint()
	case "7":
		method.AddRecord("employee")
	case "8":
		fmt.Println("> Goodbye")
		fmt.Println()
		goto Exit
	default:
		print.Invalid()
	}
	EmployeeMenu(login)
Exit:
}

// Authenticate = login and password input
// param1 = identify which table "customer" or "employee"
func Authenticate(who string) {
	var email, pass, hold string
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
	row := (database.DBCon).QueryRow(sqlStatement, email)
	row.Scan(&hold)

	if pass == hold {
		fmt.Println("> Login Successful")
		fmt.Println()
		if who == "customer" {
			CustomerMenu(email)
		} else {
			EmployeeMenu(email)
		}
	} else {
		fmt.Println("> Login ID or Password do not match.")
		fmt.Println()
	}
}
