package db

import (
	"database/sql"
	"fmt"
)

// List lists all itesm in the vending machine
func List(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM machine")
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name) // Scan the row and paste it into id and name
		fmt.Println(id, name)
	}
}
