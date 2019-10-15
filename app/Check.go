package app

import (
	"database/sql"
)

/*
Check documentation
*/
func Check(db *sql.DB, application *Tech) int {
	r := 1
	var acc, pas, com, fir, las string
	appAcc := application.account
	rows, _ := db.Query("SELECT * FROM servicers;")
	for rows.Next() {
		rows.Scan(&acc, &pas, &com, &fir, &las)
		if appAcc == acc {
			r = r - 1
		}
	}
	return r
}
