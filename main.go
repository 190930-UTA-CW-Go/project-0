package main

import (
	_ "bufio"
	"fmt"
	_ "os"
)

func main() {
	fmt.Println("Welcome to my banking app")
	var custArr []customers
	//var arr [] customer.Customer
	cust1 := customers{"Dio", "Brando", "DIO", "ZAWARUDO", 100000.64}
	cust2 := customers{"Jotaro", "Kujo", "JOJO", "Star Platinum", 100.00}
	cust3 := customers{"Ren", "Amamiya", "persona", "Joker", 2000.25}
	cust4 := customers{"Peter", "Peng", "NAMVP", "Doublelift", 50000.50}
	cust5 := customers{"Zachary", "Scuderi", "C9Poggers", "C9Sneaky", 100000.64}
	cust6 := customers{"Mikey", "Xiong", "Leeshalove", "Ideals", 7500.25}
	custArr = append(custArr, cust1, cust2, cust3, cust4, cust5, cust6)

	//needed to look at how objects are being stored in employee
	// employ := employee{custArr}
	// employ.List()

	// testing register function
	// newcust := newcustomer{}
	// newcust.Register()

	//testing addcustomer() function
	cust1.addCustomer()

	//cust1.add(cust2)
	//fmt.Println(cust1.Amount())
	// cust1.Amount()
	// fmt.Println(custArr)

	// cust1.Deposit(500)
	// fmt.Println(cust1.Amount())
	// var tagList []customer.Customer
	// results := []customer.Customer{customer.Customer{Firstname: "Jeff", Lastname: "Bogard", Password: "1234", Username: "Jefe", Balance:1000.64}}
	// for _, details := range results{
	// 	tagList = append(tagList, customer.Customer{Firstname: details.Firstname, Lastname: details.Lastname,
	// 		 Password: details.Password, Username: details.Username, Balance: details.Balance})
	// 	}
	// 	fmt.Println("Customers: ", tagList)

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
}
