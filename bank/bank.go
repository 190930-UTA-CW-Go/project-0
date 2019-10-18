package bank

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/190930-UTA-CW-Go/project-0/cfunds"
	"github.com/190930-UTA-CW-Go/project-0/name"
	//"github.com/190930-UTA-CW-Go/project-0/navigation"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

//Account checks for user account
func Account() bool {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	name := name.GetName()
	var funds float64
	rows, err := db.Query("SELECT name, funds FROM bankaccount where name = $1;", name)
	if err != nil {
		// handle err
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &funds)
		if err != nil {
			log.Fatal(err)
		}
		//print to check if its working
		//fmt.Println(name, funds)
		return true
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return false
}

//Bank starts when player enters the Bank
func Bank() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	name := name.GetName()
	cfunds := cfunds.Cfunds()
	fmt.Println("Currently holding", cfunds, "Gold")
	if Account() == true {
		fmt.Println("Welcome back", name, "how can I help you? [W]ithdraw [D]eposit")
		var moneydo string
		fmt.Scanln(&moneydo)
		var moneydolower = strings.ToLower(moneydo)
		if moneydolower == "w" {
			fmt.Println("You currently have", Pbfunds(), "gold in your account \nand you have", cfunds, "on you, how much would you like to withdraw?")
			//gets how much they want to withdraw
			var fundschoice float64
			fmt.Scanln(&fundschoice)

			// fmt.Scanln(&tolocation)
			// var tolocationlower = strings.ToLower(tolocation)

			// switch fundschoice {
			// case "s":
			// 	shop.Shop()
			// case "a":
			// 	arena.Arena()
			// case "b":
			// 	bank.Bank()
			// default:
			// 	fmt.Println("Sorry please choose a valid destination")
			// }
			// check := isNumeric(fundschoice)
			//checks if they have enough money to make the withdraw
			if fundschoice > Pbfunds() {
				// switch fundschoice {
				// case fundschoice > Pbfunds():
				fmt.Println("You do not have", fundschoice, "in your account, get out and earn money!")
				// else if fundschoice < Pbfunds(){
				// 	fmt.Println("Ok here you go. ", name, "withdrew", fundschoice, "from their account")

				// } else {
				// 	fmt.Println("Do you know what numbers are?")
				// }
				//check to see if they acctually put in numbers
				// } else if  fundschoice, err := strconv.ParseInt(userInput, 10, 0)
				// 	if err != nil {
				// } else if Numcheck() {
				// 	fmt.Println("Do you know what numbers are?")
				//} else if Numcheck(fundschoice) == nil {
				// fmt.Println("Do you know what numbers are?")
			} else if fundschoice <= Pbfunds() && fundschoice != 0 {
				//case fundschoice <= Pbfunds(): //try & fundschoice != 0
				fmt.Println("Ok here you go.", name, "withdrew", fundschoice, "gold from their account")
				//Withdraw(name, fundschoice, Pbfunds())
				db.Exec("UPDATE bankaccount SET funds = $1 WHERE name = $2", Pbfunds()-fundschoice, name)
				db.Exec("UPDATE bankaccount SET cfunds = $1 WHERE name = $2", cfunds+fundschoice, name)
				// } else if Numcheck(err == nil) { //need to validate type
				// 	fmt.Println("Please go figure out a proper amount then come back")
			} else {
				fmt.Println("That is not a valid number please get out and learn about math.") //keeps going to next input
				//Bank()
				// var fundschoice float64
				// fmt.Scanln(&fundschoice)
				Fail()
			}

		} else if moneydolower == "d" { //check both w and d to see if they have enough
			fmt.Println("You currently have", Pbfunds(), "gold in your account \nand you have", cfunds, "on you, how much would you like to Deposit?")
			//gets how much they want to withdraw
			var fundschoice float64
			fmt.Scanln(&fundschoice)
			if fundschoice > cfunds {
				fmt.Println("You do not have", fundschoice, "on you, you have", cfunds, "get out and earn money!")
				//checks to see if they have enough funds on them
				// if fundschoice > Pbfunds() {
				// 	fmt.Println("You do not ha
			} else if fundschoice <= cfunds && fundschoice != 0 {
				fmt.Println("Ok", name, "deposited", fundschoice, "gold into their account")
				//Withdraw(name, fundschoice, Pbfunds())
				db.Exec("UPDATE bankaccount SET funds = $1 WHERE name = $2", Pbfunds()+fundschoice, name)
				db.Exec("UPDATE bankaccount SET cfunds = $1 WHERE name = $2", cfunds-fundschoice, name)
			} else {
				fmt.Println("That is not a valid number please get out and learn about math.") //keeps going to next input
				//Bank()
				// var fundschoice float64
				// fmt.Scanln(&fundschoice)
				Fail()
			}
		} else {
			fmt.Println("[W]ithdraw [D]eposit please")
			// var moneydo string
			// fmt.Scanln(&moneydo)
			// var moneydolower = strings.ToLower(moneydo)
		}
	} else {
		fmt.Println("Welcome", name, "Thanks for playing my game, I see you do not have an account with us would you like to set one up? [Y] or [N]")
		//create account
		var acreate string
		fmt.Scanln(&acreate)
		var acreatelower = strings.ToLower(acreate)
		//check for user input
		//fmt.Println(acreatelower)
		if acreatelower == "y" {
			fmt.Println("Alright let me create that for you.\nJust so you know if you store your gold at the bank you will be unable to bet it \nhowever you wont lose it if you are beaten. Also you get 10.5 gold for making an account.")
			//add new account with a 10.5 gold balance under the player name .......remove cfunds
			db.Exec("INSERT INTO bankaccount (name, funds, cfunds, health, attack) values ($1, 10.5, 0, 30, 2)", name)
			//check table
			//getAll(db)
		} else {
			fmt.Println("Well, good luck with that. \nJust so you know if you dont have an account at the bank you will be UNABLE PLAY!!!")
			/*
				var tolocation string
				fmt.Println("Where would you like to go?. \n [S]hop [A]rena [B]ank")
				fmt.Scanln(&tolocation)
				var tolocationlower = strings.ToLower(tolocation)
			*/
		}
	}
}

//remove this func use to check if table is working
// func getAll(db *sql.DB) {
// 	rows, _ := db.Query("SELECT * FROM bankaccount")
// 	for rows.Next() {
// 		var name string
// 		var funds float64
// 		var cfunds float64
// 		var health int
// 		var attack int
// 		rows.Scan(&name, &funds, &cfunds, &health, &attack)
// 		fmt.Println(name, funds, cfunds, health, attack)
// 	}
// }

// //Checkint checks if the user input is an int
// func Checkint(s string) bool {
// 	for _, c := range s {
// 		if !unicode.IsDigit(c) {
// 			return false
// 		}
// 	}
// 	return true
// }

//isIntegral checks that the number given is a float64
// func isIntegral(val float64) bool {
// 	fmt.Println("Printing the thing")
// 	return val == float64(int(val))
// }

// //isIntegral checks that the number given is a float64
// func isNumeric(s string) bool {
// 	_, err := strconv.ParseFloat(s, 64)
// 	return err == nil
// }

// func IsNumeric(s float64) bool {
// 	_, err := strconv.ParseFloat(64)
// 	return err == nil
// }

// func IntCheck() bool {
// 	if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get(fundschoice)); !m {
// 		return false
// 	}
// }

// Pbfunds prints the players bank funds
func Pbfunds() float64 {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	name := name.GetName()
	var funds float64
	rows, err := db.Query("SELECT name, funds FROM bankaccount where name = $1;", name)
	if err != nil {
		// handle err
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &funds)
		if err != nil {
			log.Fatal(err)
		}
		//check to be sure its working
		//fmt.Println(funds)
	}
	return funds
}

//Fail removes bad user input
func Fail() {
	var fail string
	fmt.Scanln(&fail)
}

//Numcheck checks to see if the user input was a number or not
// func Numcheck(s string) bool {
// 	fundschoice, err := strconv.ParseFloat(s, 64)
// 	return err == nil
// }

//to do
//check and be sure that they are putting in numbers
