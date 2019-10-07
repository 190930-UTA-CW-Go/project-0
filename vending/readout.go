package vending

// This interface allows us to us Print instead of Println
type Readout interface {
	String() string
}