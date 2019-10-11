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

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected5!")
}

var newCustomer NewCustomer

func main() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	//ping(db)
	//db.Exec("INSERT INTO newCustomer VALUES (1, 'Vevee','1234')")
	//db.Exec("INSERT INTO account VALUES (2, 'Fevee')")
	//getAll(db)

	fmt.Println("banking system is running")

	//a := New("mmm", "nad", "234", "nadine", "1132", 233.23)
	//c := New1("mimi", "1234",)

	newCustomer = NewCustomer{}
	newCustomer.Register()
	//newCustomer.login()

}

func getAll(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM newCustomer")
	for rows.Next() {
		var id int
		var userName string
		var password string
		rows.Scan(&id, &userName, &password)
		fmt.Println(id, userName, password)
	}
}
