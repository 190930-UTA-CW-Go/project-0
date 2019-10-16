package vend

import (
	"database/sql"
	"fmt"
)

/*
BuyDrink documentation
*/
func BuyDrink(db *sql.DB) {
	var selection string
	index, name, stock, _ := GetDrinks(db)

	fmt.Println("Enter your drink selection: ")
	fmt.Scanln(&selection)

	for i := range index {
		if (selection == index[i]) && (stock[i] == 0) {
			fmt.Println("Congradulations! You got nothing! Thanks for being cheeky.")
		} else if (selection == index[i]) && (stock[i] != 0) {
			fmt.Println("Now dispensing " + name[i])
			Dispense(db, index[i], stock[i])
			fmt.Println(" ")
			fmt.Println("Enjoy!")
		}
	}
}
