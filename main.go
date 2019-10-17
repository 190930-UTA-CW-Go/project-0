package main

import (
	"database/sql"
	"fmt"
	"os"

	dbconnection "github.com/NGKlaure/project-0/dbConnection"
)

var newCustomer NewCustomer

func main() {

	newCustomer = NewCustomer{}
	fmt.Println("banking system is running")

	db := dbconnection.DbConnection()
	ping(db)
	Menu()

}

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected3!")
}

func Menu() {
	fmt.Println("select an option to perform")
	fmt.Print("==============================")
	fmt.Println("============================")
	fmt.Println("1:Register new customer")
	fmt.Println("")
	fmt.Println("2:Login")
	fmt.Println("")
	fmt.Println("3:Employee Login")
	fmt.Println("")
	fmt.Println("4: Exit")

	var choice string

	fmt.Scanln(&choice)
	switch choice {
	case "1":
		newCustomer.Register()
	case "2":
		newCustomer.login()
	case "3":
		ELogin()
	case "4":
		os.Exit(0)
	}

}
