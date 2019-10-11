package gen

import (
	"database/sql"
	"fmt"
)

/*
MakeBeverage generates a slice that wil tell the database which
type of drink goes where.
*/
func MakeBeverage(db *sql.DB, rows int, columns int) []string {
	BeverageList := make([]string, rows*columns)
	tablerows, _ := db.Query("SELECT * FROM drinklist")
	list := make([]string, 11)

	brand := UseSeed(0, 2)
	var id, x, i, sta, sto int
	var name, tablebrand string
	var prob float64

	switch brand {
	case 0:
		sta = 0
	case 1:
		sta = 11
	case 2:
		sta = 22
	default:
		sta = 0
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
	fmt.Println(BeverageList)
	return BeverageList
}
