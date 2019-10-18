package navigation

//Attempt at a seperate nav page it works accept you cant have it call on other packages and then have
//those packages call this one after

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/190930-UTA-CW-Go/project-0/bank"
// 	"github.com/190930-UTA-CW-Go/project-0/name"
// )

// //Nav allows the player to choose their destination
// func Nav() string {
// 	name := name.GetName()
// 	var tolocation string
// 	fmt.Println("Where would you like to go", name, "? \n [S]hop [A]rena [B]ank")
// 	fmt.Scanln(&tolocation)
// 	var tolocationlower = strings.ToLower(tolocation)

// 	switch tolocationlower {
// 	case "s":
// 		fmt.Println("Sorry shop not avalible yet")
// 	case "a":
// 		fmt.Println("Sorry arena not avalible yet")
// 	case "b":
// 		bank.Bank()
// 	default:
// 		fmt.Println("Sorry not avalible yet")
// 		// var tolocation string
// 		// fmt.Println("Sorry not an option please choose [S]hop [A]rena [B]ank")
// 		// fmt.Scanln(&tolocation)
// 		// var tolocationlower = strings.ToLower(tolocation)
// 	}
// 	return "nope"
// }
