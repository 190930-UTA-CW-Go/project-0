package employee

import (
	"bufio"
	"database/sql"
	_ "flag" //no
	"fmt"
	_ "log" //no
	"math"
	"os"
	_ "os/exec" //no
	"strconv"
	_ "strconv" //no
	"text/tabwriter"
	_ "text/tabwriter" //no

	_ "github.com/lib/pq" // no
)

var userName string //global variable that is overwritten everytime a new user logs in

// NewAcc Opens prompt to create a new account.
func NewAcc() {
	var password, fname, lname string
	fmt.Println("================================================================================")
	fmt.Printf("Creating a new account:")
	fmt.Println()
	fmt.Printf("Enter a username:  ")
	fmt.Scanln(&userName)
	fmt.Printf("Enter a password:  ")
	fmt.Scanln(&password)
	fmt.Printf("Enter your first name:  ")
	fmt.Scanln(&fname)
	fmt.Printf("Enter your last name:  ")
	fmt.Scanln(&lname)
	db := OPEN()
	defer db.Close()
	db.Exec("INSERT INTO employeeLogin(userName,password,fname, lname)"+
		"VALUES($1,$2,$3, $4)", userName, password, fname, lname)
	fmt.Println("Employee account for " + fname + " " + lname + " has been created.")
	fmt.Println("Congradjulashens you have made an account")
}

//SearchUser Test function to search for a specific user.
func SearchUser(searchvalue string) {
	db := OPEN()
	defer db.Close()
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
	fmt.Println("================================================================================")
	fmt.Println("---Main Terminal---")
	fmt.Println()
	fmt.Println("What would you like to do? Press number to choose: ")
	fmt.Println("1: Employee login")
	fmt.Println("2: Manager Login")
	fmt.Println("3: Create employee account")
	fmt.Println("4: Create manager account")
	fmt.Println("5: Exit application")
	fmt.Println()
	var choice int
	fmt.Printf("Enter choice: ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		employeeLogin()
	case 2:
		ManagerLogin()
	case 3:
		NewAcc()
		fmt.Println()
		fmt.Println("Redirecting you to main menu. . . .")
		Welcome()
	case 4:
		createManager()
		fmt.Println("REDIRECTING. . . .")
		Welcome()
	case 5:
		fmt.Println("EXITING. . . .")
		os.Exit(0)
	}
}

func employeeLogin() {
	var tryPassword, actualPassword string
	var fname, lname string
	var serial int
	fmt.Println("================================================================================")
	fmt.Printf("Enter credentials:  ")
	fmt.Scanln(&userName)
	fmt.Printf("Enter password:  ")
	fmt.Scanln(&tryPassword)
	db := OPEN()
	defer db.Close()
	row := db.QueryRow("SELECT * FROM employeeLogin WHERE userName = $1", userName)
	row.Scan(&serial, &userName, &actualPassword, &fname, &lname)
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
	fmt.Println("================================================================================")
	fmt.Println("Employee Portal")
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

//ManagerLogin d
func ManagerLogin() {
	var tryPassword, password, adminLogin string
	fmt.Println("================================================================================")
	fmt.Printf("Enter credentials:  ")
	fmt.Scanln(&adminLogin)
	fmt.Printf("Enter password:  ")
	fmt.Scanln(&tryPassword)
	db := OPEN()
	defer db.Close()
	row := db.QueryRow("SELECT * FROM employeeAccounts WHERE adminLogin = $1", adminLogin)
	row.Scan(&adminLogin, &password)
	if tryPassword == password {
		fmt.Println("Logged in as " + adminLogin)
		welcomeAdmin()
	} else {
		fmt.Println("Failed")
		fmt.Println("Program will now abort")
		os.Exit(0)
	}
}

func welcomeAdmin() {
	fmt.Println("================================================================================")
	fmt.Println("ADMIN PORTAL")
	fmt.Println("WHAT YOU WANT")
	fmt.Println("1: LIST ALL OF THE PRISONERS WITH JOBS")
	fmt.Println("2: LIST PENDING REIMBURSEMENTS")
	fmt.Println("3: LIST ALL REIMBURSEMENTS")
	fmt.Println("4: APPROVE/DENY $$$")
	fmt.Println("5: LOG OUT")
	fmt.Println("6: EXIT")
	fmt.Println()
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("================================================================================")
		fmt.Println("HERE ARE THE PRISONERS WITH JOBS:")
		db := OPEN()
		defer db.Close()
		rows, _ := db.Query("SELECT * FROM EMPLOYEELOGIN")
		w := new(tabwriter.Writer)
		var fullstring string
		w.Init(os.Stdout, 12, 0, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)
		fmt.Fprintln(w, "USERNAME\tPASSWORD\tFIRSTNAME\tLASTNAME")
		for rows.Next() {
			var u1 int
			var u2 string
			var u3 string
			var u4 string
			var u5 string
			rows.Scan(&u1, &u2, &u3, &u4, &u5)
			fullstring = (u2 + "\t" + u3 + "\t" + u4 + "\t" + u5 + "\t")
			fmt.Fprintln(w, fullstring)
		}
		fmt.Fprintln(w)
		w.Flush()
		fmt.Println("SUCCESS. RETURNING. . . . .")

		welcomeAdmin()
	case 2:
		fmt.Println("================================================================================")
		fmt.Println("HU WANTS FREE HANDOUTS:")
		db := OPEN()
		defer db.Close()
		var what2 = " pending"
		rows, _ := db.Query("SELECT * FROM tickets WHERE what = $1", what2)
		w := new(tabwriter.Writer)
		var fullstring string
		/*rows, _ := db.Query("SELECT * FROM tickets WHERE userName = $1", userName)*/
		w.Init(os.Stdout, 12, 0, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)
		fmt.Fprintln(w, "TICKET#\tUSERNAME\tAMOUNT\tSTATUS\tREASON")
		for rows.Next() {
			var u1 int     //ticketnum
			var u2 string  //userName
			var u3 string  //fname first name
			var u4 string  //lname last name
			var u5 float64 //req amount
			var u6 string  //reason
			var u7 string  //status of req
			rows.Scan(&u1, &u2, &u3, &u4, &u5, &u6, &u7)
			str1 := strconv.Itoa(u1)
			u5 = math.Floor(u5*100) / 100
			str5 := fmt.Sprintf("%.2f", u5)
			fullstring = (str1 + "\t" + u2 + "\t$" + str5 + "\t" + u7 + "\t" + u6 + "\t")
			fmt.Fprintln(w, fullstring)
		}
		fmt.Fprintln(w)
		w.Flush()
		welcomeAdmin()
	case 3:
		fmt.Println("HU WANTS FREE HANDOUTS:")
		db := OPEN()
		defer db.Close()
		rows, _ := db.Query("SELECT * FROM TICKETS")
		w := new(tabwriter.Writer)
		var fullstring string

		//w.Init(os.Stdout, 0, 8, 2, '*', tabwriter.Debug|tabwriter.AlignRight)
		w.Init(os.Stdout, 12, 0, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)
		fmt.Fprintln(w, "TICKET#\tUSERNAME\tAMOUNT\tSTATUS\tREASON")
		for rows.Next() {
			var u1 int     //ticketnum
			var u2 string  //userName
			var u3 string  //fname first name
			var u4 string  //lname last name
			var u5 float64 //req amount
			var u6 string  //reason
			var u7 string  //status of req
			rows.Scan(&u1, &u2, &u3, &u4, &u5, &u6, &u7)
			str1 := strconv.Itoa(u1)
			fmt.Println("str1 is " + str1)
			u5 = math.Floor(u5*100) / 100
			str5 := fmt.Sprintf("%.2f", u5)
			fullstring = (str1 + "\t" + u2 + "\t$" + str5 + "\t" + u7 + "\t" + u6 + "\t")
			fmt.Fprintln(w, fullstring)
		}
		fmt.Fprintln(w)
		w.Flush()
		welcomeAdmin()
	case 4:
		Approvedeny()
	case 5:
		Welcome()
	case 6:
		fmt.Println("EXITING. . . . .")
		os.Exit(0)
	}
}

func createManager() {
	var adminLogin string
	var askpassword string
	var password string
	var masterpassword = "master"
	fmt.Println("================================================================================")
	fmt.Println("CREATING AN ADMINISTRATOR ACCOUNT:")
	fmt.Println("PLEASE ENTER THE MASTER PASSWORD TO PROCEED:")
	fmt.Println("")
	fmt.Scanln(&askpassword)
	if askpassword == masterpassword {
		fmt.Println("CORRECT")
	} else {
		fmt.Println("FAILURE")
		fmt.Println("FAILURE")
		fmt.Println("FAILURE")
		fmt.Println("FAILURE")
		fmt.Println("FAILURE")
		fmt.Println("FAILURE")
		fmt.Println("FAILURE")
		fmt.Println("FAILURE")
		fmt.Println("PROGRAM WILL NOW ABORT")
		os.Exit(0)
	}
	fmt.Println("================================================================================")
	fmt.Printf("Enter a username:  ")
	fmt.Scanln(&adminLogin)
	fmt.Printf("Enter a password:  ")
	fmt.Scanln(&password)
	db := OPEN()
	defer db.Close()
	db.Exec("INSERT INTO employeeAccounts(adminLogin,password)"+
		"VALUES($1,$2)", adminLogin, password)
	fmt.Println("SUCCESS. RETURNING YOU. . .")
	Welcome()
}

func reimburseReq() {
	var reimburse float64
	var reason, fname, lname string
	var what = " pending"
	fmt.Println("================================================================================")
	fmt.Println("Enter reimbursement amount  ")
	fmt.Scanln(&reimburse)
	reimburse = math.Floor(reimburse*100) / 100
	fmt.Println("Reason:  ")
	fmt.Println(" ")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	reason = input
	db := OPEN()
	defer db.Close()
	row := db.QueryRow("SELECT * FROM employeeLogin WHERE userName = $1", userName)
	var password string
	var serial int
	row.Scan(&serial, &userName, &password, &fname, &lname)
	db.Exec("INSERT INTO tickets(userName,fname, lname, reimburse, reason, what)"+
		"VALUES($1,$2,$3, $4, $5, $6)", userName, fname, lname, reimburse, reason, what)
	fmt.Println("Ticket successfully submitted")
	fmt.Println("SUCCESS Redirecting you. . . . .")
	welcomeEmployee()
}

func viewMyreimburses() {
	db := OPEN()
	defer db.Close()
	fmt.Println("================================================================================")
	fmt.Println("Displaying reimbursement tickets for " + userName)
	//fmt.Println(userName)
	rows, _ := db.Query("SELECT * FROM tickets WHERE userName = $1", userName)
	w := new(tabwriter.Writer)
	var fullstring string
	w.Init(os.Stdout, 12, 0, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(w, "TICKET#\tAMOUNT\tSTATUS\tREASON")
	for rows.Next() {
		var u1 int     //ticketnum
		var u2 string  //userName
		var u3 string  //fname first name
		var u4 string  //lname last name
		var u5 float64 //req amount
		var u6 string  //reason
		var u7 string  //status of req
		rows.Scan(&u1, &u2, &u3, &u4, &u5, &u6, &u7)
		str1 := strconv.Itoa(u1)
		u5 = math.Floor(u5*100) / 100
		str5 := fmt.Sprintf("%.2f", u5)
		fullstring = (str1 + "\t$" + str5 + "\t" + u7 + "\t" + u6 + "\t")
		fmt.Fprintln(w, fullstring)
		//fmt.Println(u1, u2, u5, u6, u7)
	}
	fmt.Fprintln(w)
	w.Flush()
	fmt.Println("SUCCESS. Redirecting you. . . . .")
	welcomeEmployee()
}

//Approvedeny d
func Approvedeny() {
	fmt.Println("================================================================================")
	fmt.Println("APPROVE // DENY REQUESTS. SELECT:")
	fmt.Println("1: INPUT SERIAL PRIMARY KEY OF THE TICKET")
	fmt.Println("2: VIEW TICKETS")
	//view pending tickets
	fmt.Println("3: RETURN")
	fmt.Println("4: EXIT")
	fmt.Println()
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("================================================================================")
		fmt.Println("SERIAL PRIMARY KEY:")
		var ticketid int
		fmt.Scanln(&ticketid)
		db := OPEN()
		defer db.Close()
		fmt.Println("id is ")
		fmt.Println(ticketid)
		row := db.QueryRow("SELECT * FROM tickets WHERE ticketNum = $1", ticketid)
		w := new(tabwriter.Writer)
		var fullstring string

		w.Init(os.Stdout, 12, 0, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)
		fmt.Fprintln(w, "TICKET#\tUSERNAME\tAMOUNT\tSTATUS\tREASON")
		var u1 int     //ticketnum
		var u2 string  //username
		var u3 string  //first name
		var u4 string  //last name
		var u5 float64 //amount
		var u6 string  //reason
		var u7 string  //status
		row.Scan(&u1, &u2, &u3, &u4, &u5, &u6, &u7)
		str1 := strconv.Itoa(u1)
		u5 = math.Floor(u5*100) / 100
		str5 := fmt.Sprintf("%.2f", u5)
		fullstring = (str1 + u2 + "\t$" + str5 + "\t" + u7 + "\t" + u6 + "\t")
		fmt.Fprintln(w, fullstring)
		fmt.Fprintln(w)
		w.Flush()
		fmt.Println("WHAT WOULD YOU LIKE TO DO ABOUT THIS REIMBURSEMENT?")
		fmt.Println("1: APPROVE")
		fmt.Println("2: DENY")
		fmt.Println("3. IGNORE")
		var choice2 int
		fmt.Scanln(&choice2)
		switch choice2 {
		case 1:
			var yess = "APPROVED"
			db := OPEN()
			defer db.Close()
			db.Exec("UPDATE tickets SET what = $1 WHERE ticketNum = $2", yess, u1)
			fmt.Println("SUCCESSFULLY APPROVED TICKET. RETURNING YOU. . .")
			Approvedeny()
		case 2:
			var noo = "  DENIED"
			db := OPEN()
			defer db.Close()
			db.Exec("UPDATE tickets SET what = $1 WHERE ticketNum = $2", noo, u1)
			fmt.Println("SUCCESSFULLY DENIED TICKET. RETURNING YOU. . .")
			Approvedeny()
		case 3:
			fmt.Println("NO ACTION TAKEN. RETURNING YOU. . .")
			Approvedeny()
		}
	case 2:
		fmt.Println("================================================================================")
		fmt.Println("HERE ARE THE PRISONERS WITH JOBS:")
		db := OPEN()
		defer db.Close()
		rows, _ := db.Query("SELECT * FROM TICKETS")
		w := new(tabwriter.Writer)
		var fullstring string

		//w.Init(os.Stdout, 0, 8, 2, '*', tabwriter.Debug|tabwriter.AlignRight)
		w.Init(os.Stdout, 12, 0, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)
		fmt.Fprintln(w, "TICKET#\tUSERNAME\tAMOUNT\tSTATUS\tREASON")
		for rows.Next() {
			var u1 int     //ticketnum
			var u2 string  //userName
			var u3 string  //fname first name
			var u4 string  //lname last name
			var u5 float64 //req amount
			var u6 string  //reason
			var u7 string  //status of req
			rows.Scan(&u1, &u2, &u3, &u4, &u5, &u6, &u7)
			str1 := strconv.Itoa(u1)
			fmt.Println("str1 is " + str1)
			u5 = math.Floor(u5*100) / 100
			str5 := fmt.Sprintf("%.2f", u5)
			fullstring = (str1 + "\t" + u2 + "\t$" + str5 + "\t" + u7 + "\t" + u6 + "\t")
			fmt.Fprintln(w, fullstring)
		}
		fmt.Fprintln(w)
		w.Flush()
		Approvedeny()
	case 3:
		welcomeAdmin()
	case 4:
		fmt.Println("EXITING. . . . .")
		os.Exit(0)
	}
}

//OPEN opens database; needed to clean up code
func OPEN() *sql.DB {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", datasource)
	//defer db.Close()
	if err != nil {
		panic(err)
	}
	return db
}
