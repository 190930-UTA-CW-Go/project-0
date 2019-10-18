// package main

// import (
// 	"database/sql"
// 	"fmt"

// 	"github.com/jeinostroza/projecttester/project-0/register"
// 	_ "github.com/lib/pq"
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "postgres"
// 	dbname   = "postgres"
// )

// //ClientMenu show the client's menu
// // func ClientMenu() {
// // 	fmt.Println("WELCOME TO THE BANK")
// // 	fmt.Println("")
// // 	fmt.Println("MENU")
// // 	var clientM = make(map[int]string)
// // 	clientM[1] = "1.- Apply to open and account"
// // 	clientM[2] = "2.- Deposit"
// // 	clientM[3] = "3.- Withdraw"
// // 	clientM[4] = "4.- Transfer"
// // 	clientM[5] = "5.- Balance"

// // 	var i int
// // 	for i = 0; i <= len(clientM); i++ {
// // 		fmt.Println(clientM[i])
// // 	}
// // }

// //EmplMenu is the employee menu
// func EmplMenu() {
// 	fmt.Println("EMPLOYEE MENU")
// 	fmt.Println("")
// 	var empMenu = make(map[int]string)
// 	empMenu[1] = "1.- Customer Information(by one)"
// 	empMenu[2] = "2.- Customer Information(all)"
// 	empMenu[3] = "3.- Applications"
// 	empMenu[4] = "4.- Joint Accounts"

// 	var i int
// 	for i = 0; i <= len(empMenu); i++ {
// 		fmt.Println(empMenu[i])
// 	}
// }

// //PMenu is a menu
// func PMenu() {
// 	fmt.Println("WELCOME TO GO BANK INC.")
// 	fmt.Println("")
// 	fmt.Println("MAIN MENU")
// 	fmt.Println("")
// 	var princ = make(map[int]string)
// 	princ[1] = "1.- Customer"
// 	princ[2] = "2.- Employee"

// 	var i int
// 	for i = 0; i <= len(princ); i++ {
// 		fmt.Println(princ[i])
// 	}
// }

// //Customer is the structure of bank's customer
// type Customer struct {
// 	firstname, lastname, street, city, statec, zip, email, username, pass string
// }

// //GetAll show all the DB
// func GetAll(db *sql.DB) {
// 	result, _ := db.Query("select firstname, lastname, street, city, statec, zip, email, username, pass, montlyincomes, monthlyexpenses from client")
// 	for result.Next() {
// 		var firstname, lastname, street, city, statec, zip, email, username, pass string
// 		var montlyincomes, monthlyexpenses float64
// 		result.Scan(&firstname, &lastname, &street, &city, &statec, &zip, &email, &username, &pass, &montlyincomes, &monthlyexpenses)
// 		fmt.Print("First name: ")
// 		fmt.Println(firstname)
// 		fmt.Print("Last name: ")
// 		fmt.Println(lastname)
// 		fmt.Print("Street: ")
// 		fmt.Println(street)
// 		fmt.Print("City: ")
// 		fmt.Println(city)
// 		fmt.Print("State: ")
// 		fmt.Println(statec)
// 		fmt.Print("Zip code: ")
// 		fmt.Println(zip)
// 		fmt.Print("Email: ")
// 		fmt.Println(email)
// 		fmt.Print("Username: ")
// 		fmt.Println(username)
// 		fmt.Print("Monthly incomes: ")
// 		fmt.Println(montlyincomes)
// 		fmt.Print("Monthly Expenses: ")
// 		fmt.Println(monthlyexpenses)
// 		fmt.Println("========================")
// 	}
// }

// //GetAllClient return the client's information
// func GetAllClient(db *sql.DB, usernameget string) {
// 	row := db.QueryRow("select firstname, lastname, street, city, statec, zip, email, montlyincomes, monthlyexpenses from client where username = $1", usernameget)
// 	var firstname, lastname, street, city, statec, zip, email string
// 	var montlyincomes, monthlyexpenses float64
// 	row.Scan(&firstname, &lastname, &street, &city, &statec, &zip, &email, &montlyincomes, &monthlyexpenses)
// 	fmt.Print("First name: ")
// 	fmt.Println(firstname)
// 	fmt.Print("Last name: ")
// 	fmt.Println(lastname)
// 	fmt.Print("Street: ")
// 	fmt.Println(street)
// 	fmt.Print("City: ")
// 	fmt.Println(city)
// 	fmt.Print("State: ")
// 	fmt.Println(statec)
// 	fmt.Print("Zip code: ")
// 	fmt.Println(zip)
// 	fmt.Print("Email: ")
// 	fmt.Println(email)
// 	fmt.Print("Monthly incomes: ")
// 	fmt.Println(montlyincomes)
// 	fmt.Print("Monthly Expenses: ")
// 	fmt.Println(monthlyexpenses)

// }

// //SearchByUsername validate if a username exist
// func SearchByUsername(db *sql.DB, searchvalue string) string {
// 	row := db.QueryRow("select username from client where username = $1", searchvalue)
// 	var username string
// 	row.Scan(&username)
// 	return username

// }

// //SearchByNameLastname return the username
// func SearchByNameLastname(db *sql.DB, searchvalue string, searchvalue2 string) string {
// 	row := db.QueryRow("select username from client where firstname = $1 and lastname = $2", searchvalue, searchvalue2)
// 	var username string
// 	row.Scan(&username)
// 	return username

// }

// //SearchByPass return the password
// func SearchByPass(db *sql.DB, searchvalue1 string) string {
// 	row := db.QueryRow("select pass from client where pass = $1", searchvalue1)
// 	var pass string
// 	row.Scan(&pass)
// 	return pass

// }

// //UpdateIncomes update the monthly income record
// func UpdateIncomes(db *sql.DB, income float32, username string) {
// 	row := db.QueryRow(`update client set montlyincomes = $1 where username = $2`, income, username)
// 	var incomeup float32
// 	row.Scan(&incomeup)

// }

// //UpdateExpenses update the monthly expenses record
// func UpdateExpenses(db *sql.DB, expenses float32, username string) {
// 	row := db.QueryRow(`update client set monthlyexpenses = $1 where username = $2`, expenses, username)
// 	var expensesup float32
// 	row.Scan(&expensesup)

// }

// //ClientApplying is a query that show the client that are apply
// func ClientApplying(db *sql.DB) {
// 	row, _ := db.Query("select client_id, firstname, lastname, montlyincomes, monthlyexpenses, approve from client where montlyincomes > 1 and monthlyexpenses >1 and approve is null")
// 	for row.Next() {
// 		var clientid int
// 		var approve string
// 		var firstname string
// 		var lastname string
// 		var montlyincomes, monthlyexpenses float64
// 		row.Scan(&clientid, &firstname, &lastname, &montlyincomes, &monthlyexpenses, &approve)
// 		fmt.Print("Client ID: ")
// 		fmt.Println(clientid)
// 		fmt.Print("First Name: ")
// 		fmt.Println(firstname)
// 		fmt.Print("Last Name: ")
// 		fmt.Println(lastname)
// 		fmt.Print("Monthly Incomes: ")
// 		fmt.Println(montlyincomes)
// 		fmt.Print("Monthly Expenses: ")
// 		fmt.Println(monthlyexpenses)
// 		fmt.Print("Aproved: ")
// 		fmt.Println(approve)
// 		fmt.Println("====================")
// 	}
// }

// //Clientcheck insert info in aproved clients
// func Clientcheck(db *sql.DB, id int) {
// 	row, _ := db.Exec("insert into CheckingAccoClient(client_id) values($1)", id)
// 	fmt.Println(row)
// }

// //Transaction to the account
// func Transaction(db *sql.DB, id int, cha int, amount float64) {
// 	row, _ := db.Exec("insert into accounts(client_id, checkingaccount, amount) values($1,$2,$3)", id, cha, amount)
// 	fmt.Println(row)

// }

// //ApprovedClient change the client's status
// func ApprovedClient(db *sql.DB, des string, id int) {
// 	row := db.QueryRow(`update client set approve = $1 where client_id =$2`, des, id)
// 	var appro int
// 	row.Scan(&appro)
// }

// //Getaccountnum return the acocunt number
// func Getaccountnum(db *sql.DB, id int) int {
// 	row := db.QueryRow("select checkingaccount from CheckingAccoClient where client_id = $1", id)
// 	var chaccount int
// 	row.Scan(&chaccount)
// 	return chaccount
// }

// //GetID return the client's id
// func GetID(db *sql.DB, usernameget string) int {
// 	row := db.QueryRow("select client_id from client where username = $1", usernameget)
// 	var clientid int
// 	row.Scan(&clientid)
// 	return clientid
// }

// //Getbalance return the client's balance
// func Getbalance(db *sql.DB, id int) float64 {
// 	row := db.QueryRow("select sum(amount) from accounts where client_id = $1", id)
// 	var bal float64
// 	row.Scan(&bal)
// 	return bal
// }

// //JointAccount to jooint accounts
// func JointAccount(db *sql.DB, jointid int, first string, last string) {
// 	row := db.QueryRow("update client set joint = $1 where firstname = $2 and lastname=$3", jointid, first, last)
// 	var joint1 int
// 	row.Scan(&joint1)
// }

// //GetJoint return the joint code
// func GetJoint(db *sql.DB, usernamet string) int {
// 	row := db.QueryRow("select joint from client where username= $1", usernamet)
// 	var unt int
// 	row.Scan(&unt)
// 	return unt
// }

// func increment(value *int) int {
// 	*value++
// 	return *value
// }

// func main() {
// 	register.Website()
// 	var selec int          // store the option of the main menu
// 	var option int         //store the option of the client menu
// 	var userN string       //store the username type in the login - option 1
// 	var passU string       //store the password type in the login - option 1
// 	var opt2 string = "y"  //store y or n if you want to do something else
// 	var usersearch string  //store the result of the search for username
// 	var passsearch string  //store the result of the search for password
// 	var empoption int      //store the option of the employee menu
// 	var firstnemp string   //store the name that the employee is searching
// 	var lastnemp string    // store the name that the employee is searching
// 	var usernameEmp string //store the username to pull the infor
// 	var approvesel int     //store the option for aproved client
// 	var opt3 string = "y"  //store the option for employees opt3
// 	var monExp float32
// 	var montlyincomes float32
// 	var damount float64 //store the deposit's amount
// 	var id int
// 	var accnumb int
// 	var wamount float64
// 	var wamountf float64
// 	var balance float64
// 	var employeeuser string
// 	var employeeuser1 string = "employee"
// 	var employeepass string
// 	var employeepass1 string = "123456"
// 	var jointfirst1, jointfirst2, jointlast1, jointlast2 string
// 	var jointoption string
// 	var jointid int = 0
// 	var jointid1 int
// 	var jointcode1 int
// 	var jointcode2 int
// 	var firstnamet string
// 	var lastnamet string
// 	var usert string
// 	var amountt, amounttc, finalbalance float64
// 	var idtransf1, idtransf2, accountt1, accountt2 int
// 	var idclientbalance int

// 	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", datasource)
// 	defer db.Close()
// 	if err != nil {
// 		panic(err)
// 	}

// 	PMenu()
// 	fmt.Println("")
// 	fmt.Println("Please select an option")
// 	fmt.Scanln(&selec)
// 	for opt3 == "y" {
// 		if selec == 1 {
// 			for opt2 == "y" {
// 				ClientMenu()
// 				fmt.Println("What do you want to do?")
// 				fmt.Scanln(&option)

// 				switch option {
// 				case 1:
// 					fmt.Println("Please enter your user name")
// 					fmt.Scanln(&userN)
// 					usersearch = SearchByUsername(db, userN)
// 					if usersearch == userN {
// 						fmt.Println("Please enter your password")
// 						fmt.Scanln(&passU)
// 						passsearch = SearchByPass(db, passU)
// 						SearchByUsername(db, passU)
// 						if passsearch == passU {
// 							id = GetID(db, userN)
// 							accnumb = Getaccountnum(db, id)
// 							if accnumb == 0 {
// 								fmt.Println("Please enter your monthly incomes:")
// 								fmt.Scanln(&montlyincomes)
// 								UpdateIncomes(db, montlyincomes, userN)
// 								fmt.Println("Please enter your mountly expenses:")
// 								fmt.Scanln(&monExp)
// 								UpdateExpenses(db, monExp, userN)
// 								fmt.Println("")
// 								fmt.Println("Thank you for your application,")
// 								fmt.Println("We are reviewing your information")
// 							} else {
// 								fmt.Println("You already apply for an account.")
// 							}

// 						} else {
// 							fmt.Println("Wrong Password")
// 						}
// 					} else {
// 						fmt.Println("Wrong Username")
// 					}
// 				case 2:
// 					fmt.Println("")
// 					fmt.Println("")
// 					fmt.Println("DEPOSIT")
// 					fmt.Println("")
// 					fmt.Println("Please enter your user name")
// 					fmt.Scanln(&userN)
// 					usersearch = SearchByUsername(db, userN)
// 					if usersearch == userN {
// 						fmt.Println("Please enter your password")
// 						fmt.Scanln(&passU)
// 						passsearch = SearchByPass(db, passU)
// 						SearchByUsername(db, passU)
// 						if passsearch == passU {
// 							fmt.Println("========================")
// 							//fmt.Println("DEPOSIT")
// 							fmt.Println("")
// 							fmt.Println("Deposit amount: ")
// 							fmt.Scanln(&damount)
// 							id = GetID(db, userN)
// 							fmt.Println(id)
// 							accnumb = Getaccountnum(db, id)
// 							fmt.Println(accnumb)
// 							Transaction(db, id, accnumb, damount)
// 							balance = Getbalance(db, id)
// 							fmt.Print("Your balance is ")
// 							fmt.Println(balance)
// 						} else {
// 							fmt.Println("Wrong Password")
// 						}
// 					} else {
// 						fmt.Println("Wrong Username")
// 					}
// 				case 3:
// 					fmt.Println("Please enter your username")
// 					fmt.Scanln(&userN)
// 					usersearch = SearchByUsername(db, userN)
// 					if usersearch == userN {
// 						fmt.Println("Please enter your password")
// 						fmt.Scanln(&passU)
// 						passsearch = SearchByPass(db, passU)
// 						SearchByUsername(db, passU)
// 						if passsearch == passU {
// 							fmt.Println("")
// 							fmt.Println("")
// 							fmt.Println("")
// 							fmt.Println("WITHDRAW")
// 							fmt.Println("")
// 							fmt.Println("Withdraw amount")
// 							fmt.Scanln(&wamount)
// 							wamountf = (wamount * -1)
// 							id = GetID(db, userN)
// 							accnumb = Getaccountnum(db, id)
// 							Transaction(db, id, accnumb, wamountf)
// 							balance = Getbalance(db, id)
// 							fmt.Print("Your new balance is ")
// 							fmt.Println(balance)
// 						} else {
// 							fmt.Println("Wrong Password")
// 						}
// 					} else {
// 						fmt.Println("Wrong Username")
// 					}

// 				case 4:
// 					fmt.Println("")
// 					fmt.Println("TRANFER MONEY")
// 					fmt.Println("")
// 					fmt.Println("Please enter your username")
// 					fmt.Scanln(&userN)
// 					usersearch = SearchByUsername(db, userN)
// 					if usersearch == userN {
// 						fmt.Println("Please enter your password")
// 						fmt.Scanln(&passU)
// 						passsearch = SearchByPass(db, passU)
// 						SearchByUsername(db, passU)
// 						if passsearch == passU {
// 							jointcode1 = GetJoint(db, userN)
// 							idtransf1 = GetID(db, userN)
// 							accountt1 = Getaccountnum(db, idtransf1)

// 							fmt.Println("TRANSFER TO:")
// 							fmt.Println("First name:")
// 							fmt.Scan(&firstnamet)
// 							fmt.Println("Last name:")
// 							fmt.Scan(&lastnamet)
// 							usert = SearchByNameLastname(db, firstnamet, lastnamet)
// 							jointcode2 = GetJoint(db, usert)
// 							idtransf2 = GetID(db, usert)
// 							accountt2 = Getaccountnum(db, idtransf2)
// 							fmt.Println(jointcode2)
// 							if jointcode1 == jointcode2 {
// 								fmt.Println("Tansfer amount")
// 								fmt.Scan(&amountt)
// 								amounttc = amountt * -1
// 								Transaction(db, idtransf2, accountt2, amountt)
// 								Transaction(db, idtransf1, accountt1, amounttc)
// 							} else {
// 								fmt.Println("Sorry, you can't transfer money to this person. Please call your bank to joint the account")
// 							}
// 						} else {
// 							fmt.Println("Wrong Password")
// 						}
// 					} else {
// 						fmt.Println("Wrong Username")

// 					}
// 				case 5:
// 					fmt.Println("")
// 					fmt.Println("BALANCE")
// 					fmt.Println("")
// 					fmt.Println("Please enter your username")
// 					fmt.Scanln(&userN)
// 					usersearch = SearchByUsername(db, userN)
// 					if usersearch == userN {
// 						fmt.Println("Please enter your password")
// 						fmt.Scanln(&passU)
// 						passsearch = SearchByPass(db, passU)
// 						SearchByUsername(db, passU)
// 						if passsearch == passU {
// 							fmt.Println("Your current balance is")
// 							idclientbalance = GetID(db, userN)
// 							finalbalance = Getbalance(db, idclientbalance)
// 							fmt.Println(finalbalance)
// 						} else {
// 							fmt.Println("Wrong Password")
// 						}
// 					} else {
// 						fmt.Println("Wrong Username")

// 					}

// 				}
// 				fmt.Println("Do you want to continue?")
// 				fmt.Scanln(&opt2)

// 			}
// 		} else if selec == 2 {
// 			fmt.Println("USER NAME:")
// 			fmt.Scanln(&employeeuser)
// 			if employeeuser == employeeuser1 {
// 				fmt.Println("PASSWORD:")
// 				fmt.Scanln(&employeepass)
// 				if employeepass == employeepass1 {
// 					EmplMenu()
// 					fmt.Println("")
// 					fmt.Println("Please select an option")
// 					fmt.Scanln(&empoption)
// 					switch empoption {
// 					case 1:
// 						fmt.Println("Enter Customers's first name")
// 						fmt.Scanln(&firstnemp)
// 						fmt.Println("Enter Customers's last name")
// 						fmt.Scanln(&lastnemp)
// 						usernameEmp = SearchByNameLastname(db, firstnemp, lastnemp)
// 						fmt.Println("")
// 						fmt.Println("==============================")
// 						fmt.Println("CLIENT INFORMATION")
// 						fmt.Println("")
// 						GetAllClient(db, usernameEmp)
// 					case 2:
// 						fmt.Println("")
// 						fmt.Println("==============================")
// 						fmt.Println("CLIENT INFORMATION")
// 						fmt.Println("")
// 						GetAll(db)
// 					case 3:
// 						fmt.Println("CLIENTS APPLYING FOR A CHECKING ACCOUNT")
// 						fmt.Println("")
// 						ClientApplying(db)
// 						fmt.Println("")
// 						fmt.Println("Please enter de ID of the client that you want to approve(or press 0): ")
// 						fmt.Scanln(&approvesel)
// 						if approvesel != 0 {
// 							ApprovedClient(db, "yes", approvesel)
// 							Clientcheck(db, approvesel)
// 							fmt.Println("Checking account was succesfully created")
// 						} else {
// 							fmt.Println("")
// 						}
// 					case 4:
// 						fmt.Println("JOINT ACCOUNTS")
// 						fmt.Println("")
// 						fmt.Println("===================")
// 						fmt.Println("First Client")
// 						fmt.Println("")
// 						fmt.Println("First Name")
// 						fmt.Scan(&jointfirst1)
// 						fmt.Println("Last Name")
// 						fmt.Scan(&jointlast1)
// 						fmt.Println("=================")
// 						fmt.Println("Second Client")
// 						fmt.Println("============")
// 						fmt.Println("First Name")
// 						fmt.Scan(&jointfirst2)
// 						fmt.Println("Last Name")
// 						fmt.Scan(&jointlast2)
// 						fmt.Println("Are you sure that you want to joint this accounts?(y/n)")
// 						fmt.Scan(&jointoption)
// 						if jointoption == "y" {
// 							jointid1 = increment(&jointid)
// 							JointAccount(db, jointid1, jointfirst1, jointlast1)
// 							JointAccount(db, jointid1, jointfirst2, jointlast2)
// 							fmt.Println("Transaction Succesful")
// 						}
// 					}
// 				} else {
// 					fmt.Println("Wrong password")
// 				}
// 			} else {
// 				fmt.Println("Wrong Username")
// 			}
// 		}
// 		fmt.Println("Do you want to continue:")
// 		fmt.Scanln(&opt3)

// 	}
// }
