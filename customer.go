package main

import (
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
var customerlist []customers

var registrationlist = make([]newcustomer, 5)

type customers struct {
	Firstname, Lastname, Password, Username string
	Balance                                 float64
}

type newcustomer struct {
	username, password string
}

func (nc newcustomer) Register() {
	var usrname string
	var password string
	fmt.Println("Please enter a username")
	fmt.Scanln(&usrname)
	nc.username = usrname
	fmt.Println("Please enter a password")
	fmt.Scanln(&password)
	nc.password = password

	fl := searchForUsername(usrname)
	if fl == true {
		fmt.Println("This customer is already in the list")
	} else {
		registrationlist = append(registrationlist, nc)
		fmt.Println("This customer was succesfully added")
	}
	fmt.Println("The people who are registare are ", registrationlist)
}

func (c customers) addCustomer() {
	var frname string
	var lsname string
	var pswd string
	var usrnm string
	var bal float64
	fmt.Println("Please enter your first name: ")
	fmt.Scanln(&frname)
	c.Firstname = frname
	fmt.Println("Please enter your last name: ")
	fmt.Scanln(&lsname)
	c.Lastname = lsname
	fmt.Println("Please enter your password: ")
	fmt.Scanln(&pswd)
	c.Password = pswd
	fmt.Println("Please enter your username: ")
	fmt.Scanln(&usrnm)
	c.Username = usrnm
	fmt.Println("Please Deposit an initial payment to set up your balance")
	fmt.Scanln(&bal)
	c.Balance = bal

	fl := searchForUsername(usrnm)

	if fl == true {
		fmt.Println("This username has aleady signed up ")
	} else {
		customerlist = append(customerlist, c)
		fmt.Println("Customer was successfully added")
	}
	fmt.Println("The customers who've signed up are", customerlist)
}
func searchForUsername(usrname string) bool {
	for i := 0; i < len(customerlist); i++ {
		if customerlist[i].Username == usrname {
			return true
		}
	}
	return false
}

func (c *customers) add(cust customers) {
	customerlist = append(customerlist, cust)
	fmt.Println("Added", cust.Firstname, " ", cust.Lastname)
}

func createCustomer(fname string, lname string, pw string, usrnm string, balance float64) customers {
	c := customers{fname, lname, pw, usrnm, balance}
	return c
}

func (c *customers) Amount() float64 {
	return c.Balance
}

func (c *customers) Withdraw(money float64) {
	fmt.Println("Your current balance is: ", c.Balance)
	if c.Balance < money {
		fmt.Println("Sorry you're out of cash!")
	} else {
		c.Balance -= money
	}
}

func (c *customers) Deposit(money float64) {
	c.Balance += money
	fmt.Println("Your new balance is", c.Balance)
}
