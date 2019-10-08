package main

import (
	"fmt"
	"github.com/project-0/person"
	//"bufio"
	//"os"
)

type Customer struct{
	firstname, lastname, password, username string
	balance int
}

type Employee struct{

}
func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Please enter a username")
	// username, _:= reader.ReadString('\n')
	// fmt.Println("Please enter a password")
	// password, _:= reader.ReadString('\n')
	// fmt.Println(username,password)

	cust1 := Customer{
		firstname: "Paul",
		lastname: "Walker",
		password: "fast6",
		username: "fast&furious",
		balance: 10000,
	}
	fmt.Println("Customer is ", cust1)

	var persona = person.person{
		name: "Khang",
		age: 26,
		state: "MN",
	}
	fmt.Println(person{"khang",26,"MN"})
}