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
	}

	//db.Exec("INSERT INTO pokemon VALUES (4, 'Eevee')")
	//db.Exec("INSERT INTO pokemon VALUES (8, 'tyrannitar')")
	//	getAll(db)
	GetAll3(db)
	//searchByName(db, "Eeesevee")
	//employee.NewAcc()
	GetAll2(db)
	//employee.SearchUser("thirdacc")
	//employee.Welcome()
	//employee.NewAcc()
	////db.Exec("INSERT INTO employeeAccounts VALUES ('adf', 'Eeeevee')")
	//GetAll3(db)
	//fmt.Println("die")
	//SearchByName2(db, "password")
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

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func getAll(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM POKEMON")
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}

func searchByName(db *sql.DB, searchvalue string) {
	row := db.QueryRow("SELECT * FROM pokemon WHERE name = $1", searchvalue)
	var id int
	var name string
	row.Scan(&id, &name)
	fmt.Println(id, name)
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
	rows, _ := db.Query("SELECT * FROM TICKETS")
	for rows.Next() {
		var u1 int
		var u2 string
		var u3 string
		var u4 string
		var u5 float32
		var u6 string
		var u7 string
		rows.Scan(&u1, &u2, &u3, &u4, &u5, &u6, &u7)
		fmt.Println(u1, u2, u3, u4, u5, u6, u7)
	}
}
