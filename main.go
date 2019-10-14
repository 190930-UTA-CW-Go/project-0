package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//pending.json is used for all pending accounts waiting approval
//users.json stores all approved

/*
Bank struct is a slice of Users to hold everyone
*/
type Bank struct {
	Bank []User `json:"users"`
}

/*
	User is a struct to represent the user
	Holds the username and password

*/
type User struct {
	Username string  `json:"username"`
	Password string  `json:"password`
	Account  Account `json:account`
	// Approved bool      `json:approved`
}

type Account struct {
	Balance float32 `json:balance`
}

func main() {

	file, err := os.Open("users.json")
	defer file.Close()
	if err != nil {
		fmt.Println("Could not find users.json")

	}
	fmt.Println("Successfully Opened users.json")

	byteValue, _ := ioutil.ReadAll(file)

	var users Bank

	json.Unmarshal(byteValue, &users)

	//print everything in json
	fmt.Println(users)

}

/*
Signup adds a username and password to pending json file

*/
func Signup() {

}

/*


 */
func Login() {

}

/*

 */
func Viewbalance() {

}

/*

 */
func Withdraw() {

}

/*

 */
func Deposit() {

}
