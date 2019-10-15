package app

import (
	"database/sql"
)

/*
WriteTo documentation
*/
func WriteTo(db *sql.DB, application *Tech) {
	var acc, pas, com, fir, las string
	acc = application.account
	pas = application.password
	com = application.company
	fir = application.firstname
	las = application.lastname
	db.Exec("INSERT INTO servicers VALUES ($1, $2, $3, $4, $5);",
		acc, pas, com, fir, las)
}
