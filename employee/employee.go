package employee

import (
	"bufio"
	"database/sql"
	_ "flag" //no
	"fmt"
	_ "log" //no
	"math"
	"os"
	"os/exec"
	_ "os/exec" //no
	"strconv"
	_ "strconv" //no
	"text/tabwriter"
	_ "text/tabwriter" //no
	_ "time"           //yes

	"github.com/gookit/color"
	_ "github.com/gookit/color" //no
	_ "github.com/lib/pq"       // no
)

var userName string //global variable that is overwritten everytime a new user logs in
var hwatt string

// NewAcc Opens prompt to create a new account.
func NewAcc() {
	var password, fname, lname string
	Clear()
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
	Clear()
	fmt.Printf("Employee account for ")
	color.Cyan.Printf(fname)
	fmt.Printf(" ")
	color.Cyan.Printf(lname)
	fmt.Printf(" has been created!")
	fmt.Println()
	fmt.Println("Congradjulashens")
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
	Clear()
	//color.Yellow.Println("---Main Menu---")
	color.Yellow.Println(` __  __    _    ___  _  _   __  __  ___  _  _  _   _ `)
	color.Yellow.Println(`|  \/  |  /_\  |_ _|| \| | |  \/  || __|| \| || | | |`)
	color.Yellow.Println(`| |\/| | / _ \  | | | .  | | |\/| || _| | .  || |_| |`)
	color.Yellow.Println(`|_|  |_|/_/ \_\|___||_|\_| |_|  |_||___||_|\_| \___/ `)
	color.Yellow.Println(`													`)

	//	fmt.Println("---Main Menu---")
	fmt.Println()
	fmt.Println("What would you like to do? Press number to choose: ")
	fmt.Println()
	fmt.Println("	1: Employee login")
	fmt.Println("	2: Manager Login")
	fmt.Println("	3: Create employee account")
	fmt.Println("	4: Create manager account")
	fmt.Println("	5: Exit application")
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
		Prompt()
		Welcome()
	case 4:
		createManager()
		Prompt()
		Welcome()
	case 5:
		Clear()
		color.Red.Println("Exiting . . .")
		os.Exit(0)
	default:
		fmt.Println()
		fmt.Println("( ⚆ _ ⚆ ) Pick an appropriate option")
		fmt.Println()
		Prompt()
		Welcome()
	}
}

func employeeLogin() {
	var tryPassword, actualPassword string
	var fname, lname string
	var serial int
	Clear()
	color.Yellow.Println("---Employee Login---")
	fmt.Println()
	fmt.Printf("Enter credentials:  ")
	fmt.Scanln(&userName)
	fmt.Printf("Enter password:  ")
	fmt.Scanln(&tryPassword)
	db := OPEN()
	defer db.Close()
	row := db.QueryRow("SELECT * FROM employeeLogin WHERE userName = $1", userName)
	row.Scan(&serial, &userName, &actualPassword, &fname, &lname)
	if tryPassword == actualPassword {
		fmt.Println()
		//	fmt.Println("Logged in as " + userName)
		fmt.Printf("Logged in as ")
		color.Cyan.Println(userName)
		Prompt()
		welcomeEmployee()
	} else {
		Clear()
		color.Red.Println("Failure, incorrect credentials")
		color.Red.Println("Redirecting you. . .")
		fmt.Println()
		fmt.Println("		(ノಥ﹏ಥ)ノ彡┻━┻")
		fmt.Println()
		Prompt()
		employeeLogin()
	}
}

func welcomeEmployee() {
	Clear()
	color.Yellow.Println("---Employee Menu---")
	fmt.Println()
	fmt.Println("What would you like to do? Press number to choose: ")
	fmt.Println()
	fmt.Println("	1: Submit reimbursement ticket")
	fmt.Println("	2: View reimbursement status")
	fmt.Println("	3: Log out")
	fmt.Println("	4: Exit")
	fmt.Println()
	fmt.Println("Selection: ")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		reimburseReq()
		Prompt()
		welcomeEmployee()
	case 2:
		viewMyreimburses()
		Prompt()
		welcomeEmployee()
	case 3:
		Welcome()
	case 4:
		Clear()
		color.Red.Println("Exiting. . .")
		os.Exit(0)
	default:
		fmt.Println()
		fmt.Println("( ⚆ _ ⚆ ) Pick an appropriate option")
		fmt.Println()
		Prompt()
		welcomeEmployee()
	}
}

//ManagerLogin d
func ManagerLogin() {
	var tryPassword, password, adminLogin string
	Clear()
	color.Blue.Println("===MANAGER LOGIN===")
	fmt.Println()
	fmt.Printf("ENTER CREDENTIALS:  ")
	fmt.Scanln(&adminLogin)
	fmt.Printf("ENTER PASSWORD:  ")
	fmt.Scanln(&tryPassword)
	db := OPEN()
	defer db.Close()
	row := db.QueryRow("SELECT * FROM employeeAccounts WHERE adminLogin = $1", adminLogin)
	row.Scan(&adminLogin, &password)
	if tryPassword == password {
		Clear()
		fmt.Printf("WELCOME, ")
		color.Magenta.Println(adminLogin)
		Prompt()
		welcomeAdmin()
	} else {
		Clear()
		color.Red.Println("Failure, incorrect credentials")
		color.Red.Println("Try harder")
		fmt.Println()
		fmt.Println("		(ノಥ﹏ಥ)ノ彡┻━┻")
		fmt.Println()
		Prompt()
		ManagerLogin()
	}
}

func welcomeAdmin() {
	Clear()
	color.Blue.Println("===MANAGER MENU===")
	fmt.Println()
	fmt.Println("MAKE A SELECTION:")
	fmt.Println()
	fmt.Println("	1: LIST ALL OF THE PRISONERS WITH JOBS")
	fmt.Println("	2: LIST PENDING REIMBURSEMENT TICKETS")
	fmt.Println("	3: LIST ALL REIMBURSEMENT TICKETS")
	fmt.Println("	4: APPROVE/DENY TICKET REQUESTS")
	fmt.Println("	5: LOG OUT")
	fmt.Println("	6: EXIT")
	fmt.Println()
	fmt.Println("SELECTION: ")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		Clear()
		color.Blue.Println("===PRISONERS WITH JOBS===")
		fmt.Println()
		db := OPEN()
		defer db.Close()
		rows, _ := db.Query("SELECT * FROM EMPLOYEELOGIN")
		w := new(tabwriter.Writer)
		var fullstring string
		w.Init(os.Stdout, 12, 0, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)
		fmt.Fprintln(w, "USERNAME\tPASSWORD\tFIRSTNAME\tLASTNAME")
		fmt.Fprintln(w, "========\t========\t=========\t========")
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
		Prompt()
		welcomeAdmin()
	case 2:
		Clear()
		color.Blue.Println("===HU WANTS FREE HANDOUTS===")
		fmt.Println()
		db := OPEN()
		defer db.Close()
		var what2 = " pending"
		rows, _ := db.Query("SELECT * FROM tickets WHERE what = $1", what2)
		w := new(tabwriter.Writer)
		var fullstring string
		/*rows, _ := db.Query("SELECT * FROM tickets WHERE userName = $1", userName)*/
		w.Init(os.Stdout, 12, 0, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)
		fmt.Fprintln(w, "TICKET#\tUSERNAME\tAMOUNT\tSTATUS\tREASON")
		fmt.Fprintln(w, "=======\t========\t======\t======\t======")
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
		Prompt()
		welcomeAdmin()
	case 3:
		Clear()
		color.Blue.Println("===ALL TICKETS===")
		fmt.Println()
		db := OPEN()
		defer db.Close()
		rows, _ := db.Query("SELECT * FROM TICKETS")
		w := new(tabwriter.Writer)
		var fullstring string

		//w.Init(os.Stdout, 0, 8, 2, '*', tabwriter.Debug|tabwriter.AlignRight)
		w.Init(os.Stdout, 12, 0, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)
		fmt.Fprintln(w, "TICKET#\tUSERNAME\tAMOUNT\tSTATUS\tREASON")
		fmt.Fprintln(w, "=======\t========\t======\t======\t======")
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
		Prompt()
		welcomeAdmin()
	case 4:
		Approvedeny()
		Prompt()
		welcomeAdmin()
	case 5:
		Welcome()
	case 6:
		Clear()
		color.Red.Println("Exiting. . .")
		os.Exit(0)
	default:
		fmt.Println()
		fmt.Println("( ⚆ _ ⚆ ) Pick an appropriate option")
		fmt.Println()
		Prompt()
		welcomeAdmin()
	}
}

func createManager() {
	var adminLogin string
	var askpassword string
	var password string
	var masterpassword = "master"
	Clear()
	color.Blue.Println("===CREATING AN ADMINISTRATOR ACCOUNT===")
	fmt.Println()
	color.Blue.Println("PLEASE ENTER THE MASTER PASSWORD TO PROCEED:")
	fmt.Println("")
	fmt.Scanln(&askpassword)
	if askpassword == masterpassword {
		fmt.Println("CORRECT")
	} else {
		Clear()
		color.Red.Println("Failure, incorrect credentials")
		color.Red.Println("nuuuuu")
		fmt.Println()
		fmt.Println("		(ノಥ﹏ಥ)ノ彡┻━┻")
		fmt.Println()
		Prompt()
	}
	Clear()
	fmt.Println()
	fmt.Printf("PICK A USER NAME:  ")
	fmt.Scanln(&adminLogin)
	fmt.Println()
	fmt.Printf("PICK A PASSWORD:  ")
	fmt.Scanln(&password)
	db := OPEN()
	defer db.Close()
	db.Exec("INSERT INTO employeeAccounts(adminLogin,password)"+
		"VALUES($1,$2)", adminLogin, password)
	fmt.Println()
	fmt.Printf("CREATED ")
	color.Magenta.Printf(adminLogin)
	fmt.Println(" AS MANAGER ACCOUNT")
}

func reimburseReq() {
	Clear()
	color.Yellow.Println("---Employee Reimbursement Menu---")
	fmt.Println()
	var reimburse float64
	var reason, fname, lname string
	var what = " pending"
	Clear()
	fmt.Println("Enter a reimbursement amount:  $")
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
}

func viewMyreimburses() {
	Clear()
	color.Yellow.Println("---Viewing your tickets---")
	fmt.Println()
	db := OPEN()
	defer db.Close()
	Clear()
	fmt.Printf("Displaying reimbursement tickets for ")
	color.Cyan.Println(userName)
	//fmt.Println(userName)
	rows, _ := db.Query("SELECT * FROM tickets WHERE userName = $1", userName)
	w := new(tabwriter.Writer)
	var fullstring string
	w.Init(os.Stdout, 12, 0, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(w, "TICKET#\tAMOUNT\tSTATUS\tREASON")
	fmt.Fprintln(w, "======#\t======\t======\t======")
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
	fmt.Println("Successfully displayed all tickets for this user.")
}

//Approvedeny d
func Approvedeny() {
	Clear()
	color.Blue.Println("===APPROVE OR DENY REIMBURSEMENT TICKETS===")
	fmt.Println()
	fmt.Println("SELECT AN OPTION: ")
	fmt.Println()
	fmt.Println("	1: APPROVE/DENY TICKET")
	fmt.Println("	2: VIEW PENDING TICKETS")
	//view pending tickets
	fmt.Println("	3: RETURN TO MANAGER MENU")
	fmt.Println("	4: EXIT")
	fmt.Println()
	fmt.Println("SELECTION")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println()
		fmt.Println("INDICATE TICKET# :")
		var ticketid int
		fmt.Scanln(&ticketid)
		db := OPEN()
		defer db.Close()

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
		fullstring = (str1 + "\t" + u2 + "\t$" + str5 + "\t" + u7 + "\t" + u6 + "\t")
		fmt.Fprintln(w, fullstring)
		fmt.Fprintln(w)
		w.Flush()
		fmt.Println("WHAT WOULD YOU LIKE TO DO ABOUT THIS TICKET?")
		fmt.Println()
		fmt.Println("	1: APPROVE")
		fmt.Println("	2: DENY")
		fmt.Println("	3. IGNORE")
		var choice2 int
		fmt.Scanln(&choice2)
		switch choice2 {
		case 1:
			var yess = "APPROVED"
			db := OPEN()
			defer db.Close()
			db.Exec("UPDATE tickets SET what = $1 WHERE ticketNum = $2", yess, u1)
			fmt.Println("SUCCESSFULLY APPROVED TICKET.")
			Prompt()
			Approvedeny()
		case 2:
			var noo = "  DENIED"
			db := OPEN()
			defer db.Close()
			db.Exec("UPDATE tickets SET what = $1 WHERE ticketNum = $2", noo, u1)
			fmt.Println("SUCCESSFULLY DENIED TICKET.")
			Prompt()
			Approvedeny()
		case 3:
			fmt.Println("NO ACTION TAKEN.")
			Prompt()
			Approvedeny()
		default:
			fmt.Println()
			fmt.Println("( ⚆ _ ⚆ ) Pick an appropriate option")
			fmt.Println()
			Approvedeny()
		}
	case 2:
		Clear()
		color.Blue.Println("===PENDING TICKETS===")
		fmt.Println()
		db := OPEN()
		defer db.Close()
		var what2 = " pending"
		rows, _ := db.Query("SELECT * FROM tickets WHERE what = $1", what2)
		w := new(tabwriter.Writer)
		var fullstring string
		/*rows, _ := db.Query("SELECT * FROM tickets WHERE userName = $1", userName)*/
		w.Init(os.Stdout, 12, 0, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)
		fmt.Fprintln(w, "TICKET#\tUSERNAME\tAMOUNT\tSTATUS\tREASON")
		fmt.Fprintln(w, "=======\t========\t======\t======\t======")
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
		Prompt()
		Approvedeny()
	case 3:
		welcomeAdmin()
	case 4:
		Clear()
		color.Red.Println("Exiting. . .")
		os.Exit(0)
	default:
		fmt.Println()
		fmt.Println("( ⚆ _ ⚆ ) Pick an appropriate option")
		fmt.Println()
		Prompt()
		welcomeAdmin()
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

//Clear d
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

//Prompt pauses
func Prompt() {
	fmt.Println()
	color.Green.Println("+===========================+")
	color.Green.Println("| PRESS ANY KEY TO CONTINUE |")
	color.Green.Println("+===========================+")
	fmt.Scanln(&hwatt)
}
