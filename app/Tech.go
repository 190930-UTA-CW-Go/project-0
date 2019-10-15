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

func (te Tech) result(r int, comp string) {
	if r == 0 {
		fmt.Println("After careful review by our team of experts, who is definatly not a single robot, we have decided to decline your application.")
	} else {
		fmt.Println("Congradulations! You are now a certified vending machine technician of " + comp + "!")
		fmt.Println("You can restock any " + comp + "vending machine by using your username and password to login.")
	}
}
