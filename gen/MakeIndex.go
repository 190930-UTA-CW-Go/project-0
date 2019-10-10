package gen

import "strconv"

/*
MakeIndex generates an index key for the database. The key is a letter
followed by a number (such as A3). A vending machine can have multiple
rows (denoted by the letter) and multiple columns (denoted by the number).
main.go in the main package specifies how many rows and columns the vending
machine will have. The index will always start at "A1."
*/
func MakeIndex(rows int, columns int) []string {
	RowIndex := make([]string, rows*columns)
	count := 0
	let := 1
	num := 1

	for i := 0; i < rows; i++ {
		for n := 0; n < columns; n++ {
			RowIndex[count] = ToCharStr(let) + strconv.Itoa(num)
			num++
			count++
		}
		num = 1
		let++
	}
	return RowIndex
}

// ToCharStr increments the letter for each new row.
func ToCharStr(i int) string {
	return string('A' - 1 + i)
}
