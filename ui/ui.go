package ui

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/artytheparty/project-0/applications"
	"github.com/artytheparty/project-0/bank"
	"github.com/artytheparty/project-0/usr"
)

//Menu Displays the main menu
func Menu(db *sql.DB) {
	var choice string
	fmt.Println("\n*** MAIN MENU ***")
	fmt.Println("1) Customer LogIn")
	fmt.Println("2) Employee LogIn")
	fmt.Println("3) Apply for an account")
	fmt.Println("4) Exit")
	fmt.Println("Make a choice: ")
	fmt.Scan(&choice)
	switch choice {
	case "1":
		var uHolder string
		var pHolder string
		fmt.Println("Enter your username: ")
		fmt.Scan(&uHolder)
		fmt.Println("Enter your password: ")
		fmt.Scan(&pHolder)
		userHolder := bank.GetUsrInfo(uHolder, db)
		if userHolder.GetUsrPassword() == pHolder {
			fmt.Println("Success, Welcome!")
			UserMenu(bank.GetUsrInfo(uHolder, db), db)
		} else {
			fmt.Println("Wrong username or password! Try again!")
			Menu(db)
		}
	case "2":
		var uHolder string
		var pHolder string
		fmt.Println("Enter your username: ")
		fmt.Scan(&uHolder)
		fmt.Println("Enter your password: ")
		fmt.Scan(&pHolder)
		employeeHolder := bank.GetEmployeeInfo(uHolder, db)
		if employeeHolder.GetEmployeePass() == pHolder {
			fmt.Println("Success, Welcome!")
			EmployeeMenu(db)
		} else {
			fmt.Println("Wrong username or password!")
			Menu(db)
		}
	case "3":
		applications.AskNewAccount(db)
		fmt.Println("Your Application has been submitted!")
		Menu(db)
	case "4":
		os.Exit(0)
	default:
		Menu(db)
	}
}

//UserMenu will display the menu forthe user
func UserMenu(a usr.User, db *sql.DB) {
	var choice string
	fmt.Println("\n*** User Menu ***")
	fmt.Println("1) Deposit money into account")
	fmt.Println("2) Withdraw money from account")
	fmt.Println("3) View Account Information")
	fmt.Println("4) Back to Main Menu")
	//fmt.Println("3) Transfer money to another account(must know the acount number)")
	fmt.Scan(&choice)
	switch choice {
	case "1":
		var userHolder usr.User
		var amount float64
		fmt.Println("How much would you like to deposit?: ")
		fmt.Scan(&amount)
		userHolder = bank.Deposit(a, amount, db)
		fmt.Println("Deposit Successful!")
		bank.UpdateUserDB(db, userHolder)
		UserMenu(a, db)
	case "2":
		var userHolder usr.User
		var amount float64
		fmt.Println("How much would youlike to withdraw?: ")
		fmt.Scan(&amount)
		userHolder = bank.Withdraw(a, amount, db)
		fmt.Println("Withdraw Successful!")
		bank.UpdateUserDB(db, userHolder)
		UserMenu(a, db)
	case "3":
		bank.PrintUserInfo(db, a.GetUsrUsername())
		UserMenu(a, db)
	case "4":
		Menu(db)
	default:
		UserMenu(a, db)
	}
}

//EmployeeMenu Will print out the menu for employees
func EmployeeMenu(db *sql.DB) {
	var unapprovedAccounts []applications.AccountHolder
	unapprovedAccounts = applications.ReadFile()
	var choice string
	fmt.Println("\n*** Employee Menu ***")
	fmt.Println("1) View Pending Accounts")
	fmt.Println("2) Approve Pending Accounts")
	fmt.Println("3) Back to Main Menu")
	fmt.Println("Please make a selection")
	fmt.Scan(&choice)
	switch choice {
	case "1":
		applications.PrintHolder(unapprovedAccounts)
		EmployeeMenu(db)
	case "2":
		applications.ApproveAndAddToDB(db, unapprovedAccounts)
		EmployeeMenu(db)
	case "3":
		Menu(db)
	default:
		EmployeeMenu(db)
	}

}
