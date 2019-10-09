package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5431
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

	db.Exec("INSERT INTO customer VALUES (2, 'Ivysaury', 'pAsSwOrdd', 'Ivy', 'Saur', 3400)")
	getAll(db)
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
		var userID int
		var userName, password, firstName, lastName string
		var balance int
		//var isApproved bool
		rows.Scan(&userID, &userName, &password, &firstName,
			&lastName, &balance)
		fmt.Println(userID, userName, password, firstName,
			lastName, balance)
	}
}

func searchByName(db *sql.DB, searchvalue string) {
	row := db.QueryRow("SELECT * FROM customer WHERE name = $1", searchvalue)
	var userID int
	var userName, password, firstName, lastName string
	var balance int
	//var isApproved bool
	row.Scan(&userID, &userName, &password, &firstName,
		&lastName, &balance)
	fmt.Println(userID, userName, password, firstName,
		lastName, balance)
}
