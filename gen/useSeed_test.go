package gen

import (
	"testing"
)

func TestUseSeed(t *testing.T) {
	x := make([]int, 4)
	for i := 0; i <= 3; i++ {
		x[i] = UseSeed(0, 100)
	}
	if (x[0] == x[1]) && (x[0] == x[2]) && (x[0] == x[3]) {
		t.Errorf("UseSeed is not generating random numbers.")
	} else {
		t.Log("UseSeed works normally")
	}
}
