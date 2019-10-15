package employee

import (
	"bufio"
	_ "bufio" //no
	"database/sql"
	_ "flag" //no
	"fmt"
	_ "log" //no
	"os"
	_ "os" //no
	_ "os/exec"
	_ "strconv"

	_ "github.com/lib/pq" // no
)

var userName string

// NewAcc Opens prompt to create a new account.
func NewAcc() {

	var password string
	var fname string
	var lname string
	fmt.Printf("Creating a new account:")
	fmt.Println()
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
	db.Exec("INSERT INTO employeeLogin(userName,password,fname, lname)"+
		"VALUES($1,$2,$3, $4)", userName, password, fname, lname)
}

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

//SearchUser d
func SearchUser(searchvalue string) {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	row := db.QueryRow("SELECT * FROM employeeLogin WHERE userName = $1", searchvalue)
	var u1 int
	var u2 string
	var u3 string
	var u4 string
	var u5 string
	row.Scan(&u1, &u2, &u3, &u4, &u5)
	fmt.Println(u1, u2, u3, u4, u5)
}

//Welcome starts dis
func Welcome() {
	fmt.Println("Welcome to Employee reimbursement app")
	fmt.Println("What would you like to do? Press number to choose: ")
	fmt.Println("1: Log in as Employee")
	fmt.Println("2: Log in as Manager")
	fmt.Println("3: Create an account")
	fmt.Println("4: Exit application")
	fmt.Println()

	var choice int

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		employeeLogin()
	case 2:
		managerLogin()
	case 3:
		NewAcc()
		fmt.Println()
		fmt.Println("Congradjulashens you have made an account")
		fmt.Println("Redirecting you to main menu. . . .")
		fmt.Println()
		Welcome()
	case 4:
		os.Exit(0)
	}
}

func employeeLogin() {
	var tryPassword, actualPassword string
	var fname, lname string
	var serial int
	fmt.Printf("Enter credentials:  ")
	fmt.Scanln(&userName)
	fmt.Printf("Enter password:  ")
	fmt.Scanln(&tryPassword)
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	row := db.QueryRow("SELECT * FROM employeeLogin WHERE userName = $1", userName)
	row.Scan(&serial, &userName, &actualPassword, &fname, &lname)
	fmt.Println("actual password" + actualPassword)
	fmt.Println("try password" + tryPassword)
	if tryPassword == actualPassword {
		fmt.Println("Logged in as " + userName)
		welcomeEmployee()
	} else {
		fmt.Println("Failed")
		fmt.Println("Program will now abort")
		os.Exit(0)
	}
}

func welcomeEmployee() {
	fmt.Println()
	fmt.Println("EMPLOYEE PORTAL")
	fmt.Println("What would you like to do? Press number to choose: ")
	fmt.Println("1: Submit reimbursement ticket")
	fmt.Println("2: View reimbursement status")
	fmt.Println("3: Log out")
	fmt.Println("4: Exit")
	fmt.Println()
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		reimburseReq()
	case 2:
		viewMyreimburses()
	case 3:
		Welcome()
	case 4:
		os.Exit(0)
	}
}

func managerLogin() {
	//
}

func reimburseReq() {
	var reimburse float64
	var reason, fname, lname string
	fmt.Println("Enter reimbursement amount  ")
	fmt.Scanln(&reimburse)
	fmt.Println("Reason:  ")
	fmt.Println(" ")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	reason = input

	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	row := db.QueryRow("SELECT * FROM employeeLogin WHERE userName = $1", userName)
	var password string
	var serial int

	row.Scan(&serial, &userName, &password, &fname, &lname)
	/*db.Exec("INSERT INTO employeeLogin(userName,password,fname, lname)"+
	"VALUES($1,$2,$3, $4)", userName, password, fname, lname)*/
	db.Exec("INSERT INTO tickets(userName,fname, lname, reimburse, reason)"+
		"VALUES($1,$2,$3, $4, $5)", userName, fname, lname, reimburse, reason)
	fmt.Println()
	fmt.Println("Ticket successfully submitted")
	fmt.Println("Redirecting you. . . . .")
	fmt.Println()
	welcomeEmployee()
}

func viewMyreimburses() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("Displaying reimbursement tickets:")
	fmt.Println()
	fmt.Println(userName)
	//rows, _ := db.Query("SELECT ticketNum, userName, fName, lName, reimburse, reason FROM tickets WHERE userName = $1", userName)
	rows, _ := db.Query("SELECT * FROM tickets WHERE userName = $1", userName)
	for rows.Next() {
		var u1 int
		var u2 string
		var u3 string
		var u4 string
		var u5 float64
		var u6 string
		rows.Scan(&u1, &u2, &u3, &u4, &u5, &u6)
		fmt.Println(u1, u2, u3, u4, u5, u6)
	}
	/* func GetAll3(db *sql.DB) {
		rows, _ := db.Query("SELECT * FROM TICKETS")
		for rows.Next() {
			var u1 int
			var u2 string
			var u3 string
			var u4 string
			var u5 float32
			var u6 string
			rows.Scan(&u1, &u2, &u3, &u4, &u5, &u6)
			fmt.Println(u1, u2, u3, u4, u5, u6)
		}
	}
	*/
	fmt.Println("Redirecting you. . . . .")
	fmt.Println()
	welcomeEmployee()

}
