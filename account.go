package main

import "fmt"

type account struct {
	custUserName string
	custName     string
	custNum      string
	accountName  string //checking or saving
	accountNum   string
	availableBal float64
}

/*
func New(custUserName string, custName string, custNum string, accountName string, accountNum string, availableBal float64) account {
	a := account{custUserName, custName, custNum, accountName, accountNum, availableBal}
	return a
} */

type NewCustomer struct {
	userName    string
	password    string
	userAccList []account
}

/* func New1(userName string, password string, userAccList []account) NewCustomer {
	c := NewCustomer{userName, password, userAccList}
	return c
} */

//method to add an account to a list of customer account

func (n *NewCustomer) addNewAccount(a account) {
	n.userAccList = append(n.userAccList, a)
}

//method to list the account af a customer
func (n *NewCustomer) listCustAccount() {
	for _, account := range n.userAccList {
		fmt.Println("the list of account :", account.accountName, account.accountNum, account.custUserName)
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

	if index1 >= 0 {
		fmt.Println("already register,please login")
		c.login()
	} else {
		c.addNewCustomer()
		fmt.Println("successfully register")
		c.login()
	}
	//fmt.Println("customer list", customerList)

}

func (c *NewCustomer) login() {
	fmt.Println("the customerlist has:", customerList)
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

	if index >= 0 {
		fmt.Println(" login successfully")

		c.managebank()
	} else {
		fmt.Println("register first and try again")
	}
}

//customer mager the bank
func (c *NewCustomer) managebank() {
	fmt.Println("select an option  c to create an account, d to deposit,w to withdraw")
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

	}

}

//method to seaarch customer password in the customer list
func searchCustomerPass(pass string) int {
	for i := 0; i < len(customerList); i++ {
		if customerList[i].password == pass {
			fmt.Println(i)
			return i
		}
	}
	return -1
}

func (a *account) AccountSummary() {
	fmt.Printf("hello\n")
	fmt.Println("your account information: \n", a.custUserName, a.custName, a.custNum, a.accountName, a.accountNum, a.availableBal)
}

//method to search customer ssn in the account list
func searchCustomerSsn(ssn string) int {
	for i := 0; i < len(accountList); i++ {
		if accountList[i].custNum == ssn {
			return i
		}
	}
	return -1
}

//this method create an account
func (c *NewCustomer) CreateNewAccount() {
	var a account
	fmt.Println("enter a customer name to create your account")
	var name string
	fmt.Scanln(&name)
	a.custName = name
	fmt.Println("enter a customer ssn number")
	var ssNum string
	fmt.Scanln(&ssNum)
	a.custNum = ssNum
	fmt.Println("enter a account type ")
	var accType string
	fmt.Scanln(&accType)
	a.accountName = accType
	fmt.Println("enter the account number")
	var accNum string
	fmt.Scanln(&accNum)
	a.accountNum = accNum

	//check if a customer alredy have an account
	//we search if he is in the accountlist
	index := searchCustomerSsn(a.custNum)
	if index < 0 {

		//a.addAccount()
		c.addNewAccount(a)
		fmt.Println("account create succeffully")
		//fmt.Println("the available balance is", a.availableBal)

	} else {
		fmt.Println("already have an account")

	}

	fmt.Println("the account list has:", c.userAccList)
}

//function to add and account in an account list
func (a *account) addAccount() {
	accountList = append(accountList, *a)

}

//method to add new customer to a customer list
func (c *NewCustomer) addNewCustomer() {
	customerList = append(customerList, *c)
}

func (c *NewCustomer) Withdraw() {
	var a account
	fmt.Println(a.availableBal)
	fmt.Println("enter the amount you want to withdraw")
	var amount float64
	fmt.Scanln(&amount)
	if a.availableBal < amount {
		fmt.Println("available balance is less than the amount you want to withdraw")
	} else {

		a.availableBal = a.availableBal - amount
		fmt.Println("withdraw successfully")
		//fmt.Println(" %d is your remaining balance", a.availableBal)

	}

}
func (c *NewCustomer) Deposit() {
	var a account
	fmt.Println("enter the deposit amount")
	var amount float64
	fmt.Scanln(&amount)
	a.availableBal = a.availableBal + amount
	fmt.Println("current balance", a.availableBal)

}
