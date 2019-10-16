package vend

import (
	"database/sql"
	"fmt"
)

/*
Login documentation
*/
func Login(db *sql.DB, brand string, capacity int) {
	i, loginErr := 0, 0
	var userIn, passIn string
	account := make([]string, 3)
	pass := make([]string, 3)
	company := make([]string, 3)
	firstname := make([]string, 3)
	lastname := make([]string, 3)

	rows, _ := db.Query("SELECT * FROM servicers;")
	for rows.Next() {
		rows.Scan(&account[i], &pass[i], &company[i], &firstname[i], &lastname[i])
		i++
	}

	for i := 0; i >= 1; i = i + 0 {
		fmt.Printf("Please enter your username: ")
		fmt.Scanln(&userIn)
		fmt.Printf("Please enter your password: ")
		fmt.Scanln(&passIn)

		for a := range account {
			if (userIn == account[a]) && (passIn == pass[i]) {
				fmt.Println("Welcome, " + firstname[a] + " " + lastname[a] + ".")
				if company[a] == brand {
					fmt.Println("Thank you for restocking the vending machine.")
					Refill(db, capacity)
				} else {
					fmt.Println("Unfortunately, as a servicer of " + company[a] +
						", you are unable to service a " + brand + "vending machine.")
				}
			} else {
				fmt.Println("Your username and/or password was incorrect.")
				fmt.Printf("Type 1 to retry login, type 0 to exit login: ")
				fmt.Scanln(&loginErr)
				if loginErr == 0 {
					i++
				}
			}

		}
	}

}
