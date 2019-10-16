package method

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gittingdavid/project-0/database"
	"github.com/gittingdavid/project-0/print"
)

// Length of account id number
const idLength = 3

// SelectQuery = shortcut function for QueryRow
func SelectQuery(target string, table string, condition string, value string) (hold string) {
	sqlStatement := `select $1 from $2 where $3 = $4`
	result := (database.DBCon).QueryRow(sqlStatement, target, table, condition, value)
	result.Scan(&hold)
	fmt.Println(hold)
	return
}

// GenerateID = randomly generate id that doesn't start with 0
// Check it doesn't already exist in "account" table
// return = the generated id
func GenerateID() (s string) {
Top:
	var x int
	for i := 0; i < idLength; i++ {
		x = rand.Intn(9)
		for i == 0 && x == 0 {
			x = rand.Intn(9)
		}
		s += strconv.Itoa(x)
	}

	hold := SelectQuery("acc_id", "account", "acc_id", s)
	if hold == "" {
		return
	}
	// Else
	goto Top
}

// AddRecord = insert new record to table
// param1 = identify which table "customer" or "employee"
func AddRecord(who string) {
	var email, pass, first, last string
	fmt.Print("Insert Email: ")
	fmt.Scan(&email)
	fmt.Print("Insert Password: ")
	fmt.Scan(&pass)
	fmt.Print("Insert First Name: ")
	fmt.Scan(&first)
	fmt.Print("Insert Last Name: ")
	fmt.Scan(&last)
	fmt.Println()

	sqlStatement := ``
	if who == "customer" {
		sqlStatement = `
		insert into customer (email, pass, first_name, last_name)
		values ($1, $2, $3, $4)`
	} else {
		sqlStatement = `
		insert into employee (email, pass, first_name, last_name)
		values ($1, $2, $3, $4)`
	}

	_, err := (database.DBCon).Exec(sqlStatement, email, pass, first, last)
	if err != nil {
		panic(err)
	}
}

// DeleteRecord = delete record based on login id
// param1 = identify either "customer" or "employee"
func DeleteRecord(who string) {
	var email string
	fmt.Print("Login ID: ")
	fmt.Scan(&email)

	sqlStatement := ``
	if who == "customer" {
		sqlStatement = `delete from customer where email = $1`
	} else {
		sqlStatement = `delete from employee where email = $1`
	}

	res, err := (database.DBCon).Exec(sqlStatement, email)
	if err == nil {
		count, err := res.RowsAffected()
		if err == nil {
			if count == 0 {
				print.Invalid()
			} else {
				fmt.Println("> Successfully Deleted")
				fmt.Println()
			}
		}
	}
}

// OpenAccount = create new "account" record
// param1 = customer login id
func OpenAccount(login string) {
	var name string
	var balance float32
	fmt.Print("Insert Account Name: ")
	fmt.Scan(&name)
	fmt.Print("Insert Account Balance: $")
	fmt.Scan(&balance)
	fmt.Println()

	sqlStatement := `
	insert into account (email, acc_type, acc_balance, acc_id)
	values ($1, $2, $3, $4)`

	_, err := (database.DBCon).Exec(sqlStatement, login, name, balance, GenerateID())
	if err != nil {
		panic(err)
	}
}

// ApplyJoint = submit a joint account request
// param1 = login id of customer requesting it
func ApplyJoint(login string) {
	var oneNumber string
	var twoNumber string
	fmt.Print("Input Your Account Number: ")
	fmt.Scan(&oneNumber)
	fmt.Print("Input Joint Account Number: ")
	fmt.Scan(&twoNumber)
	fmt.Println()

	if oneNumber == twoNumber {
		print.Invalid()
	} else {
		var hold1 string
		var hold2 string
		var hold3 string
		var hold4 string

		// Get email values
		sqlStatement := `select email from account where acc_id = $1`
		result1 := (database.DBCon).QueryRow(sqlStatement, oneNumber)
		result1.Scan(&hold1)

		result2 := (database.DBCon).QueryRow(sqlStatement, twoNumber)
		result2.Scan(&hold2)

		// Get account names
		sqlStatement2 := `select acc_type from account where acc_id = $1`
		result3 := (database.DBCon).QueryRow(sqlStatement2, oneNumber)
		result3.Scan(&hold3)

		result4 := (database.DBCon).QueryRow(sqlStatement2, hold4)
		result4.Scan(&hold4)

		if hold1 == "" || hold2 == "" || hold1 != login || hold1 == hold2 ||
			hold3 == "JOINT" || hold4 == "JOINT" {
			print.Invalid()
		} else {
			fmt.Println("Submitted Joint Account Request")
			fmt.Println()
			sqlStatement = `
			insert into joint (email1, email2, num1, num2)
			values ($1, $2, $3, $4)`

			_, err := (database.DBCon).Exec(sqlStatement, hold1, hold2, oneNumber, twoNumber)
			if err != nil {
				panic(err)
			}
		}
	}
}

// VerifyJoint = approve/deny customer joint requests
func VerifyJoint() {
	count, slice := print.Joints()
	var input string
	var hold string

	if count != 0 {
		fmt.Print("Input: ")
		fmt.Scan(&input)
		convInput, _ := strconv.Atoi(input)
		newInput := slice[convInput-1]

		sqlStatement := `select index from joint where index = $1`
		result := (database.DBCon).QueryRow(sqlStatement, newInput)
		result.Scan(&hold)

		if hold == "" {
			print.Invalid()
		} else {
			var choice string
			fmt.Println()
			fmt.Println("1) Approve")
			fmt.Println("2) Deny")
			fmt.Print(": ")
			fmt.Scan(&choice)

			switch choice {
			case "1":
				// Get acc_id values
				var idOne, idTwo string
				sqlOne := `select num1 from joint where index = $1`
				sqlTwo := `select num2 from joint where index = $1`
				resOne := (database.DBCon).QueryRow(sqlOne, input)
				resOne.Scan(&idOne)
				resTwo := (database.DBCon).QueryRow(sqlTwo, input)
				resTwo.Scan(&idTwo)

				// Use acc_id values to get acc_balance
				var balOne, balTwo float32
				sqlThree := `select acc_balance from account where acc_id = $1`
				resThree := (database.DBCon).QueryRow(sqlThree, idOne)
				resThree.Scan(&balOne)
				resFour := (database.DBCon).QueryRow(sqlThree, idTwo)
				resFour.Scan(&balTwo)

				// Update the affected records
				var newID string = GenerateID()
				sqlUpdate := `
				update account
				set acc_type = $1, acc_balance = $2, acc_id = $3
				where acc_id = $4`
				_, err := (database.DBCon).Exec(sqlUpdate, "JOINT", balOne+balTwo, newID, idOne)
				if err != nil {
					panic(err)
				}

				_, err = (database.DBCon).Exec(sqlUpdate, "JOINT", balOne+balTwo, newID, idTwo)
				if err != nil {
					panic(err)
				}

				// Delete the joint record now that it's been approved
				print := "> Joint Application Approved\n> Joint Account Number is " + newID
				DeleteJoint(newInput, print)

			case "2":
				DeleteJoint(newInput, "> Joint Application Denied")
			default:
			}
		}
		fmt.Println()
	}
}

// DeleteJoint = deletes record from "joint"
// param1 = index primary key to delete record
// param2 = string message to output
func DeleteJoint(input string, print string) {
	sqlStatement := `delete from joint where index = $1`
	res, err := (database.DBCon).Exec(sqlStatement, input)
	if err == nil {
		count, err := res.RowsAffected()
		if err == nil {
			if count == 0 {
				////////////////////////////////////////////////////////////////////////////
				//print.Invalid()
				fmt.Println("> Invalid Input")
				fmt.Println()
				////////////////////////////////////////////////////////////////////////////
			} else {
				fmt.Println(print)
				fmt.Println()
			}
		}
	}
}
