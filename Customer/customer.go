package customer

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
