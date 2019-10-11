package gen

import "database/sql"

// WriteTo documentation
func WriteTo(db *sql.DB, index []string, beverage []string, stock []int) {
	for i := range index {
		in := index[i]
		be := beverage[i]
		st := stock[i]

		db.Exec("INSERT INTO machine (%d, %s, %d)", in, be, st)
	}
}
