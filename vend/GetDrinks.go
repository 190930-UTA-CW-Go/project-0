package vend

import "database/sql"

/*
GetDrinks documentation
*/
func GetDrinks(db *sql.DB) ([]string, []string, []int, string) {
	var index, name, brand string
	var stock, i, count int

	rows0, _ := db.Query("SELECT COUNT(*) as count FROM machine;")
	for rows0.Next() {
		rows0.Scan(&count)
	}

	inDB := make([]string, count)
	naDB := make([]string, count)
	stDB := make([]int, count)

	rows, _ := db.Query("SELECT * FROM machine;")
	i = 0
	for rows.Next() {
		rows.Scan(&index, &name, &stock, &brand)

		inDB[i] = index
		naDB[i] = name
		stDB[i] = stock
		i++
	}

	return inDB, naDB, stDB, brand
}
