package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/gittingdavid/project-0/database"
	"github.com/gittingdavid/project-0/menu"
	_ "github.com/lib/pq"
)

// Connection string information
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	// Connecting to database
	var err error
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	(database.DBCon), err = sql.Open("postgres", datasource)
	defer (database.DBCon).Close()
	if err != nil {
		panic(err)
	}

	//Set RNG Seed
	rand.Seed(time.Now().UTC().UnixNano())

	//Main Menu
	menu.Menu()
}
