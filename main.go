package main

import (
	"database/sql"
	"fmt"
	_ "project-0/employee"

	_ "github.com/lib/pq"
)

func main() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
		fmt.Println("die")
	} else {
		fmt.Println("live")
	}
	//employee.NewAcc()
	db.Exec("INSERT INTO employeeAccounts VALUES ('adf', 'Eeeevee')")
	GetAll3(db)
	fmt.Println("die")
	SearchByName2(db, "password")
	//employee.Welcome()
	//GetAll2(db)
	/*db.Exec("INSERT INTO pokemon VALUES (6, 'Eeeevee')")
	employee.GetAll2(db)
	employee.SearchByName2(db, "adfh")
	db.Exec("INSERT INTO customerLogin VALUES ('fffff', 'passwords', 'adsffff', 'miakhdddddalifa')")
	employee.SearchByName2(db, "passwords")
	employee.NewAcc()
	employee.GetAll2(db)*/
}

//GetAll2 comment
func GetAll2(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM EMPLOYEELOGIN")
	for rows.Next() {
		var u1 int
		var u2 string
		var u3 string
		var u4 string
		var u5 string
		rows.Scan(&u1, &u2, &u3, &u4, &u5)
		fmt.Println(u1, u2, u3, u4, u5)
	}
}

//GetAll3 suck
func GetAll3(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM employeeAccounts")
	for rows.Next() {

		var adminLogin string
		var password string

		rows.Scan(&adminLogin, &password)
		fmt.Println(adminLogin, password)
	}
}

//SearchByName2 d
func SearchByName2(db *sql.DB, searchvalue string) {
	row := db.QueryRow("SELECT * FROM employeeAccounts WHERE userName = $1", searchvalue)
	var u1 string
	var u2 string

	row.Scan(&u1, &u2)
	fmt.Println(u1, u2)
}
