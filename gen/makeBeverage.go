package gen

import "math/rand"

func makeBeverage(rows int, columns int) []string {
	BeverageList := make([]string, rows*columns)
	list := make([]string, 6)
	list[0] = "Coke"
	list[1] = "Diet Coke"
	list[2] = "Sprite"
	list[3] = "Water"
	list[4] = "Mountain Dew"
	list[5] = "Green Tea"

	for i := 0; i < (rows * columns); i++ {
		BeverageList[i] = list[rand.Intn(5)]
	}

	return BeverageList
}
