package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/Tony-Moon/project-0/app"
	"github.com/Tony-Moon/project-0/gen"
	"github.com/Tony-Moon/project-0/vend"
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
	data, err := sql.Open("postgres", datasource)
	defer data.Close()

	if err != nil {
		log.Println(err)
		panic(err)
	}

	row := flag.Int("row", 3, "number of rows in vending machine")
	col := flag.Int("col", 5, "number of columns in vending machine")
	cap := flag.Int("cap", 10, "number of capacity each slot in vending machine has")
	apply := flag.String("apply", "none", "apply to be technician, follow with d , s or t ")
	flag.Parse()

	if *apply != "none" {
		app.Apply(data, *apply)
	}

	r := gen.Generate(data, *row, *col, *cap)
	if r == true {
		vend.Navigate(data)
	} else {
		fmt.Println("Failed to generate vending machine.")
	}

}
