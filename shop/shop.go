package shop

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/190930-UTA-CW-Go/project-0/cfunds"
	"github.com/190930-UTA-CW-Go/project-0/name"
	"github.com/190930-UTA-CW-Go/project-0/stats"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

//Shop for when the player is at the shop
func Shop() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	name := name.GetName()
	fmt.Println("Welcome to the Shop. See anything you like?")
	cfunds := cfunds.Cfunds()
	fmt.Println("Currently holding", cfunds, "Gold.")
	pattack := stats.Pattack()
	phealth := stats.Phealth()
	fmt.Println("Current stats are", pattack, "attack and", phealth, "health.")
	//array with options
	var inventory [7]string
	inventory[0] = "[P]late Mail 30g - adds 60 hp" //add descritpions once they do stuff
	inventory[1] = "[S]hort Sword 10g - adds 2 attack"
	inventory[2] = "[L]ong Sword 30g - adds 6 attack"
	inventory[3] = "[G]auntlets 15g - adds 25 hp"
	inventory[4] = "[B]oots 5g - adds 5 hp"
	inventory[5] = "[C]arrot of Death 300g - adds 900 attack"
	inventory[6] = "[E]xit shop"
	for i, s := range inventory {
		fmt.Println(i, s)
	}
	//fmt.Println(inventory)
	var moneyspend string
	fmt.Scanln(&moneyspend)
	var moneyspendlower = strings.ToLower(moneyspend)
	//need to Check if they can afford it and subtract currently held funds
	if moneyspendlower == "p" {
		if 30 <= cfunds {
			fmt.Println("Aquired Plate Mail")
			db.Exec("UPDATE bankaccount SET cfunds = $1 WHERE name = $2", cfunds-30, name)
			db.Exec("UPDATE bankaccount SET health = $1 WHERE name = $2", phealth+60, name)
			fmt.Println("Thanks for the purchase.")
			Shop()
		} else {
			fmt.Println("You lack the funds buddy")
			Shop()
		}
	} else if moneyspendlower == "s" {
		if 10 <= cfunds {
			fmt.Println("Aquired Short sword")
			db.Exec("UPDATE bankaccount SET cfunds = $1 WHERE name = $2", cfunds-10, name)
			db.Exec("UPDATE bankaccount SET health = $1 WHERE name = $2", pattack+2, name)
			fmt.Println("Thanks for the purchase.")
			Shop()
		} else {
			fmt.Println("You lack the funds buddy")
			Shop()
		}
	} else if moneyspendlower == "l" {
		if 30 <= cfunds {
			fmt.Println("Aquired Long sword")
			db.Exec("UPDATE bankaccount SET cfunds = $1 WHERE name = $2", cfunds-30, name)
			db.Exec("UPDATE bankaccount SET health = $1 WHERE name = $2", pattack+6, name)
			fmt.Println("Thanks for the purchase.")
			Shop()
		} else {
			fmt.Println("You lack the funds buddy")
			Shop()
		}
	} else if moneyspendlower == "g" {
		if 15 <= cfunds {
			fmt.Println("Aquired Gauntlets")
			db.Exec("UPDATE bankaccount SET cfunds = $1 WHERE name = $2", cfunds-15, name)
			db.Exec("UPDATE bankaccount SET health = $1 WHERE name = $2", phealth+25, name)
			fmt.Println("Thanks for the purchase.")
			Shop()
		} else {
			fmt.Println("You lack the funds buddy")
			Shop()
		}
	} else if moneyspendlower == "b" {
		if 5 <= cfunds {
			fmt.Println("Aquired Boots")
			db.Exec("UPDATE bankaccount SET cfunds = $1 WHERE name = $2", cfunds-5, name)
			db.Exec("UPDATE bankaccount SET health = $1 WHERE name = $2", phealth+5, name)
			fmt.Println("Thanks for the purchase.")
			Shop()
		} else {
			fmt.Println("You lack the funds buddy")
			Shop()
		}
	} else if moneyspendlower == "c" {
		if 300 <= cfunds {
			fmt.Println("Aquired Carrot of Death")
			db.Exec("UPDATE bankaccount SET cfunds = $1 WHERE name = $2", cfunds-300, name)
			db.Exec("UPDATE bankaccount SET health = $1 WHERE name = $2", pattack+900, name)
			fmt.Println("Thanks for the purchase.")
			Shop()
		} else {
			fmt.Println("You lack the funds buddy")
			Shop()
		}
	} else if moneyspendlower == "e" {
		fmt.Println("Thats too bad.")
	} else {
		fmt.Println("Clearly you must have hit your head in the arena \ncause I have no idea what your going on about just get out and relearn how to speak.")
	}
}

//add effects to the equipment and clean up text printout
