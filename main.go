package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("welcome to my bank App!")
	fmt.Println("What can i help you,press the menu option")
	var menu int
	fmt.Scanln(&menu)

	switch menu {

	case 1:
		fmt.Println("Newaccount")
		fmt.Println("The  account number of the customer ")
		fmt.Println()
		fmt.Println(rand.Intn(100))
	case 2:
		fmt.Println("Deposit")
	case 3:
		fmt.Println("Withdraw")
	case 4:
		fmt.Println("Transfert")
	default:
		fmt.Println("createaccount")
	}

}
