package vend

import "database/sql"

/*
GetDrinks documentation
*/
func GetDrinks(db *sql.DB) ([]string, []string, []int, string) {
	rows, _ := db.Query("SELECT * FROM machine;")
	i := 0
	index := make([]string, 2)
	name := make([]string, 2)
	stock := make([]int, 2)
	var brand string

	for rows.Next() {
		rows.Scan(&index[i], &name[i], &stock[i], &brand)
		i++
	}

	return index, name, stock, brand
}
