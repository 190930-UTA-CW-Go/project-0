package app

import "fmt"

/*
Tech documentation
*/
type Tech struct {
	account   string
	password  string
	company   string
	firstname string
	lastname  string
}

func (te Tech) print(comp string) {
	fmt.Println("Thank you, " + te.firstname + ", for applying to " + comp + ".")
}

func (te Tech) result(stat int, comp string) {
	if stat == 0 {
		fmt.Println("Sorry, but that user name has already been taken. Please choose a different username:")
	} else if stat == 1 {
		fmt.Println("After careful review by our team of experts, who is definatly not a single robot, we have decided to decline your application.")
		fmt.Println(" ")
	} else {
		fmt.Println("Congradulations! You are now a certified vending machine technician of " + comp + "!")
		fmt.Println("You can restock any " + comp + " vending machine by using your username and password to login.")
		fmt.Println(" ")
	}
}
