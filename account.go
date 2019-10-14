package main

import (
	"database/sql"
	"fmt"
	"os"

	dbconnection "github.com/NGKlaure/project-0/dbConnection"
)

type account struct {
	accountNum string
	custName   string

	accountType  string //checking or saving
	availableBal float64
}

var db *sql.DB

//Newcustomer data
type NewCustomer struct {
	userName    string
	password    string
	userAccList []account
}

//method to add an account to a list of customer account

func (n *NewCustomer) addNewAccount(a account) {
	n.userAccList = append(n.userAccList, a)

}

//method to display the account af a customer
func (n *NewCustomer) displayCustAccountInfos() {
	for _, account := range n.userAccList {
		fmt.Println("the list of account :", account.accountNum, account.custName, account.accountType, account.availableBal)
	}
}

var accountList []account

var customerList []NewCustomer

//var customerList = make([]newCustomer, 2)

func (c *NewCustomer) Register() {
	fmt.Println("enter a userName to register")
	var uName string
	fmt.Scanln(&uName)
	c.userName = uName
	fmt.Println("enter a password")
	var psswrd string
	fmt.Scanln(&psswrd)
	c.password = psswrd

	//fmt.Println("customer length is:", len(customerList))
	index1 := searchCustomerPass(psswrd)

	if index1 == psswrd {
		fmt.Println(" the password you enter already exist")
		fmt.Println("Please login to selection and action")
		c.login()
	} else {
		db := dbconnection.DbConnection()
		defer db.Close()
		//c.addNewCustomer()
		db.Exec("INSERT INTO newCustomer (name, password) VALUES ($1,$2)", uName, psswrd)

		fmt.Println(" successfully register")
		fmt.Println("  Now login to select an action")
		c.login()
	}
}

func (c *NewCustomer) login() {
	//fmt.Println("the customerlist has:", customerList)
	fmt.Println("please enter login information ")
	fmt.Println("please enter your user name ")
	var usName string
	fmt.Scanln(&usName)
	c.userName = usName
	fmt.Println("enter your password")
	var pass string
	fmt.Scanln(&pass)
	c.password = pass

	index := searchCustomerPass(pass)

	if index == pass {
		fmt.Println(" login successfully")

		c.managebank()
	} else {
		fmt.Println("register first and try again")
	}
}

//customer mager the bank
func (c *NewCustomer) managebank() {
	fmt.Println("select an option  c to create an account, d to deposit,w to withdraw , e to exit")
	var choice string
	fmt.Scanln(&choice)
	switch choice {
	case "C":
		c.CreateNewAccount()
	case "c":
		c.CreateNewAccount()
	case "w":
		c.Withdraw()
	case "d":
		c.Deposit()
	case "e":
		os.Exit(0)

	}
}

//method to seaarch customer password in the customer list
func searchCustomerPass(password string) string {

	var upass string

	db := dbconnection.DbConnection()
	defer db.Close()
	row := db.QueryRow("select password from newCustomer where password = $1", password)
	row.Scan(&upass)

	return upass
}

//method to search customer accountnum in the account list
func searchCustomeaccNum(accountNum string) string {
	var acnum string

	db := dbconnection.DbConnection()
	defer db.Close()
	row := db.QueryRow("select accountNum from account where accountNum  = $1", accountNum)
	row.Scan(&acnum)

	return acnum

}

//this method create an account
func (c *NewCustomer) CreateNewAccount() {
	var a account
	fmt.Println("enter a customer name to create your account")
	var name string
	fmt.Scanln(&name)
	a.custName = name

	fmt.Println("enter a account type ")
	var accType string
	fmt.Scanln(&accType)
	a.accountType = accType
	fmt.Println("enter the account number")
	var accNum string
	fmt.Scanln(&accNum)
	a.accountNum = accNum
	var availbal float64 = 0.0
	a.availableBal = availbal

	//check if a customer alredy have an account
	//we search if the account number is alredy in our db
	index := searchCustomeaccNum(accNum)
	if index == accNum {
		fmt.Println("already have an account")

	} else {
		db := dbconnection.DbConnection()
		defer db.Close()
		//c.addNewCustomer()
		db.Exec("INSERT INTO account (accountNum, custName,accountType,availableBal) VALUES ($1,$2,$3,$4)", name, accType, accNum, availbal)
		//a.addAccount()
		c.addNewAccount(a)
		fmt.Println("account create succeffully")
		//fmt.Println("the available balance is", a.availableBal)

	}

	fmt.Println("the account list has:", c.userAccList)
}

//method to add new customer to a customer list

//method to return the balance of a given account number
func getAccountBalance(accountNum string) float64 {

	var balance float64
	db := dbconnection.DbConnection()
	defer db.Close()
	row := db.QueryRow("select availableBal from account where accountNum = $1", accountNum)
	row.Scan(&balance)

	return balance

}

func (c *NewCustomer) Withdraw() {

	fmt.Println("enter the account number you want to withdraw from")
	var accountNum string
	fmt.Scanln(&accountNum)
	fmt.Println("enter the amount you want to withdraw")
	var amount float64
	fmt.Scanln(&amount)
	var availBalance float64

	if accountNum != searchCustomeaccNum(accountNum) {
		fmt.Println("invalid account number")
	} else if accountNum == "" {
		fmt.Println("invalid account number")

	} else {
		availBalance = getAccountBalance(accountNum)

		if amount > availBalance {
			fmt.Println(" not enough money to withdraw from")
		} else {
			db := dbconnection.DbConnection()
			defer db.Close()
			db.Exec("UPDATE account SET balance =$1 WHERE accountNum =$2", availBalance-amount, accountNum)
			fmt.Println("withdraw successfull")
			fmt.Println("the remainning balance is:", availBalance-amount)
		}
	}
}

//method to deposit into an  account
func (c *NewCustomer) Deposit() {
	fmt.Println("enter the account number you want to Deposit into")
	var accountNum string
	fmt.Scanln(&accountNum)
	fmt.Println("enter the amount you want to Deposit")
	var amount float64
	fmt.Scanln(&amount)
	var availBalance float64
	if accountNum != searchCustomeaccNum(accountNum) {
		fmt.Println("invalid account number")
	} else if accountNum == "" {
		fmt.Println("invalid account number")
	} else {
		availBalance = getAccountBalance(accountNum)

		db := dbconnection.DbConnection()
		defer db.Close()
		db.Exec("UPDATE account SET balance =$1 WHERE accountNum =$2", availBalance+amount, accountNum)
		fmt.Println("Deposit successfull")
		fmt.Println("the totalbalance is:", availBalance+amount)
	}

}
