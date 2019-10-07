package main

import (
	"fmt"

	"github.com/NGKlaure/project-0/account"
)

func main() {
	fmt.Println("banking system is running")
	//a := account.New()

	a := account.New("nad", "234", "nadine", "1132", 233.23)
	a.Summary()
	a.CreateNewAccount()

	a.Withdraw()
	a.Deposit()

}
