package main

import (
	"database/sql"
	"fmt"

	"github.com/Tony-Moon/project-0/db"
	"github.com/Tony-Moon/project-0/gen"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "usr"
	password = "pass"
	dbname   = "projcet0db"
)

func main() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	daba, err := sql.Open("postgres", datasource)
	defer daba.Close()

	if err != nil {
		panic(err)
	}

	gen.Generate(daba, 5, 5, 10)
	db.List(daba)
}
