package main

import (

	//"database/sql"

	//"database/sql"
	"database/sql"
	"fmt"
	"os"

	dbconnection "github.com/NGKlaure/project-0/dbConnection"
	//_ "github.com/lib/pq"
)

var newCustomer NewCustomer

//var emplyee Employee

func main() {

	newCustomer = NewCustomer{}
	fmt.Println("banking system is running")

	db := dbconnection.DbConnection()
	ping(db)
	Menu()
	//db.Exec("INSERT INTO newCustomer VALUES (7, 'Vevee','1234')")
	//db.Exec("INSERT INTO newCustomer values (6, 'Fevee','2345')")
	//db.Exec("INSERT INTO account VALUES ('1234567','nad',34,'checking',123.00)")
	//db.Exec("delete from newCustomer ")
	//getAll(db)
	//getAllAcc(db)

	//getAllJointAcc(db)
	//searchByName(db, "nadine")

	//emplyee = Employee{}
	//ELogin()

	//newCustomer.Register()
	//newCustomer.login()
	//newCustomer.CreateNewAccount()
	//newCustomer.Applyforjoint()
	//newCustomer.Withdraw()
	//newCustomer.Deposit()
	//getAllAcc(db)
	//getAllJointAcc(db)
	//getAllEmp(db)

}

/* func getAll(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM newCustomer")
	defer db.Close()
	for rows.Next() {

		var userName string
		var password string
		rows.Scan(&userName, &password)
		fmt.Println(userName, password)
	}
} */

/* func getAllEmp(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM employee")
	defer db.Close()
	for rows.Next() {

		var userName string
		var password string
		rows.Scan(&userName, &password)
		fmt.Println(userName, password)
	}
} */

/* func searchByName(db *sql.DB, searchvalue string) {
	row := db.QueryRow("SELECT * FROM newCustomer WHERE name = $1", searchvalue)

	var name string
	var password string
	row.Scan(&name, &password)
	fmt.Println(name, password)
} */

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected3!")
}

/* func getAllAcc(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM account")
	defer db.Close()
	for rows.Next() {
		var accountNum string
		var custName string
		var age int
		var accountType string
		var availableBal float64
		rows.Scan(&accountNum, &custName, &age, &accountType, &availableBal)
		fmt.Println(accountNum, custName, age, accountType, availableBal)
	}
} */

/* func getAllJointAcc(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM jointAccount")
	defer db.Close()
	for rows.Next() {
		var id int
		var name1 string
		var name2 string
		var accType string
		var availableBal float64
		rows.Scan(&id, &name1, &name2, &accType, &availableBal)
		fmt.Println(id, name1, name2, accType, availableBal)
	}
} */

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
