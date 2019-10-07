package vending

import (
	//"math"
	"strconv"
)

func Generate(rows int, columns int) []string{
	RowIndex := make([]string, rows*columns)
	count := 0
	let := 1
	num := 1

	for i:=0; i<rows; i++ {
		for n:=0; n<columns; n++ {
			RowIndex[count] = toCharStr(let) + strconv.Itoa(num)
			num++
			count++
		}
		num = 1
		let++
	}
	return RowIndex
}

func toCharStr(i int) string {
	return string('A' - 1 + i)
}