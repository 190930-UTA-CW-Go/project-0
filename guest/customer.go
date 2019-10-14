package guest

import (
	"database/sql"
	"fmt"
	_ "log" //no
	_ "os"  //no
	_ "os/exec"
	_ "strconv"

	_ "github.com/lib/pq" // no
)

// Customer data
type Customer struct {
	userName string
	password string
	name     string
}

// NewAccGuest fdsf
func NewAccGuest() {
	var userName string
	var password string
	var fname string
	var lname string
	fmt.Printf("Creating a new account:")
	fmt.Printf("Enter a username  ")
	fmt.Scanln(&userName)
	fmt.Printf("Enter a password  ")
	fmt.Scanln(&password)
	fmt.Printf("Enter your first  ")
	fmt.Scanln(&fname)
	fmt.Printf("Enter your last  ")
	fmt.Scanln(&lname)
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	db.Exec("INSERT INTO customerLogin(userName,password,fname, lname)"+
		"VALUES($1,$2,$3, $4)", userName, password, fname, lname)
}

// NewCustomer is a Constructor for Customer
func NewCustomer(userName string, password string,
	name string) {
	n := Customer{
		userName: userName,
		password: password,
		name:     name,
	}

	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	db.Exec("INSERT INTO customer"+"(n.userName,n.password,n.name)"+
		"VALUES($1,$2,$3)", n.userName, n.password, n.name)

	//return &n
}

//SearchByName func
func SearchByName(userName string) {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	row := db.QueryRow("SELECT userName, password, fname, lname FROM customerLogin WHERE userName = $1", userName)
	var n1, n2, n3, n4 string
	//var isApproved bool
	row.Scan(&n1, &n2, &n3, &n4)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

}

//GetAll is a func
func GetAll(db *sql.DB) {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	rows, _ := db.Query("SELECT * FROM customerLogin")
	for rows.Next() {
		var userName, password, fname, lname string
		//var isApproved bool
		rows.Scan(&userName, &password, &fname, &lname)
		fmt.Println(userName, password, fname, lname)
	}
}

/*
func (a Customer) String() string {
	var output string
	t := fmt.Sprintf("%.2f", a.balance)
	output = a.userName + " | " + a.password + " | " + a.name + " | $" + t + "\n"
	return output
}

// Balance returns the amount of money in a customer's balance
func (a *Customer) Balance() float64 {
	return a.balance
}

// Withdraw removes money from a customer's balance
func (a *Customer) Withdraw(i float64) {
	if a.balance < i {
		fmt.Println("Insufficient funds, transaction canceled")
	} else {
		a.balance -= i
	}

}

// Deposit adds money to a customer's balance
func (a *Customer) Deposit(i float64) {
	a.balance += i
}

// Transfer moves money from one customer to another customer's balance
func (a *Customer) Transfer(i float64, b *Customer) {
	if a.balance < i {
		fmt.Println("Insufficient funds, transaction canceled")
	} else {
		a.Withdraw(i)
		b.balance += i
	}
}*/
