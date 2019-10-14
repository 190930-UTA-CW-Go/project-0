package vend

import "fmt"

/*
Vending documentation
*/
func Vending() int {
	r, nav := 0, 0

	for n := 0; n < 1; n = n + 0 {
		fmt.Println("You bought a drink!")
		fmt.Println("What would you like to do next?")
		fmt.Println("[1] Buy another drink")
		fmt.Println("[2] Restock the machine")
		fmt.Println("[3] Leave")
		fmt.Scanln(&nav)

		switch nav {
		case 1:
			fmt.Println("Feeling dehydrated?")
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
