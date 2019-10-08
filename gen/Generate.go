package gen

/*
Generate is the main function inside the gen package. Really, Generate should
be the only function in the gen package to be called outside of the gen and the
gen_test package. Additionally, it should only be called once, at the start of
main.go inside the main package.
*/
func Generate(rows int, columns int, max int) ([]string, []string, []int) {

	Index := MakeIndex(rows, columns)
	Beverage := MakeBeverage(rows, columns)
	Stock := MakeStock(rows, columns, max)

	return Index, Beverage, Stock

}
