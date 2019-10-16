package print

import (
	"fmt"
	"strconv"

	"github.com/gittingdavid/project-0/database"
)

// Table = Prints table
// param1 = identify which table "customer" or "employee"
func Table(who string) {
	var count int
	var email, pass, first, last string
	fmt.Printf("%-30v", "Login ID:")
	fmt.Printf("%-20v", "Password:")
	fmt.Printf("%-20v", "First Name:")
	fmt.Printf("%-20v", "Last Name:")
	fmt.Println()
	fmt.Print("================================================")
	fmt.Println("==============================================")

	sqlStatement := ``
	if who == "customer" {
		sqlStatement = `select * from customer`
	} else {
		sqlStatement = `select * from employee`
	}
	rows, _ := (database.DBCon).Query(sqlStatement)

	for rows.Next() {
		// count variable used as empty table error checker
		count++
		rows.Scan(&email, &pass, &first, &last)
		fmt.Printf("%-30v", email)
		fmt.Printf("%-20v", pass)
		fmt.Printf("%-20v", first)
		fmt.Printf("%-20v", last)
		fmt.Println()
	}

	if count == 0 {
		fmt.Println("No Data in Table")
	}

	fmt.Print("================================================")
	fmt.Println("==============================================")
	fmt.Println()
}

// Accounts = Print accounts associated with login id
// param1 = customer login id
func Accounts(login string) {
	var count, number int
	var email, name string
	var balance float32
	fmt.Printf("%-30v", "Login ID:")
	fmt.Printf("%-20v", "Account Type:")
	fmt.Printf("%-20v", "Account Balance:")
	fmt.Printf("%-20v", "Account Number:")
	fmt.Println()
	fmt.Print("================================================")
	fmt.Println("==============================================")

	sqlStatement := ""
	if login == "" {
		sqlStatement = `select * from account order by email`
		rows, _ := (database.DBCon).Query(sqlStatement)
		for rows.Next() {
			// count variable used as empty table error checker
			count++
			rows.Scan(&email, &name, &balance, &number)
			fmt.Printf("%-30v", email)
			fmt.Printf("%-20v$", name)
			s := fmt.Sprintf("%.2f", balance)
			fmt.Printf("%-20v", s)
			fmt.Printf("%-20v", number)
			fmt.Println()
		}
	} else {
		sqlStatement = `select * from account where email = $1`
		rows, _ := (database.DBCon).Query(sqlStatement, login)
		for rows.Next() {
			// count variable used as empty table error checker
			count++
			rows.Scan(&email, &name, &balance, &number)
			fmt.Printf("%-30v", email)
			fmt.Printf("%-20v$", name)
			s := fmt.Sprintf("%.2f", balance)
			fmt.Printf("%-20v", s)
			fmt.Printf("%-20v", number)
			fmt.Println()
		}
	}

	if count == 0 {
		fmt.Println("No Data in Table")
	}

	fmt.Print("================================================")
	fmt.Println("==============================================")
	fmt.Println()
}

// Joints = Prints joint table
func Joints() (count int, slice []string) {
	count = 0
	var index, email1, email2 string
	var num1, num2 int

	fmt.Print("   ")
	fmt.Printf("%-25v", "#1 Login:")
	fmt.Printf("%-25v", "#2 Login:")
	fmt.Printf("%-20v", "#1 Account ID:")
	fmt.Printf("%-20v", "#2 Account ID:")
	fmt.Println()
	fmt.Print("================================================")
	fmt.Println("==============================================")

	sqlStatement := "select * from joint"
	rows, _ := (database.DBCon).Query(sqlStatement)
	for rows.Next() {
		count++

		rows.Scan(&index, &email1, &email2, &num1, &num2)
		slice = append(slice, index)
		fmt.Print(strconv.Itoa(count) + ") ")
		fmt.Printf("%-25v", email1)
		fmt.Printf("%-25v", email2)
		fmt.Printf("%-20v", num1)
		fmt.Printf("%-20v", num2)
		fmt.Println()
	}

	if count == 0 {
		fmt.Println("No Data in Table")
	}

	fmt.Print("================================================")
	fmt.Println("==============================================")
	fmt.Println()
	return
}

// Invalid =
func Invalid() {
	fmt.Println("> Invalid Input")
	fmt.Println()
}
