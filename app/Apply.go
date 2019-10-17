package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Tony-Moon/project-0/gen"
)

/*
Apply documentation
*/
func Apply(db *sql.DB, app string) {
	var acc, pass, comp, first, last string
	var stat, check int

	switch app {
	case "d":
		comp = "Duda-Cola"
	case "s":
		comp = "Salt-PhD"
	case "t":
		comp = "TipsyCo"
	default:
		comp = "Duda-Cola"
	}

	fmt.Println("Welcome to the " + comp + " application page.")

	if len(os.Args) >= 6 {
		fmt.Println("Thank you for choosing the quick apply option.")
		fmt.Println(" ")
		acc, pass, first, last = QuickApply()
	} else if (len(os.Args) > 2) && (len(os.Args) < 6) {
		fmt.Println("It appears that your quick apply was partially or fully incomplete.")
		fmt.Println("Please use the application form below.")
		fmt.Println(" ")
		acc, pass, first, last = Form()
	} else {
		fmt.Println("Please fill out this form to apply.")
		fmt.Println(" ")
		acc, pass, first, last = Form()
	}
	application := NewTech(acc, pass, comp, first, last)
	application.print(comp)

	for i := 0; i == 0; i = i + 0 {
		stat = 0
		check = Check(db, application)
		if check == 1 {
			stat = gen.UseSeed(1, 2)
			i++
		}
		application.result(stat, comp)

		if stat == 0 {
			fmt.Scanln(&acc)
			application.account = acc
		}
	}

	if stat == 2 {
		WriteTo(db, application)
	}

}
