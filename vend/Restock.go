package vend

// Restock adds beverages to the stock
func (r *Row) Restock(amount int) int {
	r.Stock = r.Stock + amount
	return r.Stock
}
