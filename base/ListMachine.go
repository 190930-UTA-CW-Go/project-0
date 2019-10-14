package db

import (
	"database/sql"
	"fmt"
)

// ListMachine lists all items in the machine table
func ListMachine(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM machine")
	var id int
	var name, brand string
	var prob float64
	for rows.Next() {
		rows.Scan(&id, &name, &brand, &prob)
		fmt.Println(id, name, brand, prob)
	}
}
