package app

import (
	"fmt"
	"os"
)

/*
Apply documentation
*/
func Apply(app string) {
	var acc, pass, comp, first, last string
	switch app {
	case "d":
		comp = "Duda-Cola"
	case "s":
		comp = "Salt-PhD"
	case "t":
		comp = "TipsyCo"
	}

	fmt.Println("Welcome to the", comp, "application page.")

	if len(os.Args) >= 6 {
		fmt.Println("Thank you for choosing the quick apply option.")
		fmt.Println(" ")
		acc, pass, first, last = QuickApply()
	} else if (len(os.Args) > 2) && (len(os.Args) < 6) {
		fmt.Println("It appears that your quick apply was partially complete.")
		fmt.Println("Please use the application form below.")
		fmt.Println(" ")
		acc, pass, first, last = Form()
	} else {
		fmt.Println("Please fill out this form to apply.")
		fmt.Println(" ")
		acc, pass, first, last = Form()
	}
	application := NewTech(acc, pass, comp, first, last)
	application.print(comp)

}
