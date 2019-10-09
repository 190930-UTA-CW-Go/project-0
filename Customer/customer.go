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
	
	type Customer struct{
		Firstname, Lastname, Password, Username string
		Balance float64
	}

	func CreateCustomer(fname string, lname string, pw string, usrnm string, balance float64) Customer{
		return Customer{fname, lname, pw, usrnm, balance} 
	}

	func (c *Customer) Amount() float64{
		return c.Balance
	}

	func (c *Customer) Withdraw(i float64) {
		if c.Balance < i{
			fmt.Println("Can't withdraw over 0")
		}else{
			c.Balance -= i
		}
	}
