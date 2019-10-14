package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/Tony-Moon/project-0/base"
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

	r := gen.Generate(data, 5, 5, 10)
	if r == true {
		vend.Navigate()
	} else {
		fmt.Println("Failed to generate vending machine.")
	}

}
