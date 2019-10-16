package method

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/gittingdavid/project-0/database"
	"github.com/gittingdavid/project-0/print"
)

// Length of account id number
const idLength = 3

// GenerateID = randomly generate id and check it doesn't already exist
// return = the generated id
func GenerateID() (s string) {
Top:
	var x int
	var hold string
	for i := 0; i < idLength; i++ {
		x = rand.Intn(9)
		for i == 0 && x == 0 {
			x = rand.Intn(9)
		}
		s += strconv.Itoa(x)
	}

	sqlStatement := `select acc_id from account where acc_id = $1`
	row := (database.DBCon).QueryRow(sqlStatement, s)
	row.Scan(&hold)

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
		print.Invalid()
		//panic(err)
	}
}

// DeleteRecord = delete record based on login id
// param1 = identify either "customer" or "employee"
func DeleteRecord(who string) {
	var email, hold string
	fmt.Print("Login ID: ")
	fmt.Scan(&email)

	if email == "user" && who == "employee" {
		print.Invalid()
	} else {
		sql1 := ``
		sql2 := ``
		if who == "customer" {
			sql1 = `select email from customer where email = $1`
			sql2 = `delete from customer where email = $1`
		} else {
			sql1 = `select email from employee where email = $1`
			sql2 = `delete from employee where email = $1`
		}

		result := (database.DBCon).QueryRow(sql1, email)
		result.Scan(&hold)

		if hold == "" {
			print.Invalid()
		} else {

			(database.DBCon).Exec(sql2, email)
			fmt.Println("> Successfully Deleted")
			fmt.Println()
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

	id := GenerateID()
	fmt.Println(">", name, "account", id, "opened")
	fmt.Println()
	_, err := (database.DBCon).Exec(sqlStatement, login, name, balance, id)
	if err != nil {
		panic(err)
	}
}

// ApplyJoint = submit a joint account request
// param1 = customer login id
func ApplyJoint(login string) {
	var acc1, acc2 string
	fmt.Print("Input Your Account Number: ")
	fmt.Scan(&acc1)
	fmt.Print("Input Joint Account Number: ")
	fmt.Scan(&acc2)
	fmt.Println()

	if acc1 == acc2 {
		print.Invalid()
	} else {
		var email1, email2, type1, type2 string

		// Get emails
		sql1 := `select email from account where acc_id = $1`
		result1 := (database.DBCon).QueryRow(sql1, acc1)
		result1.Scan(&email1)
		result2 := (database.DBCon).QueryRow(sql1, acc2)
		result2.Scan(&email2)

		// Get account types
		sql2 := `select acc_type from account where acc_id = $1`
		result3 := (database.DBCon).QueryRow(sql2, acc1)
		result3.Scan(&type1)
		result4 := (database.DBCon).QueryRow(sql2, type2)
		result4.Scan(&type2)

		if email1 == "" || email2 == "" || email1 != login || email1 == email2 ||
			type1 == "JOINT" || type2 == "JOINT" {
			print.Invalid()
		} else {
			fmt.Println("> Submitted Joint Account Request")
			fmt.Println()
			sql := `
			insert into joint (email1, email2, num1, num2)
			values ($1, $2, $3, $4)`

			_, err := (database.DBCon).Exec(sql, email1, email2, acc1, acc2)
			if err != nil {
				panic(err)
			}
		}
	}
}

// VerifyJoint = approve/deny customer joint requests
func VerifyJoint() {
	count, slice := print.Joints()
	var input, hold string

	if count != 0 {
		fmt.Print("Input: ")
		fmt.Scan(&input)
		convInput, err := strconv.Atoi(input)
		if err != nil || convInput < 1 || convInput > count {
			fmt.Println()
			print.Invalid()
		} else {
			newInput := slice[convInput-1]

			sql := `select index from joint where index = $1`
			result := (database.DBCon).QueryRow(sql, newInput)
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
				fmt.Println()

				switch choice {
				case "1":
					// Get acc_id values
					var acc1, acc2 string
					sql1 := `select num1 from joint where index = $1`
					sql2 := `select num2 from joint where index = $1`
					result1 := (database.DBCon).QueryRow(sql1, newInput)
					result1.Scan(&acc1)
					result2 := (database.DBCon).QueryRow(sql2, newInput)
					result2.Scan(&acc2)

					// Use acc_id values to get acc_balance
					var bal1, bal2 float32
					sql3 := `select acc_balance from account where acc_id = $1`
					result3 := (database.DBCon).QueryRow(sql3, acc1)
					result3.Scan(&bal1)
					result4 := (database.DBCon).QueryRow(sql3, acc2)
					result4.Scan(&bal2)

					// Update the affected records
					var newID string = GenerateID()
					sqlUpdate := `
					update account
					set acc_type = $1, acc_balance = $2, acc_id = $3
					where acc_id = $4`
					_, err := (database.DBCon).Exec(sqlUpdate, "JOINT", bal1+bal2, newID, acc1)
					if err != nil {
						panic(err)
					}

					_, err = (database.DBCon).Exec(sqlUpdate, "JOINT", bal1+bal2, newID, acc2)
					if err != nil {
						panic(err)
					}

					// Delete the joint record now that it's been approved
					print := "> Joint Application Approved\n> Joint Account Number is " + newID
					DeleteJoint(newInput, print)

				case "2":
					DeleteJoint(newInput, "> Joint Application Denied")
				default:
					print.Invalid()
				}
			}
			fmt.Println()
		}
	}
}

// DeleteJoint = deletes record from "joint"
// param1 = index primary key to delete record
// param2 = string message to output
func DeleteJoint(input string, message string) {
	var hold string
	sql1 := `select index from joint where index = $1`
	result1 := (database.DBCon).QueryRow(sql1, input)
	result1.Scan(&hold)

	if hold == "" {
		print.Invalid()
	} else {
		sql2 := `delete from joint where index = $1`
		(database.DBCon).Exec(sql2, input)
		fmt.Println(message)
		fmt.Println()
	}
}

// Money = set up withdraw/deposit money process
// param1 = customer login id
// param2 = identify either withdraw or deposit
func Money(login string, who string) {
	var id, amount, balance, check string
	fmt.Print("Insert Account ID: ")
	fmt.Scan(&id)
	fmt.Print("Insert Amount: $")
	fmt.Scan(&amount)
	fmt.Println()

	// Check account id number is valid
	sql := `select email, acc_balance from account where acc_id = $1`
	row := (database.DBCon).QueryRow(sql, id)
	row.Scan(&check, &balance)

	amountInt, flag := Atof(amount)
	balanceInt, _ := Atof(balance)

	if balance == "" || flag == false || login != check {
		print.Invalid()
	} else if who == "withdraw" && amountInt > balanceInt {
		fmt.Println("> Insufficient Funds")
		fmt.Println()
	} else if who == "withdraw" {
		Withdraw(balanceInt-amountInt, id)
		fmt.Println()
	} else if who == "deposit" {
		Deposit(balanceInt+amountInt, id)
		fmt.Println()
	}
}

// Withdraw = execute withdraw sql statement
// param1 = amount to withdraw
// param2 = id to withdraw from
func Withdraw(amount float32, id string) {
	sqlUpdate := `update account set acc_balance = $1 where acc_id = $2`
	_, err := (database.DBCon).Exec(sqlUpdate, amount, id)
	if err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%.2f", amount)
	fmt.Print("> New balance $")
	fmt.Println(s, "in account", id)

}

// Deposit = execute deposit sql statement
// param1 = amount to deposit
// param2 = id to deposit to
func Deposit(amount float32, id string) {
	sqlUpdate := `update account set acc_balance = $1 where acc_id = $2`
	_, err := (database.DBCon).Exec(sqlUpdate, amount, id)
	if err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%.2f", amount)
	fmt.Print("> New balance $")
	fmt.Println(s, "in account", id)
}

// Transfer = set up transfer money process
// param1 = customer login id
func Transfer(login string) {
	var acc1, acc2, transfer string
	fmt.Print("Input Your Account Number: ")
	fmt.Scan(&acc1)
	fmt.Print("Input Transfer Account Number: ")
	fmt.Scan(&acc2)
	fmt.Print("Insert Amount: $")
	fmt.Scan(&transfer)
	fmt.Println()

	if acc1 == acc2 {
		print.Invalid()
	} else {
		var email1, email2 string
		var balance1, balance2 string

		// Check account id number is valid and return acc_balance
		sql := `select email, acc_balance from account where acc_id = $1`
		row1 := (database.DBCon).QueryRow(sql, acc1)
		row1.Scan(&email1, &balance1)

		row2 := (database.DBCon).QueryRow(sql, acc2)
		row2.Scan(&email2, &balance2)

		transferInt, flag := Atof(transfer)
		balance1Int, _ := Atof(balance1)
		balance2Int, _ := Atof(balance2)

		if email1 == "" || email2 == "" || flag == false || login != email1 {
			print.Invalid()
		} else if transferInt > balance1Int {
			fmt.Println("> Insufficient Funds")
			fmt.Println()
		} else {
			Withdraw(balance1Int-transferInt, acc1)
			Deposit(balance2Int+transferInt, acc2)
			fmt.Println()
		}
	}
}

// Atof = convert string to float32
// param1 = string to convert to float32
// return1 = converted amount as float32
// return2 = flag whether there's an error or not
func Atof(s string) (money float32, flag bool) {
	value, err := strconv.ParseFloat(s, 32)
	if err != nil {
		flag = false
	} else {
		flag = true
		money = float32(math.Floor(value*100) / 100)
	}
	return
}
