package vend

import "database/sql"

/*
Refill recieves the database information and maximum capacity of
each row. Then, Refill updates the **machine** table so the entire
stock column is set to the maximum capcity.
*/
func Refill(db *sql.DB, max int) {
	db.Exec("UPDATE machine SET stock = $1;", max)
}
