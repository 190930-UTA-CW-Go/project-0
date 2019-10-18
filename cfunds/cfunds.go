package cfunds

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/190930-UTA-CW-Go/project-0/name"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

//Cfunds tracks players on hand gold
func Cfunds() float64 {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	name := name.GetName()
	var cfunds float64
	rows, err := db.Query("SELECT name, cfunds FROM bankaccount where name = $1;", name)
	if err != nil {
		// handle err
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &cfunds)
		if err != nil {
			log.Fatal(err)
		}
		//print to check if its working
		//fmt.Println(name, cfunds)

	}
	return cfunds
}

//done
