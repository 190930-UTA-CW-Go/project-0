package main

import (
	"fmt"
	"os"

	dbconnection "github.com/NGKlaure/project-0/dbConnection"
)

func ELogin() {
	//fmt.Println("the customerlist has:", customerList)
	fmt.Print("=====================================")
	fmt.Println("===================================")
	fmt.Println("Enter employee login information ")
	fmt.Println("Enter your user name ")
	var usName string
	fmt.Scanln(&usName)
	//e.userName = usName
	fmt.Println("enter your password")
	var pass string
	fmt.Scanln(&pass)
	//e.password = pass

	index := searchEmployeePass(pass)

	if index == pass {
		fmt.Println(" login successfully")

		eManagebank()
	} else {
		fmt.Println("register first and try again")
		Menu()
	}
}

func searchEmployeePass(password string) string {
	var upass string

	db := dbconnection.DbConnection()
	defer db.Close()
	row := db.QueryRow("select password from employee where password = $1", password)
	row.Scan(&upass)

	return upass

}

func eManagebank() {
	fmt.Print("==============================")
	fmt.Println("============================")
	fmt.Println("select an option")
	fmt.Println("1) to view customer informations")
	fmt.Println("2)to approuve or delete customer")
	fmt.Println("3)to list all accounts")
	fmt.Println("4)to list all joint accounts")
	fmt.Println("5) to exit")

	var choice string

	fmt.Scanln(&choice)
	switch choice {
	case "1":
		viewCustomerInfos()
	case "2":
		aprouve()
	case "3":
		listAllAccount()
	case "4":
		listAllJointAccount()
	case "5":
		os.Exit(0)

	}
}

func viewCustomerInfos() {
	fmt.Println("enter the customer name ")
	var custName string
	fmt.Scanln(&custName)

	db := dbconnection.DbConnection()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM account WHERE custName = $1", custName)

	for rows.Next() {
		var accountNum string
		var custName string
		var age int
		var accountType string
		var availableBal float64
		rows.Scan(&accountNum, &custName, &age, &accountType, &availableBal)
		fmt.Println("Account number is:", accountNum)
		fmt.Println("Name is:", custName)
		fmt.Println("AGE IS:", age)
		fmt.Println("Account type is:", accountType)
		fmt.Println("Available balance is:", availableBal)

	}
	Menu()
}

func aprouve() {
	fmt.Println("enter the account holder name you want to approve")
	var name string
	fmt.Scanln(&name)

	if getCustAge(name) < 18 {
		db := dbconnection.DbConnection()
		defer db.Close()
		db.Exec("Delete from account where name =$1", name)
		fmt.Println("we can't aprouve it, too young")

	} else {
		fmt.Println("The account is approuved")
	}
	Menu()

}

func listAllAccount() {

	fmt.Printf("%-18v", "Account Num:")
	fmt.Printf("%-15v", "AccHolder Name:")
	fmt.Printf("%-15v", "Account age:")
	fmt.Printf("%-15v", "Account Type:")
	fmt.Printf("%-15v", "Account Balance:")
	fmt.Println()
	fmt.Print("==============================================")
	fmt.Println("============================================")
	db := dbconnection.DbConnection()
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM account")

	for rows.Next() {
		var accountNum string
		var custName string
		var age int
		var accountType string
		var availableBal float64
		rows.Scan(&accountNum, &custName, &age, &accountType, &availableBal)

		fmt.Printf("%-18v", accountNum)
		fmt.Printf("%-15v", custName)
		fmt.Printf("%-15v", age)
		fmt.Printf("%-15v", accountType)
		fmt.Printf("%-15v", availableBal)
		fmt.Println()

		//fmt.Println(accountNum, custName, age, accountType, availableBal)
	}
	Menu()

}

func listAllJointAccount() {

	fmt.Printf("%-18v", "Account Num:")
	fmt.Printf("%-15v", "Holder Name1:")
	fmt.Printf("%-15v", "Holder Name2:")
	fmt.Printf("%-15v", "Account Type:")
	fmt.Printf("%-15v", "Account Balance:")
	fmt.Println()
	fmt.Print("================================================")
	fmt.Println("==============================================")
	db := dbconnection.DbConnection()
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM jointAccount")

	for rows.Next() {
		var id int
		var name1 string
		var name2 string
		var accType string
		var availableBal float64
		rows.Scan(&id, &name1, &name2, &accType, &availableBal)

		fmt.Printf("%-15v", id)
		fmt.Printf("%-15v", name1)
		fmt.Printf("%-15v", name2)
		fmt.Printf("%-15v", accType)
		fmt.Printf("%-15v", availableBal)
		fmt.Println()

		fmt.Println(id, name1, name2, accType, availableBal)
	}
	Menu()

}

func getCustAge(name string) int {

	var aage int
	db := dbconnection.DbConnection()
	defer db.Close()
	row := db.QueryRow("select age from account where custName = $1", name)
	row.Scan(&aage)

	return aage

}
