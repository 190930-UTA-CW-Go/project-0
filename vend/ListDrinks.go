package vend

import (
	"database/sql"
	"fmt"
)

/*
ListDrinks documentation
*/
func ListDrinks(db *sql.DB) {
	index, name, stock, brand := GetDrinks(db)

	fmt.Println("Welcome to the ", brand, "vending machine!")
	for i := range stock {
		if stock[i] > 0 {
			fmt.Println("[" + index[i] + "] " + name[i])
		} else {
			fmt.Println("[" + index[i] + "] empty")
		}
	}
}
