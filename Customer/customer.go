package customer
	
	import(
		"fmt"
	)
	// Fields that start with lower case characters are package internal and not exposed, 
	// If you want to reference the field from another package it needs to start with an upper case character,
	// eg.

	// package yelk

	// type PhoneOptions struct {
    // 	Phone string
    // 	Cc    string
    // 	Lang  string
	// }	
	
	type Customers struct{
		Firstname, Lastname, Password, Username string
		Balance float64
	}

	func CreateCustomer(fname string, lname string, pw string, usrnm string, balance float64) Customers{
		return Customers{fname, lname, pw, usrnm, balance} 
	}

	func (c *Customers) Amount() float64{
		return c.Balance
	}

	func (c *Customers) Withdraw(i float64) {
		if c.Balance < i{
			fmt.Println("Can't withdraw over 0")
		}else{
			c.Balance -= i
		}
	}
	func (c *Customers) Deposit(i float64){
		c.Balance += i
	}
