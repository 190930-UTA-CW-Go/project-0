package main

import (
	"fmt"
	"strings"

	//links to bank.go

	"database/sql"

	//imports the db

	"github.com/190930-UTA-CW-Go/project-0/arena"
	"github.com/190930-UTA-CW-Go/project-0/bank"
	"github.com/190930-UTA-CW-Go/project-0/cfunds"
	"github.com/190930-UTA-CW-Go/project-0/name"
	"github.com/190930-UTA-CW-Go/project-0/shop"

	//"github.com/190930-UTA-CW-Go/project-0/navigation"
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
	//gets data from db
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	//check table
	//getAll(db)

	name.GetName()
	bank.Bank()
	//endless loop
	for {
		Nav()
		//getAll(db)
	}
}

//Nav allows the player to choose their destination
func Nav() string {
	name := name.GetName()
	var tolocation string
	cfunds := cfunds.Cfunds()
	fmt.Println("Currently holding", cfunds, "Gold")
	fmt.Println("Where would you like to go", name, "? \n[S]hop [A]rena [B]ank")
	fmt.Scanln(&tolocation)
	var tolocationlower = strings.ToLower(tolocation)

	switch tolocationlower {
	case "s":
		shop.Shop()
	case "a":
		arena.Arena()
	case "b":
		bank.Bank()
	default:
		fmt.Println("Sorry please choose a valid destination")
	}
	return "nope" //Apparently I need a return but don't really want to use it
}

//getAll
func getAll(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM bankaccount")
	for rows.Next() {
		var name string
		var funds float64
		var cfunds float64
		var health int
		var attack int
		rows.Scan(&name, &funds, &cfunds, &health, &attack)
		fmt.Println(name, funds, cfunds, health, attack)
	}
}

//done
