package gen

import "database/sql"

// WriteTo documentation
func WriteTo(db *sql.DB, index []string, beverage []string, stock []int, br string) {
	for i := range index {
		in := index[i]
		be := beverage[i]
		st := stock[i]

		db.Exec("INSERT INTO machine VALUES (%s, %s, %d, %s);", in, be, st, br)
	}
}
