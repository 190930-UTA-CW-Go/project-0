package main

import (
	"fmt"
	"github.com/project-0/Customer"
	_"bufio"
	_"os"
	_"strings"
	_"encoding/gob"
)

func main() {
	fmt.Println("Welcome to my banking app")
	cust1 := customer.Customer{"Dio", "Brando", "DIO","ZAWARUDO", 100000.64}
	cust2 := customer.Customer{"Jotaro", "Kujo", "JOJO","Star Platinum", 100.00}
	cust3 := customer.Customer{"Ren", "Amamiya", "persona","Joker", 2000.25}
	cust4 := customer.Customer{"Peter", "Peng", "NAMVP","Doublelift", 50000.50}
	cust5 := customer.Customer{"Zachary", "Scuderi", "C9Poggers","C9Sneaky", 100000.64}
	cust6 := customer.Customer{"Mikey", "Xiong", "Leeshalove","Ideals", 7500.25}

	cust1.Withdraw(1000000)
	fmt.Println(cust1.Amount())
	// var tagList []customer.Customer
	// results := []customer.Customer{customer.Customer{Firstname: "Jeff", Lastname: "Bogard", Password: "1234", Username: "Jefe", Balance:1000.64}}
	// for _, details := range results{
	// 	tagList = append(tagList, customer.Customer{Firstname: details.Firstname, Lastname: details.Lastname,
	// 		 Password: details.Password, Username: details.Username, Balance: details.Balance})
	// 	}
	// 	fmt.Println("Customers: ", tagList)

		// //create a file
		// dataFile, err := os.Create("Customer.gob")

		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }
		
		// // serialize the data
		// dataEncoder := gob.NewEncoder(dataFile)
		// dataEncoder.Encode(results)
		// fmt.Println("inside datafile", results)
		// dataFile.Close()

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Are you a customer or employee?")
	// obj, _:= reader.ReadString('\n')
	// if strings.TrimRight(obj, "\n") == "customer"{
	// 	fmt.Println("Please enter a username")
	// 	username, _:= reader.ReadString('\n')
	// 	fmt.Println("Please enter a password")
	// 	password, _:= reader.ReadString('\n')
	// 	fmt.Println("Your username is:" + username, "Your password is: " + password)
		
	// }	else{
	// 	obj = "employee"
	// 	fmt.Println("You are an " + obj)
	// }
	

	// var persona = customer.Customer{
	// 	Firstname: "Ren",
	// 	Lastname: "Amaiya",
	// 	Password: "1234",
	// 	Username: "Joker",
	// 	Balance: 1000.50,
	// }
	// fmt.Println(persona)

	// var sum int = 0
	// for i :=0; i<5; i++{
	// 	sum += i
	// 	if sum > 0{
	// 		fmt.Println(sum)
	// 	}
	// }
	// fmt.Println(sum)
}