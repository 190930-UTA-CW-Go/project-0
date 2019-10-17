package vend

import (
	"database/sql"
	"fmt"
)

/*
Restock is called when the users requests to restock the vending machine. Restock
retrieves the company associated with the vending machine through the GetDrinks
function. Then Restock calls Login to verify that the user can restock the vending
machine. Login will return a one when the user quits or restocks the machine. After
this, Restock will prompt the user if they want to buy a drink, returning them to
the Vending function then to BuyDrink, or leave, sending them back to main.go in the
main package. If Login returns a zero, Restock will restart Login.

Restock recieves the database information and the maxium capacity of the rows in
the vending machine. Restock returns an integer telling Naviage which function to
navigate to. Restock sends the database information to GetDrinks and Login. In
addition, Login also recieves from Restock the maxium capacity and company associated
with the vending machine.
*/
func Restock(db *sql.DB, capacity int) int {
	loginErr, r, nav := 0, 0, 0

	for n := 0; n < 1; n = n + 0 {
		_, _, _, brand := GetDrinks(db)

		fmt.Println(" ")
		fmt.Println("Please login to restock the machine.")
		loginErr = Login(db, brand, capacity)

		if loginErr == 1 {
			fmt.Println("What would you like to do next?")
			fmt.Println("[1] Buy a drink")
			fmt.Println("[2] Leave")
			fmt.Scanln(&nav)

			switch nav {
			case 1:
				r = 1
				n++
			case 2:
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
	}
	return r
}
