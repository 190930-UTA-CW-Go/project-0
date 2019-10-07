package account

import "fmt"

type account struct {
	custName     string
	custNum      string
	accountName  string
	accountNum   string
	availableBal float64
}

var accountList []account

//var c account

func New(custName string, custNum string, accountName string, accountNum string, availableBal float64) account {
	a := account{custName, custNum, accountName, accountNum, availableBal}
	return a
}

func (a account) Summary() {
	fmt.Printf("hello\n")
	fmt.Println("your account information: \n", "Name :"+a.custName, a.custNum, a.accountName, a.accountNum, a.availableBal)
}

func searchCustomerSsn(ssn string) int {
	for i := 0; i < len(accountList); i++ {
		if accountList[i].custNum == ssn {
			return i
		}

	}
	return -1

}

//this method create an account
func (a account) CreateNewAccount() {

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

	index := searchCustomerSsn(a.custNum)
	if index < 0 {

		accountList = append(accountList, a)
		fmt.Println("account create succeffully")
		//fmt.Println(a.availableBal)

	} else {
		fmt.Println("already have an account")

	}

	//fmt.Println(accountList)
}
func (a account) Withdraw() {
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
func (a account) Deposit() {

	fmt.Println("enter the deposit amount")
	var amount float64
	fmt.Scanln(&amount)
	a.availableBal = a.availableBal + amount
	fmt.Println("current balance", a.availableBal)

}
