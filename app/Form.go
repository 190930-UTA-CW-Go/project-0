package app

import (
	"fmt"
)

/*
Form documentation
*/
func Form() (string, string, string, string) {
	var acc, pas, fir, las string

	fmt.Println("account:")
	fmt.Scanln(&acc)

	fmt.Println("password:")
	fmt.Scanln(&pas)

	fmt.Println("first name:")
	fmt.Scanln(&fir)

	fmt.Println("last name:")
	fmt.Scanln(&las)

	return acc, pas, fir, las
}
