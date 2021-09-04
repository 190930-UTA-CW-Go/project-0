package account

import "fmt"

type account struct {
	custUserName string
	custName     string
	custNum      string
	accountName  string
	accountNum   string
	availableBal float64
}

func New(custUserName string, custName string, custNum string, accountName string, accountNum string, availableBal float64) account {
	a := account{custUserName, custName, custNum, accountName, accountNum, availableBal}
	return a
}

type newCustumer struct {
	userName string
	password string
}

func New1(userName string, password string) newCustumer {
	c := newCustumer{userName, password}
	return c
}

var accountList []account

var customerList = make([]newCustumer, 2)

func (c newCustumer) Register() {
	fmt.Println("enter a userName")
	var uName string
	fmt.Scanln(&uName)
	c.userName = uName
	fmt.Println("enter a password")
	var psswrd string
	fmt.Scanln(&psswrd)
	c.password = psswrd

	fmt.Println("customer length is:", len(customerList))
	index1 := searchCustomerPass(psswrd)

	if index1 > 0 {
		fmt.Println("already register")
	} else {
		customerList = append(customerList, c)
		fmt.Println("successfully register")
	}
	fmt.Println("customer list", customerList)

}

//method to seaarch customer password in the customer list
func searchCustomerPass(pass string) int {
	for i := 0; i < len(customerList); i++ {
		if customerList[i].password == pass {
			return i
		}
	}
	return -1
}

func (a account) Summary() {
	fmt.Printf("hello\n")
	fmt.Println("your account information: \n", "Name :"+a.custName, a.custNum, a.accountName, a.accountNum, a.availableBal)
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
func (a *account) CreateNewAccount() {

	fmt.Println("enter a customer name")
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
	fmt.Println("enter a account number")
	var accNum string
	fmt.Scanln(&accNum)
	a.accountNum = accNum

	fmt.Println("enter the amount you want to deposit")
	var initAmount float64
	fmt.Scanln(&initAmount)
	a.availableBal = initAmount

	//check if a customer alredy have an account
	//we search if he is in the accountlist
	index := searchCustomerSsn(a.custNum)
	if index < 0 {

		accountList = append(accountList, *a)
		fmt.Println("account create succeffully")
		fmt.Println("the available balance is", a.availableBal)

	} else {
		fmt.Println("already have an account")

	}

	//fmt.Println(accountList)
}
func (a *account) Withdraw() {
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
func (a *account) Deposit() {

	fmt.Println("enter the deposit amount")
	var amount float64
	fmt.Scanln(&amount)
	a.availableBal = a.availableBal + amount
	fmt.Println("current balance", a.availableBal)

}
