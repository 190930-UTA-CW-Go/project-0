package guest

import "fmt"

// Customer data
type Customer struct {
	userName string
	password string
	name     string
	balance  float64
}

// NewCustomer is a Constructor for Customer
func NewCustomer(userName string, password string,
	name string, balance float64) *Customer {
	n := Customer{
		userName: userName,
		password: password,
		name:     name,
		balance:  balance,
	}
	return &n
}

func (a Customer) String() string {
	var output string
	//t := strconv.Itoa(a.balance)
	t := fmt.Sprintf("%.2f", a.balance)
	output = a.userName + " | " + a.password + " | " + a.name + " | $" + t + "\n"
	return output
}

// Balance returns the amount of money in a customer's balance
func (a *Customer) Balance() float64 {
	return a.balance
}

// Withdraw removes money from a customer's balance
func (a *Customer) Withdraw(i float64) {
	a.balance -= i
}

// Deposit adds money to a customer's balance
func (a *Customer) Deposit(i float64) {
	a.balance += i
}
