package gen

// Generate is the main function inside the gen package
func Generate(rows int, columns int, max int) ([]string, []string, []int) {

	Index := makeIndex(rows, columns)
	Beverage := makeBeverage(rows, columns)
	Stock := makeStock(rows, columns, max)

	return Index, Beverage, Stock

}
