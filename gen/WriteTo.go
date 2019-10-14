package gen

import "database/sql"

// WriteTo documentation
func WriteTo(db *sql.DB, index []string, beverage []string, stock []int, br string) {
	var in, be string
	var st int

	for i := range index {
		in = index[i]
		be = beverage[i]
		st = stock[i]

		db.Exec("INSERT INTO machine VALUES ($1, $2, $3, $4);", in, be, st, br)
	}
}
