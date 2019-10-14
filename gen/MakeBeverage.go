package gen

import (
	"database/sql"
	"sort"
)

/*
MakeBeverage generates a slice that denotes which beverage goes where. To accomplish this,
It first pulls all possible beverages from the *drinklist* table. The *drinklist* table is
seporated into three brands. To mimic real life, only one company can own a vending machine.
MakeBeverage uses UseSeed to pick a brand at random. Then, MakeBeverage will pull all drinks
made by that brand. Drinks have varying popularity in the real world. To reflect that, each
beverage has a "weight." The higher the weight, the more likely that beverage will make it
to the list. MakeBeverage uses UseSeed in a loop to randomly choose a beverage for each element
in the slice. The slice is then ordered alphabetically so all of the same kind of drinks are
together. MakeBeverage returns the ordered slice and brand chosen for the vending machine.
*/
func MakeBeverage(db *sql.DB, rows int, columns int) ([]string, string) {
	BeverageList := make([]string, rows*columns)
	tablerows, _ := db.Query("SELECT * FROM drinklist")
	list := make([]string, 11)
	var brandName string

	brand := UseSeed(0, 2)
	var id, x, i, sta, sto int
	var name, tablebrand string
	var prob float64

	switch brand {
	case 0:
		sta = 0
		brandName = "Dud-Cola"
	case 1:
		sta = 11
		brandName = "Salt-PhD"
	case 2:
		sta = 22
		brandName = "TipsyCo"
	default:
		sta = 0
		brandName = "Dud-Cola"
	}
	sto = sta + 10

	x = 0
	i = 0
	for tablerows.Next() {
		tablerows.Scan(&id, &name, &tablebrand, &prob)
		if (x >= sta) && (x <= sto) {
			list[i] = name
			i++
		}
		x++
	}

	for i = range BeverageList {
		switch x = UseSeed(0, 35); {
		case x == 0:
			BeverageList[i] = list[10]
		case x <= 2:
			BeverageList[i] = list[9]
		case x <= 4:
			BeverageList[i] = list[8]
		case x <= 7:
			BeverageList[i] = list[7]
		case x <= 10:
			BeverageList[i] = list[6]
		case x <= 13:
			BeverageList[i] = list[5]
		case x <= 17:
			BeverageList[i] = list[4]
		case x <= 21:
			BeverageList[i] = list[3]
		case x <= 25:
			BeverageList[i] = list[2]
		case x <= 30:
			BeverageList[i] = list[1]
		default:
			BeverageList[i] = list[0]
		}
	}
	sort.Strings(BeverageList)
	return BeverageList, brandName
}
