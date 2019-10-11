package gen

import (
	"database/sql"
	"fmt"
)

/*
Generate is the main function inside the gen package. Really, Generate should
be the only function in the gen package to be called outside of the gen and the
gen_test package. Additionally, it should only be called once, at the start of
main.go inside the main package. Generate should be called to generate an index,
beverage type and a stock amount for each row in the vending machine. First the
table needs to be cleared of all data. Then, the "make" family of methods should
be called and return a slice. Generate should send those slices to be written
into the database table. Finally, Generate has a catch if it is sent a zero or
less. It does not make sense to have a vending machine with zero rows or zero
stock capacity.
*/
func Generate(db *sql.DB, rows int, columns int, max int) bool {
	r := false
	if (rows <= 0) || (columns <= 0) || (max <= 0) {
		fmt.Println("Error! Could not generate machine.")
		fmt.Println("One or more field has been left blank or made less than 1.")
	} else {
		// index := MakeIndex(rows, columns)
		beverage := MakeBeverage(db, rows, columns)
		// stock := MakeStock(rows, columns, max)
		// WriteTo(db, index, beverage, stock)
		if beverage != nil {
			r = true
		} else {
			r = false
		}
	}
	return r
}

// package gen

// import (
// 	"testing"
// )

// func TestGenerate(t *testing.T) {
// 	for c := 0; c < 3; c++ {
// 		r := Generate(c, c-1, c-2)
// 		if r == 1 {
// 			t.Errorf("Generate was passed 0(s) and slices were created.")
// 		} else {
// 			t.Log("Generate catches normally")
// 		}
// 	}
// }
