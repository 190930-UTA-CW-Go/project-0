package vend

import (
	"database/sql"
	"fmt"
)

/*
ListDrinks prints the contents of the vending machine to the console. If the user
has not bought a drink yet, ListDrinks will print a welcome message. ListDrinks will
get the index, name, stock and brand of the vending machine in slices from the
GetDrinks function. Then ListDrinks will print the index and name of the drink to
the console. If the stock is zero, ListDrinks will print the index and "empty" to
tell the user that row is out of stock.

List drinks recives the database information and the counter mentioned in the Vending
documentation. ListDrinks passes the database information to the GetDrinks function.
*/
func ListDrinks(db *sql.DB, counter int) {
	index, name, stock, brand := GetDrinks(db)

	if counter == 0 {
		fmt.Println("Welcome to the ", brand, "vending machine!")
	}

	for i := range stock {
		if stock[i] > 0 {
			fmt.Println("[" + index[i] + "] " + name[i])
		} else {
			fmt.Println("[" + index[i] + "] empty")
		}
	}
}
