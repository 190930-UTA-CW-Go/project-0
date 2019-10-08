package gen

import "math/rand"

/*
MakeStock generates a slice that wil tell the database how many
of each drink is in the row. It does by creating a random integer
between 0 and a maximum specified by main.go in the main package.
*/
func MakeStock(rows int, columns int, max int) []int {
	StockAmount := make([]int, rows*columns)

	for i := 0; i < (rows * columns); i++ {
		StockAmount[i] = rand.Intn(max)
	}

	return StockAmount
}
