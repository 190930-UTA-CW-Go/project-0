package employee

import (
	"database/sql"
	"fmt"
	_ "log" //no
	"os"
	_ "os" //no
	_ "os/exec"
	_ "strconv"

	_ "github.com/lib/pq" // no
)

// NewAccGuest Opens prompt to create a new account.
func NewAccGuest() {
	var userName string
	var password string
	var fname string
	var lname string
	fmt.Printf("Creating a new account:")
	fmt.Printf("Enter a username:  ")
	fmt.Scanln(&userName)
	fmt.Printf("Enter a password:  ")
	fmt.Scanln(&password)
	fmt.Printf("Enter your first:  ")
	fmt.Scanln(&fname)
	fmt.Printf("Enter your last:  ")
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

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

//GetAll2 exports
func GetAll2(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM CUSTOMERLOGIN")
	for rows.Next() {
		var u1 string
		var u2 string
		var u3 string
		var u4 string
		rows.Scan(&u1, &u2, &u3, &u4)
		fmt.Println(u1, u2, u3, u4)
	}
}

//SearchByName2 exports
func SearchByName2(db *sql.DB, searchvalue string) {
	row := db.QueryRow("SELECT * FROM customerLogin WHERE password = $1", searchvalue)
	var userName string
	var password string
	var fname string
	var lname string
	row.Scan(&userName, &password, &fname, &lname)
	fmt.Println(userName, password, fname, lname)
}

//Welcome starts dis
func Welcome() {
	fmt.Println("Welcome to Employee reimbursement app")
	fmt.Println("What would you like to do? Press number to choose: ")
	fmt.Println("1: Log in as Employee")
	fmt.Println("2: Log in as Manager")
	fmt.Println("3: Create an account")
	fmt.Println("4: Exit application")
	var choice int
	switch choice {
	case 1:
		employeeLogin()
	case 2:
		managerLogin()
	case 3:
		NewAccGuest()
	case 4:
		os.Exit(0)
	}
}

func employeeLogin() {
	var user, tryPassword, actualPassword string
	var userName, fname, lname string
	fmt.Printf("Enter credentials:  ")
	fmt.Scanln(&user)
	fmt.Printf("Enter password:  ")
	fmt.Scanln(&tryPassword)
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	row := db.QueryRow("SELECT * FROM customerLogin WHERE userName = $1", userName)
	row.Scan(&userName, &actualPassword, &fname, &lname)
	if tryPassword == actualPassword {
		fmt.Println("Logged in as " + userName)
		welcomeEmployee()
	} else {
		fmt.Println("Failed")
	}
}

func welcomeEmployee() {
	fmt.Println("Welcome")
	fmt.Println("What would you like to do? Press number to choose: ")
	fmt.Println("1: Submit reimbursement ticket")
	fmt.Println("2: View reimbursement status")
	fmt.Println("3: Exit")
	var choice int
	switch choice {
	case 1:
		reimburseReq()
	case 2:
		viewMyreimburses()
	case 3:
		os.Exit(0)
	}

}
