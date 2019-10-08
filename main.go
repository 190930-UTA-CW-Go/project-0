package main

import (
	"fmt"
	"github.com/project-0/Customer"
	"bufio"
	"os"
	"strings"
)

func main() {
	var tagList []customer.Customer
	results := []customer.Customer{customer.Customer{Firstname: "Jeff", Lastname: "Bogard", Password: "1234", Username: "Jefe", Balance:1000.64}}
	for _, details := range results{
		tagList = append(tagList, customer.Customer{Firstname: details.Firstname, Lastname: details.Lastname,
			 Password: details.Password, Username: details.Username, Balance: details.Balance})
		}
		fmt.Println("Customers: ", tagList)
	

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Are you a customer or employee?")
	obj, _:= reader.ReadString('\n')
	if strings.TrimRight(obj, "\n") == "customer"{
		fmt.Println("Please enter a username")
		username, _:= reader.ReadString('\n')
		fmt.Println("Please enter a password")
		password, _:= reader.ReadString('\n')
		fmt.Println("Your username is:" + username, "Your password is: " + password)
		
	}	else{
		obj = "employee"
		fmt.Println("You are an " + obj)
	}
	

	var persona = customer.Customer{
		Firstname: "Ren",
		Lastname: "Amaiya",
		Password: "1234",
		Username: "Joker",
		Balance: 1000.50,
	}
	fmt.Println(persona)

	// var sum int = 0
	// for i :=0; i<5; i++{
	// 	sum += i
	// 	if sum > 0{
	// 		fmt.Println(sum)
	// 	}
	// }
	// fmt.Println(sum)
}