package vending

import (
	"math/rand"
	"strconv"
)

func GenerateMachine(rows int, columns int, max int) ([]string, []string, []int) {

	Index    := GenerateIndex(rows, columns)
	Beverage := GenerateBeverage(rows, columns)
	Stock    := GenerateStock(rows, columns, max)
	
	return Index, Beverage, Stock

}

// This function generates the index for each row.
func GenerateIndex(rows int, columns int) []string{
	RowIndex := make([]string, rows*columns)
	count    := 0
	let      := 1
	num      := 1

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

// This function incriments the character used for the index
func toCharStr(i int) string {
	return string('A' - 1 + i)
}

func GenerateBeverage(rows int, columns int) []string {
	BeverageList := make([]string, rows*columns)
	list := make([]string, 6) 
	list[0] = "Coke"
	list[1] = "Diet Coke"
	list[2] = "Sprite"
	list[3] = "Water"
	list[4] = "Mountain Dew"
	list[5] = "Green Tea"

	for i:=0; i<(rows*columns); i++ {
		BeverageList[i] = list[rand.Intn(5)]
	}

	return BeverageList
}

func GenerateStock(rows int, columns int, max int) []int {
	StockAmount := make([]int, rows*columns)

	for i:=0; i<(rows*columns); i++ {
		StockAmount[i] = rand.Intn(max)
	}

	return StockAmount
}