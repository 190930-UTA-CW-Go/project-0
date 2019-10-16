package vend

import "database/sql"

/*
Dispense documentation
*/
func Dispense(db *sql.DB, index string, stock int) {
	stock = stock - 1
	db.Exec("UPDATE machine SET stock = $1 WHERE index = $2;", stock, index)
}
