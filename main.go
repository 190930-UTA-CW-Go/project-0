package main

import (
	"fmt"

	"github.com/NGKlaure/project-0/account"
)

func main() {
	fmt.Println("banking system is running")
	//a := account.New()

	//a := account.New("mmm", "nad", "234", "nadine", "1132", 233.23)
	c := account.New1("mimi", "1234")

	c.Register()

	//a.Summary()
	//a.CreateNewAccount()

	//a.Withdraw()
	//a.Deposit()

}
