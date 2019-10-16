package vend

import (
	"database/sql"
	"fmt"
)

/*
Restock documentation
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
