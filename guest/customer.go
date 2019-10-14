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
	var name string
	fmt.Printf("Creating a new account:")
	fmt.Printf("Enter a username  ")
	fmt.Scanln(&userName)
	fmt.Printf("Enter a password  ")
	fmt.Scanln(password)
	fmt.Printf("Enter your full name  ")
	fmt.Scanln(name)
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	db.Exec("INSERT INTO customerLogin(userName,password,name)"+
		"VALUES($1,$2,$3)", userName, password, name)
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
	row := db.QueryRow("SELECT userName FROM customer WHERE userName = $1", userName)
	var password, name string
	//var isApproved bool
	row.Scan(&userName, &password, &name)
	fmt.Println(userName, password, name)

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
