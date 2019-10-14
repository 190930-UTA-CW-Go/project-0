package app

import (
	"fmt"
)

/*
Apply documentation
*/
func Apply(app string) {
	var comp string
	switch app {
	case "d":
		fmt.Println("Thank you for applying to Duda-Cola.")
		comp = "Duda-Cola"
	case "s":
		fmt.Println("Thank you for applying to Salt-PhD.")
		comp = "Salt-PhD"
	case "t":
		fmt.Println("Thank you for applying to TipsyCo.")
		comp = "TipsyCo"
	}

	fmt.Println("Please fill out this form to apply.")
	acc, pass, first, last := Form()
	application := NewTech(acc, pass, comp, first, last)
	application.print()

}
