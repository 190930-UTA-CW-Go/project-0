package main

import (
	"fmt"
	"math/rand"
)

//ClientMenu show the client's menu
func ClientMenu() {
	fmt.Println("WELCOME TO THE BANK")
	fmt.Println("")
	fmt.Println("MENU")
	var clientM = make(map[int]string)
	clientM[1] = "1.- Register as a client"
	clientM[2] = "2.- Apply to open and account"
	clientM[3] = "3.- Withdraw"
	clientM[4] = "4.- Deposit"
	clientM[5] = "5.- Transfer"
	clientM[6] = "6.- Exit"

	var i int
	for i = 0; i < len(clientM); i++ {
		fmt.Println(clientM[i])
	}
}

//Customer is the structure of bank's customer
type Customer struct {
	FirstName, lastName, street, city, state, zip, userName, pass string
}

//Register ask for the client's information
func Register() Customer {
	fmt.Println("Personal Information")
	var customer Customer
	fmt.Println("First Name: ")
	fmt.Scanln(&customer.FirstName)
	fmt.Println("Last Name: ")
	fmt.Scanln(&customer.lastName)
	fmt.Println("street: ")
	fmt.Scanln(&customer.street)
	fmt.Println("City: ")
	fmt.Scanln(&customer.city)
	fmt.Println("State: ")
	fmt.Scanln(&customer.state)
	fmt.Println("Zip Code: ")
	fmt.Scanln(&customer.zip)
	fmt.Println("User Name")
	fmt.Scanln(&customer.userName)
	fmt.Println("Password")
	fmt.Scanln(&customer.pass)
	return customer

}

func main() {
	const INCOMES = 40000
	var option int
	var userN string
	var passU string
	var opt2 string = "y"
	var anualIncomes float32
	var monExp float32
	var chAcc int
	var saAcc int
	var customer Customer

	for opt2 == "y" {
		ClientMenu()
		fmt.Println("What do you want to do?")
		fmt.Scanln(&option)

		switch option {
		case 1:
			if customer.userName == "" {
				customer = Register()
				fmt.Println("Congratulation your registration was succecful ")
				fmt.Println(" ")
			} else {
				fmt.Println("You already have an account")
			}
		case 2:
			fmt.Println("Please neter your user name")
			fmt.Scanln(&userN)
			if userN == customer.userName {
				fmt.Println("Please enter your password")
				fmt.Scanln(&passU)
				if passU == customer.pass {
					fmt.Println("PERSONAL INFORMATION")
					fmt.Println(" ")
					fmt.Println("First Name: " + customer.FirstName)
					fmt.Println("Last Name: " + customer.lastName)
					fmt.Println("Address: " + customer.street + ", " + customer.city + ", " + customer.state + " " + customer.zip)
					fmt.Println("Please enter your anual incomes")
					fmt.Scanln(&anualIncomes)
					fmt.Println("Please enter your mountly expenses")
					fmt.Scanln(&monExp)
					if anualIncomes > INCOMES {
						if (anualIncomes/12)/2 >= monExp {
							fmt.Println("Congratulations, your request has been accepted")
							fmt.Println("Your new checking account is:")
							chAcc = rand.Intn(1000000000000)
							fmt.Println(chAcc)
							fmt.Println("Your new savings account is ")
							saAcc = rand.Intn(1000000000000)
							fmt.Println(saAcc)
							//fmt.Println(chAcc)
						}
					} else {
						fmt.Println("We are sorry but you can't open an account")
					}
				} else {
					fmt.Println("Wrong Password")
				}
			} else {
				fmt.Println("Wrong Username")
			}
		case 3:
			fmt.Println("This is the option 3")
		case 4:
			fmt.Println("This is the option 4")
		case 5:
			fmt.Println("This is the option 5")
		case 6:
			fmt.Println("This is the option 6")
		}

		fmt.Println("Do you want to continune(y/n)")
		fmt.Scanln(&opt2)

	}

}
