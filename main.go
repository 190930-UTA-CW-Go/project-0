package main

import (
	"database/sql"
	"fmt"

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

	ping(db)

}

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func showUserTable(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM user_accounts")
	for rows.Next() {
		var uniqname string
		var userfirst string
		var userlast string
		var password string
		var funds float32

		rows.Scan(&uniqname, &userfirst, &userlast, &password, &funds)
		fmt.Println(uniqname, userfirst, userlast, funds)
	}
}
