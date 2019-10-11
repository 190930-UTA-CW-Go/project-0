package gen

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	for c := 0; c < 3; c++ {
		r := Generate(c, c-1, c-2)
		if r == 1 {
			t.Errorf("Generate was passed 0(s) and slices were created.")
		} else {
			t.Log("Generate catches normally")
		}
	}
}
