package main

import (
	"database/sql"
	"fmt"
	_ "project-0/guest"
	_ "strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	//db.Exec("INSERT INTO pokemon VALUES (6, 'Eeeevee')")

	db.Exec("INSERT INTO customer VALUES ('Ivysaury', 'pAsSwOrdd', 'Saur', 3400)")
	getAll(db)

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

func getAll(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM CUSTOMER")
	for rows.Next() {
		var userName, password, name string
		var balance float64
		//var isApproved bool
		rows.Scan(&userName, &password, &name, &balance)
		fmt.Println(userName, password, name)
		fmt.Println(balance)
	}
}

func searchByName(db *sql.DB, searchvalue string) {
	row := db.QueryRow("SELECT * FROM customer WHERE name = $1", searchvalue)
	var userName, password, name string
	var balance float64
	//var isApproved bool
	row.Scan(&userName, &password, &name, &balance)
	fmt.Println(userName, password, name, balance)
}
