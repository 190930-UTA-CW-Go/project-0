package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("Welcome to my banking app")
	var custArr []account
	//var arr [] customer.Customer
	cust1 := account{"Dio ", "Brando ", "DIO ", "ZAWARUDO ", 100000.64}
	cust2 := account{"Jotaro ", "Kujo ", "JOJO ", "Star Platinum ", 100.00}
	cust3 := account{"Ren ", "Amamiya ", "persona ", "Joker ", 2000.25}
	cust4 := account{"Peter ", "Peng ", "NAMVP ", "Doublelift ", 50000.50}
	cust5 := account{"Zachary ", "Scuderi ", "C9Poggers ", "C9Sneaky ", 100000.64}
	cust6 := account{"Mikey ", "Xiong ", "Leeshalove ", "Ideals ", 7500.25}
	custArr = append(custArr, cust1, cust2, cust3, cust4, cust5, cust6)

	// Testing Database
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	db.Exec("INSERT INTO account VALUES ('Dio', 'Brando', 'DIO','ZAWARUDO',5000.75)")
	getAll(db)
	// testing Writo to file
	// var customArr map[int]account
	// customArr = make(map[int]account)
	// customArr[1] = cust1
	// customArr[2] = cust2
	// WriteToFile(customArr)

	//needed to look at how objects are being stored in employee
	// employ := employee{custArr}
	// employ.List()

	// testing register function
	// newcust := newcustomer{}
	// newcust.Register()

	//testing addcustomer() function
	//cust1.addCustomer()

	//cust1.add(cust2)
	//fmt.Println(cust1.Amount())
	// cust1.Amount()
	// fmt.Println(custArr)

	// cust1.Deposit(500)
}
