package vend

import "database/sql"

/*
Refill Documentation
*/
func Refill(db *sql.DB, max int) {
	db.Exec("UPDATE machine SET stock = $1;", max)
}
