package gentest

import (
	"fmt"
	"testing"

	"github.com/Tony-Moon/project-0/gen"
)

/*
TestGenerate is a unit test which checks the Generate function in the gen package. The
Generate function should not procede with generating a vending machine if one of its
parameters are less than one.
*/
func TestGenerate(t *testing.T) int {
	for c := 0; c < 3; c++ {
		i, b, s := gen.Generate(c, c-1, c-2)
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
	return 0
}
