package main

import (

	//"database/sql"

	//"database/sql"
	"database/sql"
	"fmt"

	dbconnection "github.com/NGKlaure/project-0/dbConnection"
	//_ "github.com/lib/pq"
)

var newCustomer NewCustomer

func main() {
	fmt.Println("banking system is running")

	db := dbconnection.DbConnection()
	ping(db)
	//db.Exec("INSERT INTO newCustomer VALUES (7, 'Vevee','1234')")
	//db.Exec("INSERT INTO newCustomer values (6, 'Fevee','2345')")
	//db.Exec("INSERT INTO account VALUES ('1234567','nad','checking',123.00)")
	//db.Exec("delete from newCustomer ")
	//getAll(db)
	//getAllAcc(db)
	//searchByName(db, "nadine")

	newCustomer = NewCustomer{}
	//newCustomer.Register()
	newCustomer.login()
	//newCustomer.CreateNewAccount()
	//newCustomer.Withdraw()
	//newCustomer.Deposit()
	getAllAcc(db)

}

func getAll(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM newCustomer")
	defer db.Close()
	for rows.Next() {

		var userName string
		var password string
		rows.Scan(&userName, &password)
		fmt.Println(userName, password)
	}
}

func searchByName(db *sql.DB, searchvalue string) {
	row := db.QueryRow("SELECT * FROM newCustomer WHERE name = $1", searchvalue)

	var name string
	var password string
	row.Scan(&name, &password)
	fmt.Println(name, password)
}

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected3!")
}

func getAllAcc(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM account")
	defer db.Close()
	for rows.Next() {
		var accountNum string
		var custName string
		var accountType string
		var availableBal float64
		rows.Scan(&accountNum, &custName, &accountType, &availableBal)
		fmt.Println(accountNum, custName, accountType, availableBal)
	}
}
