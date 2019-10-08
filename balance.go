package fund

type Balance struct {
	balance float64
}

func NewFund(i float64) *Fund {
	return &Fund{
		balance: i,
	}
}

// Balance returns current balance of a Fun
func (f *Fund) Balance() float64 {
	return f.balance
}

//Withdraw removes an amount from the current balance of a Fun
func (f *Fund) Withdraw(i float64) {
	f.balance -= i
}