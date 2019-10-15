package app

import (
	"database/sql"
	"fmt"
)

// ListDrinks lists all items in the drinklist table
func ListDrinks(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM drinklist;")
	var index int
	var name, brand string
	var prob float64
	for rows.Next() {
		rows.Scan(&index, &name, &brand, &prob)
		fmt.Println(index, name, brand, prob)
	}
}
