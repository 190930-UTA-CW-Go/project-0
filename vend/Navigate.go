package vend

import (
	"database/sql"
	"fmt"
)

/*
Navigate documentation
*/
func Navigate(db *sql.DB, capacity int) {
	//var app string
	var nav1 int
	nav1 = Encounter()
	//appCall := 0

	for n := 0; n < 1; n = n + 0 {
		switch nav1 {
		case 1:
			nav1 = Vending(db)
		case 2:
			nav1 = Restock(db, capacity)
		case 3:
			fmt.Println("Goodbye!")
			n++
		// case 4:
		// 	fmt.Println("Thank you for choosing the hidden application function.")
		// 	fmt.Println("Which company did you wish to apply to?")
		// 	fmt.Println("[d] Duda-Cola")
		// 	fmt.Println("[s] SaltPhD")
		// 	fmt.Println("[t] TipsyCo")
		// 	fmt.Scanln(&app)
		// 	n++
		// 	appCall = 1
		default:
			for i := 0; i < 1; i = i + 0 {
				fmt.Println("Whoops! That's not an option. Try again!")
				fmt.Scanln(&nav1)
				if (nav1 >= 1) && (nav1 <= 4) {
					i++
				}
			}
		}
	}
}
