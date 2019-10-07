package main

import (
	"fmt"
	"github.com/Tony-Moon/project-0/vending"
)

func main(){
	// Let's start by building a row inside the vending machine. 
	var a1 = vending.Row {
		Row:      "A1",
		Beverage: "Coke",
		Stock:    5,
	}
	var a1p vending.Readout = a1
	fmt.Println(a1p)
	
	a1.Dispense()
	a1p = a1 			// Update the readout
	fmt.Println(a1p)

	a1.Restock(5)
	a1p = a1
	fmt.Println(a1p)

	// Row A1 down, but I should consider finding a way to generate the number of rows and filling them automatically
	// Generated an array for the row index
	junk := vending.Generate(2, 5)
	fmt.Println(junk)
	fmt.Println(len(junk))
}