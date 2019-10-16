package vend

import (
	"database/sql"
	"fmt"
)

/*
Vending is the function called when the user chooses to buy a drink, either through the Encounter
function, or through the Restock function. Both options will go back to the Navigate function, then
to the Vending function. The Vending function should be passed the database information.
The Vending function first lists the contents of the vending machine using the ListDrinks function
inside the vend package. Then, Vending calls the BuyDrink function, where the user will select the
drink of their choice, and the **machine** table will be updated to dispense one of that beverage.
Finally, Vending reprompts the user to see if they want to buy another drink, restock the machine or
leave. If they wish to buy another drink, Vending loops over and recalls ListDrinks and BuyDrink.
Otherwise, Vending returns to Navigate, then returns to main.go in the main package.

Vending also has a counter which counts the number of drinks a user buys. This is passed to both
ListDrinks and BuyDrink. ListDrink will print a unique message if the user has not bought any drinks
yet. BuyDrink will increment the counter.
*/
func Vending(db *sql.DB) int {
	r, nav, counter := 0, 0, 0

	for n := 0; n < 1; n = n + 0 {
		ListDrinks(db, counter)
		counter = BuyDrink(db, counter)

		fmt.Println("What would you like to do next?")
		fmt.Println("[1] Buy another drink")
		fmt.Println("[2] Restock the machine")
		fmt.Println("[3] Leave")
		fmt.Scanln(&nav)

		switch nav {
		case 1:
			fmt.Println(" ")
			fmt.Println("Which drink would you like to purchase?")
		case 2:
			r = 2
			n++
		case 3:
			r = 3
			n++
		default:
			for i := 0; i < 1; i = i + 0 {
				fmt.Println("Whoops! That's not an option. Try again!")
				fmt.Scanln(&nav)
				if (nav >= 1) && (nav <= 3) {
					i++
				}
			}
		}
	}
	return r
}
