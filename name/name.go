package name

import (
	"fmt"
)

//Name is player name
var Name string

//GetName gets player name and returns it
func GetName() string {
	//get player name
	//var Name string
	if len(Name) > 0 {
		//check to see if it runs
		//fmt.Println("Sup")
	} else {
		fmt.Println("Hello adventurer please state your name no spaces dashes or underscores.")
		fmt.Scanln(&Name)
		// in := bufio.NewReader(os.Stdin)
		// Name, _ := in.ReadString('\n')
	}
	return Name
}
