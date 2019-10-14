package gen

/*
MakeStock generates a slice that wil tell the database how many of each drink is
in the row. It accomplishes this by creating a random integer, using UseSeed,
between 0 and a maximum specified by main.go in the main package. It chooses a
new integer for each element in the slice MakeStock returns.
*/
func MakeStock(rows int, columns int, max int) []int {
	StockAmount := make([]int, rows*columns)

	for i := 0; i < (rows * columns); i++ {
		StockAmount[i] = UseSeed(0, max)
	}

	return StockAmount
}
