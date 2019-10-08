package vend

import "strconv"

// Row - a structure for each row
type Row struct {
	Row      string
	Beverage string
	Stock    int
}

// String() method prints the readout of the row structure
func (p Row) String() string {
	return p.Row + " - " + p.Beverage + " - " + strconv.Itoa(p.Stock) + " in stock.\n"
}
