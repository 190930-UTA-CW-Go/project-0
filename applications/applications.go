package applications

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/artytheparty/project-0/bank"
)

//AccountHolder will hold the necessary information to create an accoun after it is approved
type AccountHolder struct {
	username, password, fname, lname, types string
	bal                                     float64
}

//CreateAccountHolder is a constructor for the AccountHolder struct
func CreateAccountHolder(username string, pass string, fname string, lname string, types string, bal float64) AccountHolder {
	return AccountHolder{username, pass, fname, lname, types, bal}
}

//AccountHolderInfo2 will just return the Account Information as a string.
func (a *AccountHolder) AccountHolderInfo2() string {
	//var balance string = fmt.Sprintf("%f", a.bal)
	return ("USERNAME: " + a.username +
		" PASSWORD: " + a.password +
		" FNAME: " + a.fname +
		" LNAME: " + a.lname +
		" TYPE: " + a.types +
		" BALANCE: " + strconv.FormatFloat(a.bal, 'f', 2, 64) +
		"\n")
}

//WriteToFile writes the Array of AccountHolder{s} to accounts.txt so that it can be loadedin later for employee review
func WriteToFile(c []AccountHolder) {
	f, err := os.Create("accounts.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range c {
		f.WriteString(c[i].AccountHolderInfo2())
	}
	err = f.Close()
}

//ReadFile will read the accounts.txt file so that we can recreate the slice of AccountHolder{s} for the employee
func ReadFile() []AccountHolder {
	var holder []AccountHolder
	f, err := os.Open("accounts.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewReader(f)
	var txtline string
	for {
		txtline, err = scanner.ReadString('\n')
		if err != nil {
			break
		}
		uname := txtline[10 : strings.Index(txtline, "PASSWORD: ")-1]
		pass := txtline[strings.Index(txtline, "PASSWORD: ")+10 : strings.Index(txtline, "FNAME: ")-1]
		fname := txtline[strings.Index(txtline, "FNAME: ")+7 : strings.Index(txtline, "LNAME: ")-1]
		lname := txtline[strings.Index(txtline, "LNAME: ")+7 : strings.Index(txtline, "TYPE: ")-1]
		types := txtline[strings.Index(txtline, "TYPE: ")+6 : strings.Index(txtline, "BALANCE: ")-1]
		aBal, _ := strconv.ParseFloat(txtline[strings.Index(txtline, "BALANCE: ")+9:len(txtline)-1], 64)
		holder = append(holder, CreateAccountHolder(uname, pass, fname, lname, types, aBal))
	}
	if err != io.EOF {
		fmt.Printf("failed: %v\n", err)
	}
	return holder
}

//PrintHolder prints the AccountHolder Slice | use case for showing accounts to employee for approval/denial
func PrintHolder(a []AccountHolder) {
	for i := range a {
		fmt.Println("Temp. Acc Number", i)
		fmt.Println(a[i].AccountHolderInfo2())
	}
}

//AskNewAccount collects necessary information from a potential new user when applying for a new account
func AskNewAccount(db *sql.DB) {
	var unapprovedAccounts []AccountHolder
	unapprovedAccounts = ReadFile()
	var username string
	var pass string
	var fname string
	var lname string
	var bal float64
	var types string
	fmt.Println("Enter your username: ")
	fmt.Scan(&username)
	var usernameHolder []string
	rows, _ := db.Query("SELECT username FROM users")
	for rows.Next() {
		var uN string
		rows.Scan(&uN)
		usernameHolder = append(usernameHolder, uN)
	}
	var continu bool = true
	for i := range usernameHolder {
		if usernameHolder[i] == username {
			continu = false
			break
		}
	}
	if continu {
		fmt.Println("Enter your password: ")
		fmt.Scan(&pass)
		fmt.Println("Enter your First Name: ")
		fmt.Scan(&fname)
		fmt.Println("Enter your Last Name: ")
		fmt.Scan(&lname)
		fmt.Println("Enter your Account type \"c\" for checking \"s\" for savings: ")
		fmt.Scan(&types)
		fmt.Println("Enter your Initial deposit: ")
		fmt.Scan(&bal)
		unapprovedAccounts = append(unapprovedAccounts, CreateAccountHolder(username, pass, fname, lname, types, bal))
		WriteToFile(unapprovedAccounts)
	} else {
		println("Please try again with a different username, entered Username has been taken!")
		AskNewAccount(db)
	}

}

//ApproveAndAddToDB willl add a selected accountholder to the user and accounts table in the database
func ApproveAndAddToDB(db *sql.DB, a []AccountHolder) {
	fmt.Println("Which account would you like to add/remove?")
	PrintHolder(a)
	fmt.Println("Enter temporary account number: ")
	var choice int
	fmt.Scan(&choice)
	fmt.Println("[r]emove? or [a]pprove]: ?")
	var choice2 string
	fmt.Scan(&choice2)
	holder := a[choice]
	if choice2 == "a" {
		bank.CreateNewUserEntry(holder.username, holder.password, holder.fname, holder.lname, db)
		usrholder := bank.GetUsrInfo(holder.username, db)
		bank.CreateNewAccountEntry(usrholder.GetUsrID(), holder.types, holder.bal, db)
		b := remove(a, choice)
		WriteToFile(b)
		fmt.Println("Successfully added the user and their account to Database")
	}
	if choice2 == "r" {
		b := remove(a, choice)
		WriteToFile(b)
		fmt.Println("Successfully removed Pending Account from file.")
	}
}

//remove will remove the accountHolder from a slice which has been either approved by the employee or removed
func remove(s []AccountHolder, i int) []AccountHolder {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
