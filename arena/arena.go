package arena

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

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

//Arena for when the player goes to the arena
func Arena() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	cfunds := cfunds.Cfunds()
	name := name.GetName()
	pattack := stats.Pattack()
	phealth := stats.Phealth()
	fmt.Println("Current stats are", pattack, "attack and", phealth, "health.")
	fmt.Println("Sup wanna [f]ight or are you a [c]oward?")
	var fight string
	fmt.Scanln(&fight)
	var fightlower = strings.ToLower(fight)
	if fightlower == "f" {
		fmt.Println("Heck yah, then the funds you are holding will be bet to increase your winnings")
		//Fighter = Fight()
		Fight()
		oattack := Oattack()
		ohealth := Ohealth()
		oname := Oname()
		fmt.Println("Current stats are", pattack, "attack and", phealth, "health.")
		fmt.Println("Opponent is", oname, "and their stats are", oattack, "attack and", ohealth, "health.")
		if pattack+phealth > oattack+ohealth {
			fmt.Println("You win")
			db.Exec("UPDATE bankaccount SET cfunds = $1 WHERE name = $2", cfunds+cfunds+10, name)
		} else {
			fmt.Println("You lose")
			db.Exec("UPDATE bankaccount SET cfunds = $1 WHERE name = $2", cfunds-cfunds, name)
		}
		// cfunds := cfunds.Cfunds()
		// fmt.Println("Currently holding", cfunds, "Gold")
		//option to bet
	} else {
		fmt.Println("Get lost then and stop wasting my time!")
	}
}

//Fight gets player opponent
func Fight() string {
	rand.Seed(time.Now().Unix())
	fighter := []string{
		"LokierFellheart",
		"Bob",
		"BraveSirRobin",
		"TheDude",
		"TheSpanishInquisition",
		"ARandomRabbitThatWanderedIn",
		"YourInnerDemons",
	}
	f := rand.Int() % len(fighter)
	//return fighter[f]
	opponent := fighter[f]
	return opponent
	//fmt.Println(opponent)
}

//Oattack gets opponents attack
func Oattack() int {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	name := Fight()
	var attack int
	rows, err := db.Query("SELECT name, attack FROM bankaccount where name = $1;", name)
	if err != nil {
		// handle err
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &attack)
		if err != nil {
			log.Fatal(err)
		}
		//print to check if its working
		//fmt.Println(name, cfunds)

	}
	return attack
}

//Ohealth gets opponents health
func Ohealth() int {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	name := Fight()
	var health int
	rows, err := db.Query("SELECT name, health FROM bankaccount where name = $1;", name)
	if err != nil {
		// handle err
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &health)
		if err != nil {
			log.Fatal(err)
		}
		//print to check if its working
		//fmt.Println(name, cfunds)

	}
	return health
}

//Oname get oponents name
func Oname() string {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	name := Fight()
	var nameofo string
	rows, err := db.Query("SELECT name FROM bankaccount where name = $1;", name)
	if err != nil {
		// handle err
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&nameofo)
		if err != nil {
			log.Fatal(err)
		}
		//print to check if its working
		//fmt.Println(name, cfunds)

	}
	return nameofo
}

// type Opponent struct {
// 	Name   string
// 	Rating int
// }

// //Fight runs the fight
// func Fight() {
// 	// Map animal names to color strings.
// 	// ... Create a map with composite literal syntax.
// 	fighter := map[string]string{
// 		"Hard":   "LokierFellheart",
// 		"Normal": "Bob",
// 		"Easy":   "BraveSirRobin",
// 	}

// 	// Get color of snake.
// 	f := fighter["1"]

// 	// Display string.
// 	fmt.Println(f)
// }

// func Fight() {
// 	var fighter = []string{"LokierFellheart", "BraveSirRobin", "Bob"}
// 	var fighter = make(map[string]string)
// 	fighterselect = fighter[]

// 	if locafighterselectle == "LokierFellheart" {
// 		greeting = "Hello"
// 	} else if locale == "BraveSirRobin" {
// 		greeting = "Hola"
// 	} else if locale == "Bob" {
// 		greeting = "Guten Tag"
// 	}
// }

// func Fighter() {
// 	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", datasource)
// 	defer db.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// 	name := name.GetName()
// 	// fattack :=
// 	m := make(map[string]int)
// 	m["LokierFellheart"] = 7
// 	m["BraveSirRobin"] = 13
// 	m["Bob"] = 13

// var fighter = []string{"en", "es", "de", "fr"}
// 		if name == "en" {
// 			greeting = "Hello"
// 		} else if locale == "es" {
// 			greeting = "Hola"
// 		} else if locale == "de" {
// 			greeting = "Guten Tag"
// 		} else {
// 			greeting = "Yo"
// 		}
// 	}

//to do

//array of fighters and stats

//acctual fights w/health, attack, energy and actions = attack block rest

//add bet functionality, loss penalty, win reward
//check and be sure they are putting in numbers
