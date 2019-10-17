package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var db *sql.DB

func main() {
	var err error
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	mainmenu()
	//signup()
	// login()

}

func signup() {
	var username, fname, lname, password string
	fmt.Println("Sign up for an account")
	fmt.Println("enter your first name")
	fmt.Scanln(&fname)
	fmt.Println("enter your last name")
	fmt.Scanln(&lname)
	fmt.Println("enter a username")
	fmt.Scan(&username)

	fmt.Println("enter a password")
	fmt.Scan(&password)

	sql := `
	insert into employees (fname, lname, username, pass, reimbursement, currentstatus)
	values ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(sql, fname, lname, username, password, "he", "dadkls")
	if err != nil {
		fmt.Println("usrname already exists")
	} else {

		fmt.Println("Registration successful")

	}
}

func login() {
	var username string
	var password string
	var placeholder string
	fmt.Println("Login to your account")
	fmt.Println("enter your username")
	fmt.Scan(&username)
	fmt.Println("enter your password")
	fmt.Scan(&password)

	sql := `select pass from employees where username = $1`

	result := db.QueryRow(sql, username)

	result.Scan(&placeholder)

	if placeholder != password {
		fmt.Println("Invalid username/password")
		fmt.Println()
		login()
	} else {
		fmt.Println("Login Successful")

		fmt.Println("Employee Portal")
	}

}

func mainmenu() {
	var selection string
	fmt.Println("Welcome to the reimbursemnt app")
	fmt.Println("Enter 1 to sign up for an account")
	fmt.Println("Enter 2 login")
	fmt.Scan(&selection)

	switch selection {
	case "1":
		signup()
	case "2":
		login()

	default:
		fmt.Println("Invalid Option")
		mainmenu()

	}
}

/*

cd db
sudo docker build -t project .
sudo docker run --name mydb -d -p 5432:5432 project


reset dabatase
sudo docker stop mydb
sudo docker rm mydb
sudo docker rmi project


check database
sudo docker ps-a
sudo docker images -a



to start
go run *.go
go run *.go  -hide

*/
