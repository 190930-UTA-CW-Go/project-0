package db

import (
	"database/sql"
	"fmt"
)

// ListDrinks lists all items in the drinklist table
func ListDrinks(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM drinklist")
	var id int
	var name, brand string
	var prob float64
	for rows.Next() {
		rows.Scan(&id, &name, &brand, &prob)
		fmt.Println(id, name, brand, prob)
	}
}
