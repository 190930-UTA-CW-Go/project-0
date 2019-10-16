package employee

import (
	"database/sql"
	_ "database/sql" //why
	"fmt"
	_ "fmt" //wont

	_ "github.com/lib/pq" // u save
)

//OPEN1 opens database; needed to clean up code
func OPEN1() *sql.DB {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	return db
}
