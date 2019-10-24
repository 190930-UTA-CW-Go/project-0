package usr

import (
	"fmt"

	"github.com/artytheparty/project-0/acc"
)

//User Structure
type User struct {
	id, username, password, fName, lName string
	accounts                             []acc.Account
}

//CreateUser constructor for the User struct
func CreateUser(id string, username string, password string, fName string, lName string, accounts []acc.Account) User {
	return User{id, username, password, fName, lName, accounts}
}

//GetAccounts returns the array of user accounts.
func (a *User) GetAccounts() []acc.Account {
	return a.accounts
}

//PrintAccounts will print the associated accounts of the user the metho has been called on
func (a *User) PrintAccounts() {
	for k := range a.accounts {
		fmt.Println("Account Number: ", a.accounts[k].GetAccountNum())
		fmt.Println("Account Type: ", a.accounts[k].GetAccountType())
		fmt.Println("Account Balance: $", a.accounts[k].GetAccountBal())
		fmt.Println("----------------------------------------")
	}
}

//UpdateUserAccounts method updates the account with new updated account slice
func UpdateUserAccounts(u User, a []acc.Account) User {
	id := u.id
	username := u.username
	password := u.password
	fName := u.fName
	lName := u.lName
	return CreateUser(id, username, password, fName, lName, a)
}

//GetUsrID returns a string of the user's id which function was called on
func (a *User) GetUsrID() string {
	return a.id
}

//GetUsrUsername returns a string of the user's username which function was called on
func (a *User) GetUsrUsername() string {
	return a.username
}

//GetUsrPassword returns a string of the user's password which function was called on
func (a *User) GetUsrPassword() string {
	return a.password
}

//GetUsrFName returns a string of the user's fName which function was called on
func (a *User) GetUsrFName() string {
	return a.fName
}

//GetUsrLName returns a string of the user's lName which function was called on
func (a *User) GetUsrLName() string {
	return a.lName
}
