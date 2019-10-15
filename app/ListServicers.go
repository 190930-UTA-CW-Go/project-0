package app

import (
	"database/sql"
	"fmt"
)

// ListServicers lists all items in the servicers table
func ListServicers(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM servicers;")
	var acc, pas, com, fir, las string
	for rows.Next() {
		rows.Scan(&acc, &pas, &com, &fir, &las)
		fmt.Println(acc, pas, com, fir, las)
	}
}
