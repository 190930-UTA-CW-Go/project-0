package vending

import "strconv"

// Inventory about individual rows
type Row struct {
	Row      string 
	Beverage string
	Stock    int
}

// This function dispenses one beverage
func (d *Row) Dispense() int{
	d.Stock = d.Stock - 1
	return d.Stock
}

// This function restocks the beverage
func (r *Row) Restock(amount int) int{
	r.Stock = r.Stock + amount
	return r.Stock
}

// Print the readout of the row and its stock
func (p Row) String() string{
	return p.Row + " - " + p.Beverage +  " - " + strconv.Itoa(p.Stock) + " in stock.\n"
}