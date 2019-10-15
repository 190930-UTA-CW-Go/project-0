package app

import (
	"database/sql"
	"fmt"
)

// ListMachine lists all items in the machine table
func ListMachine(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM machine;")
	var ind, bev, bra string
	var sto int
	for rows.Next() {
		rows.Scan(&ind, &bev, &sto, &bra)
		fmt.Println(ind, bev, sto, bra)
	}
}
