package app

import (
	"fmt"
)

/*
Form documentation
*/
func Form() (string, string, string, string) {
	var acc, pas, fir, las string

	fmt.Printf("First Name: ")
	fmt.Scanln(&fir)
	fmt.Printf("Last Name: ")
	fmt.Scanln(&las)
	fmt.Printf("User Name: ")
	fmt.Scanln(&acc)
	fmt.Printf("Desired Password: ")
	fmt.Scanln(&pas)

	return acc, pas, fir, las
}
