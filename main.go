package main

import (
	"database/sql"
	"fmt"
	_ "os"
	"project-0/guest"
	_ "project-0/guest"
	_ "strconv"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Welcome to Banking app.")
	fmt.Println("Press number:")
	fmt.Println("1: Log on")
	fmt.Println("2: Create Account")
	fmt.Println("3. Exit")
	/*var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("1: Guest")
		fmt.Println("2: Employee")
	case 2:
		fmt.Println("Creating a new account:")
	}*/

	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	//db.Exec("INSERT INTO pokemon VALUES (6, 'Eeeevee')")

	db.Exec("INSERT INTO customerLogin VALUES ('Ivysaury', 'pAsSwOrdd', 'bulba','Saur')")
	//getAll(db)
	//guest.NewCustomer("garner1", "pass", "vgarneDeng")
	//guest.SearchByName("garner1")
	guest.SearchByName("Bulbasaury")
	//guest.SearchByName("adfh")
	//guest.NewAccGuest()
	guest.SearchByName("adfh")
	//guest.GetAll(db)

	//getAll(db)
	/*var customer1 = guest.NewCustomer("ggarnerdeng", "badpassword", "Garner Deng",
		999.99)
	fmt.Println(customer1)
	fmt.Println(customer1.Balance())
	customer1.Withdraw(99)
	fmt.Println(customer1.Balance())
	customer1.Deposit(1.01)
	fmt.Println(customer1.Balance())
	fmt.Println(customer1)
	var customer2 = guest.NewCustomer("wat", "badpassword", "wgat Deng",
		0)
	fmt.Println(customer2)

	customer1.Transfer(5, customer2)
	fmt.Println(customer2.Balance())
	fmt.Println(customer2)

	customer2.Withdraw(100)
	fmt.Println(customer2)*/
}

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

/*func searchByName(db *sql.DB, searchvalue string) {
	row := db.QueryRow("SELECT * FROM customer WHERE name = $1", searchvalue)
	var userName, password, name string
	var balance float64
	//var isApproved bool
	row.Scan(&userName, &password, &name, &balance)
	fmt.Println(userName, password, name, balance)
}*/

/*
create table customerAccounts(
    routingNumber serial primary key,
    balance float,
    userName varchar,
    userName2 varchar
)
insert into customerAccounts values (123123, 500, 'Bulbasaury', 'adfh' );


create table employeeAccounts(
    adminLogin varchar primary key,
    password varchar NOT NULL
)
insert into employeeAccounts values ('god', 'password');*/
