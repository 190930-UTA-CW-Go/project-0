package gen

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	for c := 0; c < 3; c++ {
		i, b, s := Generate(c, c-1, c-2)
		if i != nil {
			fmt.Println("Generate was passed 0(s) and Index was created.")
		}
		if b != nil {
			fmt.Println("Generate was passed 0(s) and Beverage was created.")
		}
		if s != nil {
			fmt.Println("Generate was passed 0(s) and Stock was created.")
		}
	}
}
