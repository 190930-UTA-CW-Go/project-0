package gen

import (
	"database/sql"
	"fmt"
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func TestWriteTo(t *testing.T) {
	var index, beverage, brand string
	var stock int

	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, _ := sql.Open("postgres", datasource)
	defer db.Close()

	ind := []string{"A1", "A2", "B1", "B2"}
	bev := []string{"b1", "b2", "b3", "b4"}
	sto := []int{1, 2, 3, 4}
	brn := "Brand"
	exInd, exBev, exSto, exBrn, i := 0, 0, 0, 0, 0
	WriteTo(db, ind, bev, sto, brn)

	rows, _ := db.Query("SELECT * FROM machine")
	for rows.Next() {
		rows.Scan(&index, &beverage, stock, brand)
		if index != ind[i] {
			exInd++
		}
		if beverage != bev[i] {
			exBev++
		}
		if stock != sto[i] {
			exSto++
		}
		if brand != brn {
			exBrn++
		}
		i++
	}

	if exInd != 0 {
		t.Errorf("Incorrect Index")
	} else if exBev != 0 {
		t.Errorf("Incorrect Beverage")
	} else if exSto != 0 {
		t.Errorf("Incorrect Stock")
	} else if exBrn != 0 {
		t.Errorf("Incorrect Brand")
	} else {
		fmt.Println("WriteTo works properly")
	}

}
