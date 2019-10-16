package vend

import (
	"database/sql"
	"fmt"
)

/*
Login allows registered technicians to restock the vending machine by first checking their
creditials with a username and password. If the user successfully logs in, Login will check
if that user has privaleges to restock the vending machine. If they do, Login will call the
Refill function to set all rows to max capcity.

To do this, first Login saves all the data from the **servicers** table. Then, Login prompts
the user to enter their username and login and checks with the saved data from the servicers
table. Login will match a username from the table with the username provided by the user. If
the username does not match with any usernames in thedatabase, Login will proceed with the
first name in the database. Then, Login will check the password associated with the username.
If the passwords do not match, Login will notify the user and ask if they want to try again
or go back to Navigate. If the user chooses to retry Login returns a zero, and Restock will
recall Login. If the user chooses to quit, Login will return a one and Restock will prompt
the user if they want to buy a drink or leave.

If the user successfully logs in, Login will check the company associated with the user and
compare it to the company associated with the vending machine. If the users company and the
vending machine company do not match, the usere is sent back to asked if they want to buy a
drink or leave. If the companies do match, the Refill is called to refill the vending machine.


Login recieves the database information, the company of the vending machine and the maxium
capacity of each row. It returns an integer that determines if Login should be run again.
Login sends the databse information and the maximum capacity to Refill.
*/
func Login(db *sql.DB, brand string, capacity int) int {
	i, r, n, c, loginErr := 0, 0, 0, 0, 0
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

	for n = 1; n == 1; n = n + 0 {
		fmt.Printf("Please enter your username: ")
		fmt.Scanln(&userIn)
		fmt.Printf("Please enter your password: ")
		fmt.Scanln(&passIn)

		for a := range account {
			if userIn == account[a] {
				c = a
			}
		}

		if passIn == pass[c] {
			fmt.Println("Welcome, " + firstname[c] + " " + lastname[c] + ".")
			if company[c] == brand {
				fmt.Println("Thank you for restocking the vending machine.")
				Refill(db, capacity)
				n--
			} else {
				fmt.Println("Unfortunately, as a servicer of " + company[c] +
					", you are unable to service a " + brand + " vending machine.")
				n--
			}
		} else {
			fmt.Println("Your username and/or password was incorrect.")
			fmt.Printf("Type 1 to retry login, type 0 to exit login: ")
			fmt.Scanln(&loginErr)
			if loginErr == 0 {
				n--
			}
		}
	}
	r = 1
	return r
}
