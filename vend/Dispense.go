package vend

// Dispense will reduce stock by 1
func (d *Row) Dispense() int {
	d.Stock = d.Stock - 1
	return d.Stock
}
