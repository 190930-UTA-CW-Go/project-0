package functions

import "fmt"

//ClientMenu show the client's menu
func ClientMenu() {
	fmt.Println("WELCOME TO THE BANK")
	fmt.Println("")
	fmt.Println("MENU")
	var clientM = make(map[int]string)
	clientM[1] = "1.- Apply to open and account"
	clientM[2] = "2.- Deposit"
	clientM[3] = "3.- Withdraw"
	clientM[4] = "4.- Transfer"
	clientM[5] = "5.- Balance"

	var i int
	for i = 0; i <= len(clientM); i++ {
		fmt.Println(clientM[i])
	}
}
