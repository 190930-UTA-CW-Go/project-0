package gen

import (
	"database/sql"
	"fmt"
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func TestGenerate(t *testing.T) {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, _ := sql.Open("postgres", datasource)
	defer db.Close()
	var r bool

	for c := 0; c < 3; c++ {
		r = Generate(db, c, c-1, c-2)
		if r == true {
			t.Errorf("Generate was passed 0(s) and slices were created.")
		} else {
			t.Log("Generate catches normally")
		}
	}
}
