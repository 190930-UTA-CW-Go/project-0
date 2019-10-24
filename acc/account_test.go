package acc

import (
	"fmt"
	"testing"
)

func TestDeposit(t *testing.T) {
	hold := CreateAccount("2", "1", "s", 243.50)
	hold2 := hold.Deposit(20)
	if hold2.AccBal == 263.50 {
		fmt.Println("Success")
	} else {
		panic("BRUUUH")
	}
}
func TestWithdraw(t *testing.T) {
	hold := CreateAccount("2", "1", "s", 243.50)
	hold2 := hold.Withdraw(200)
	if hold2.AccBal == 43.50 {
		fmt.Println("Success")
	} else {
		fmt.Println("Error happened")
	}
}
func TestCreateAccount(t *testing.T) {
	hold := CreateAccount("2", "1", "s", 243.50)
	if hold.AccNum == "2" && hold.UsrID == "1" && hold.AccType == "s" && hold.AccBal == 243.50 {
		fmt.Println("Success")
	} else {
		fmt.Println("Error happened")
	}
}
func TestUpdateAccountSlice(t *testing.T) {
	var holder []Account
	hold := CreateAccount("2", "1", "s", 243.50)
	hold2 := CreateAccount("2", "1", "s", 243.50)
	tester := CreateAccount("2", "1", "s", 500)
	holder = append(holder, hold)
	holder = append(holder, hold2)
	holder = UpdateAccountSlice(holder, tester, 0)
	if holder[0].AccBal == 500 {
		fmt.Println("Success")
	} else {
		fmt.Println("Error happened")
	}
}
func TestGetAccountNum(t *testing.T) {
	hold := CreateAccount("2", "1", "s", 243.50)
	if hold.GetAccountNum() == "2" {
		println("Success")
	} else {
		fmt.Println("Error happened")
	}
}
func TestGetUsrID(t *testing.T) {
	hold := CreateAccount("2", "1", "s", 243.50)
	if hold.GetUsrID() == "1" {
		println("Success")
	} else {
		fmt.Println("Error happened")
	}
}
func TestGetAccountType(t *testing.T) {
	hold := CreateAccount("2", "1", "s", 243.50)
	if hold.GetAccountType() == "s" {
		println("Success")
	} else {
		fmt.Println("Error happened")
	}
}
func TestGetAccountBal(t *testing.T) {
	hold := CreateAccount("2", "1", "s", 243.50)
	if hold.GetAccountBal() == 243.50 {
		println("Success")
	} else {
		fmt.Println("Error happened")
	}
}
