package vend

import (
	"database/sql"
	"fmt"
)

/*
Login documentation
*/
func Login(db *sql.DB, brand string, capacity int) int {
	i, r, loginErr := 0, 0, 0
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

	for c := 0; c >= 1; c = c + 0 {
		fmt.Printf("Please enter your username: ")
		fmt.Scanln(&userIn)
		fmt.Printf("Please enter your password: ")
		fmt.Scanln(&passIn)

		for a := range account {
			fmt.Println(userIn)     //Debugging
			fmt.Println(account[a]) //Debugging
			if (userIn == account[a]) && (passIn == pass[a]) {
				fmt.Println("Welcome, " + firstname[a] + " " + lastname[a] + ".")
				if company[a] == brand {
					fmt.Println("Thank you for restocking the vending machine.")
					Refill(db, capacity)
					c++
				} else {
					fmt.Println("Unfortunately, as a servicer of " + company[a] +
						", you are unable to service a " + brand + "vending machine.")
					c++
				}
			} else {
				fmt.Println("Your username and/or password was incorrect.")
				fmt.Printf("Type 1 to retry login, type 0 to exit login: ")
				fmt.Scanln(&loginErr)
				if loginErr == 0 {
					c++
				}
			}

		}
	}
	r = 1
	return r
}
