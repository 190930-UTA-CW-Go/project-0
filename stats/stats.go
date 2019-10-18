package stats

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

//Phealth gets player health
func Phealth() int {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	name := name.GetName()
	var health int
	rows, err := db.Query("SELECT name, health FROM bankaccount where name = $1;", name)
	if err != nil {
		// handle err
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &health)
		if err != nil {
			log.Fatal(err)
		}
		//print to check if its working
		//fmt.Println(name, cfunds)

	}
	return health
}

//Pattack gets player attack
func Pattack() int {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	name := name.GetName()
	var attack int
	rows, err := db.Query("SELECT name, attack FROM bankaccount where name = $1;", name)
	if err != nil {
		// handle err
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &attack)
		if err != nil {
			log.Fatal(err)
		}
		//print to check if its working
		//fmt.Println(name, cfunds)

	}
	return attack
}
