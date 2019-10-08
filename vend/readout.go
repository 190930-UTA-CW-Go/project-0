package vend

// Readout interface allows us to us Print instead of Println insdie the vend package
type Readout interface {
	String() string
}
