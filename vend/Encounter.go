package vend

import (
	"fmt"
)

/*
Encounter Documentation
*/
func Encounter() int {
	n := 0
	fmt.Println("You encounter a vending machine! [Type number to interact]")
	fmt.Println("[1] Buy a drink.")
	fmt.Println("[2] Restock the machine")
	fmt.Println("[3] Leave.")
	fmt.Scanln(&n)
	return n
}
