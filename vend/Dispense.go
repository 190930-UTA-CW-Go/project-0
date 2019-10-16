package vend

import "database/sql"

/*
Dispense reduces the stock of the users selection by one. It recieves the database
information, the index of the users selection and the stock amount of that row. Then
it updates the **machine** table using index and then new stock amount.
*/
func Dispense(db *sql.DB, index string, stock int) {
	stock = stock - 1
	db.Exec("UPDATE machine SET stock = $1 WHERE index = $2;", stock, index)
}
