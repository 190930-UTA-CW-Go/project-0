package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

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

	_, err := db.Exec(sql, fname, lname, username, password, "", "")
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
		employeePortal(username)

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
		mainmenu()
	case "2":
		login()
	case "manager":
		Manager()

	default:
		fmt.Println("Invalid Option")
		mainmenu()

	}
}

func employeePortal(username string) {
	var selection string
	var reimbursement string
	var placeholder string
	fmt.Println("Welcome to the employee portal")
	fmt.Println("Enter 1 to submit a reimbursement")
	fmt.Println("Enter 2 to view all reimbursements")
	fmt.Println("Enter 3 to logout")

	fmt.Scan(&selection)

	switch selection {
	case "1":
		fmt.Println("Submit a reimbursement")
		reader := bufio.NewReader(os.Stdin)
		reimbursement, _ = reader.ReadString('\n')
		reimbursement = strings.TrimSuffix(reimbursement, "\n")

		//fmt.Scan(&reimbursement)

		sql := `update employees
		set reimbursement =$1, currentstatus=$2
		where username =$3`

		_, err := db.Exec(sql, reimbursement, "Pending", username)
		if err != nil {
			panic(err)
		}
		fmt.Println("")
		fmt.Println("Reimbursement submitted")
		employeePortal(username)

	case "2":
		sql := `select reimbursement from employees where username =$1`

		result := db.QueryRow(sql, username)
		result.Scan(&placeholder)
		fmt.Println(placeholder)
		employeePortal(username)

	case "3":
		fmt.Println("Logout Successful")

		fmt.Println("")
		fmt.Println("")

		mainmenu()

	default:
		fmt.Println("Invalid Option")
		employeePortal(username)

	}
}

func Manager() {
	var selection string
	fmt.Println("Welcome Manager")
	fmt.Println("Enter 1 to view all pending reimbursements")
	fmt.Println("")
	fmt.Scan(&selection)

	switch selection {
	case "1":

		// rows, _ := db.Query("SELECT fname, ")

		// rows, _ := db.Query("SELECT  username, reimbursements, currentstatus from employees")
		rows, _ := db.Query("Select * FROM employees order by id asc")
		for rows.Next() {
			var id int
			var username string
			var fname string
			var lname string
			var pass string
			var reimbursements string
			var currentstatus string
			rows.Scan(&id, &fname, &lname, &username, &pass, &reimbursements, &currentstatus)

			fmt.Printf("%-10v", id)
			fmt.Printf("%-30v", username)
			fmt.Printf("%-30v", reimbursements)
			fmt.Printf("%-30v", currentstatus)
			fmt.Println("")

		}

		var status string
		var id string

		fmt.Println("")
		fmt.Println("To accept a reimbursement type accept")
		fmt.Println("To decline a reimbursement type decline")
		fmt.Scan(&status)

		switch status {
		case "accept":
			fmt.Println("Enter an id: ")
			fmt.Scan(&id)
			sql := `update employees set currentstatus=$1 where id=$2`

			_, err := db.Exec(sql, "accepted", id)
			if err != nil {
				panic(err)
			}

		case "decline":
			fmt.Println("Enter an id: ")
			fmt.Scan(&id)

			sql := `update employees set currentstatus=$1 where id=$2`

			_, err := db.Exec(sql, "denied", id)
			if err != nil {
				panic(err)
			}

		default:
			fmt.Println("invalid string")
			Manager()
		}

	default:
		fmt.Println("Invalid Option")
		Manager()

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
