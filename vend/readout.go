package vend

// This interface allows us to us Print instead of Println
type Readout interface {
	String() string
}
