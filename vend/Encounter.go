package vend

import (
	"fmt"
)

/*
Encounter simply propmts the user the first time they find the vending machine. Then,
it records and returns the users choice.
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
