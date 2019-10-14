package main

import (
	"database/sql"
	"fmt"
	"project-0/employee"
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
	
	/*db.Exec("INSERT INTO pokemon VALUES (6, 'Eeeevee')")
	employee.GetAll2(db)
	employee.SearchByName2(db, "adfh")
	db.Exec("INSERT INTO customerLogin VALUES ('fffff', 'passwords', 'adsffff', 'miakhdddddalifa')")
	employee.SearchByName2(db, "passwords")
	employee.NewAccGuest()
	employee.GetAll2(db)*/
}
