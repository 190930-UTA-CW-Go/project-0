package gen

import "math/rand"

func makeStock(rows int, columns int, max int) []int {
	StockAmount := make([]int, rows*columns)

	for i := 0; i < (rows * columns); i++ {
		StockAmount[i] = rand.Intn(max)
	}

	return StockAmount
}
