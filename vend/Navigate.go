package vend

import (
	"fmt"
)

/*
Navigate documentation
*/
func Navigate() {
	var nav1 int
	nav1 = Encounter()
	for n := 0; n < 1; n = n + 0 {
		switch nav1 {
		case 1:
			nav1 = Vending()
		case 2:
			nav1 = Restock()
		case 3:
			fmt.Println("Goodbye!")
			n++
		default:
			for i := 0; i < 1; i = i + 0 {
				fmt.Println("Whoops! That's not an option. Try again!")
				fmt.Scanln(&nav1)
				if (nav1 >= 1) && (nav1 <= 3) {
					i++
				}
			}
		}
	}
}
