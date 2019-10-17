package vend

import (
	"database/sql"
	"fmt"
)

/*
BuyDrink is called by Vending to get the users drink selection then dispense the drink.
To do this, BuyDrink calls GetDrinks to get the index, name and stock from the **machine**
table in slices. BuyDrink then gets the index of the drink the user wanted and increments
a counter passed back to Vending at the end of the function. After a drink has been
selected, BuyDrink calls the Dispense function to reduce the stock of the machine table
by one.

BuyDrink recieves the database information and the counter. It passes the database
information to GetDrinks and Dispense. Along with the database information, BuyDrink also
sends Dispense the index and stock fo the users choice.
*/
func BuyDrink(db *sql.DB, counter int) int {
	var selection string
	index, name, stock, _ := GetDrinks(db)

	fmt.Println("Enter your drink selection: ")
	fmt.Scanln(&selection)
	counter++

	for i := range index {
		if (selection == index[i]) && (stock[i] == 0) {
			fmt.Println("Congradulations! You got nothing! Thanks for being cheeky.")
		} else if (selection == index[i]) && (stock[i] != 0) {
			fmt.Println(" ")
			fmt.Println("Now dispensing " + name[i])
			Dispense(db, index[i], stock[i])
			fmt.Println("Enjoy!")
			fmt.Println(" ")
		}
	}
	return counter
}
