package vend

import (
	"database/sql"
	"fmt"
)

/*
Navigate is the main function of the vend package. This should only be called once in main.go of the
main package. This function should be the only thing in this package called outside of the vend package.
It is important to note that the gen package should have a chance to run, via calling the Generate
function in the gen package, before this function is called.

Navigate first runs the Encounter function,then navigates to one of three options; the Vending function,
the Restock function or exiting Naviagte back to the main function in the main package. Additionally, if
the user tries to pick an option that is not one of those three options, it will continue ask the user to
retry until the user picks a valid option.I wrote in some hidden messages if you keep trying to pick
non-options, just for fun.

Naviate should be passed the database information and the max capactiy of each row in the vending machine.
Navigate passes the database information to the Vending function, and passes the database and capacity
information to the Restock function.
*/
func Navigate(db *sql.DB, capacity int) {
	var nav1, count int
	nav1 = Encounter()
	count = 0

	for n := 0; n < 1; n = n + 0 {
		switch nav1 {
		case 1:
			nav1 = Vending(db)
		case 2:
			nav1 = Restock(db, capacity)
		case 3:
			fmt.Println(" ")
			fmt.Println("Goodbye!")
			n++
		default:
			for i := 0; i < 1; i = i + 0 {
				fmt.Println("Whoops! That's not an option. Try again!")
				fmt.Scanln(&nav1)

				count++
				if count == 10 {
					fmt.Println(" ")
					fmt.Println("Listen, I am a computer. I promise I have more patience than you.")
					fmt.Println("Just pick a valid option already.")
					fmt.Println(" ")
				} else if count == 20 {
					fmt.Println(" ")
					fmt.Println("Okay. I will admit, I am impressed with your hatered to the numbers 1, 2 and 3")
					fmt.Println("But seriously, pick an option already.")
					fmt.Println(" ")
				} else if count == 30 {
					fmt.Println(" ")
					fmt.Println("This is getting old.")
					fmt.Println(" ")
				} else if count == 40 {
					fmt.Println(" ")
					fmt.Println("Wow.")
					fmt.Println(" ")
				} else if count == 50 {
					fmt.Println(" ")
					fmt.Println("I bet you think there is a prize for continuing on this path.")
					fmt.Println("Well, there's not. This is the 50th time you have not pressed 1, 2 or 3. Just stop.")
					fmt.Println(" ")
				} else if count == 60 {
					fmt.Println(" ")
					fmt.Println("Don't you have something better to do?")
					fmt.Println(" ")
				} else if count == 70 {
					fmt.Println(" ")
					fmt.Println("You've been here a while. Maybe you should take a break and walk around.")
					fmt.Println(" ")
				} else if count == 80 {
					fmt.Println(" ")
					fmt.Println("-________________________________-")
					fmt.Println(" ")
				} else if count == 90 {
					fmt.Println(" ")
					fmt.Println("I bet I'm encouraging this behaviour with these messages.")
					fmt.Println("I really don't have anything interesting to say.")
					fmt.Println(" ")
				} else if count == 100 {
					fmt.Println(" ")
					fmt.Println("Okay, forreal, I have better things to do than code these hidden messages.")
					fmt.Println("Thanks for you dedication and the goofs, but this is the last message.")
					fmt.Println("You've hit 100, BTW. Go do something better with your life.")
					fmt.Println(" ")
				} else if count == 110 {
					fmt.Println(" ")
					fmt.Println("( ͡° ͜ʖ ͡°)")
					fmt.Println(" ")
				}

				if (nav1 >= 1) && (nav1 <= 4) {
					i++
				}
			}
		}
	}
}
