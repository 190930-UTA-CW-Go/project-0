package gen

import (
	"testing"
)

func TestMakeStock(t *testing.T) {
	x := MakeStock(2, 2, 2)
	for i := range x {
		if (x[i] > 2) || (x[i] < 0) {
			t.Errorf("Stock out of range")
		} else {
			t.Log("Stock works normally")
		}
	}
}
