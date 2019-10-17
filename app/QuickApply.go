package app

import (
	"os"
)

/*
QuickApply documentation
*/
func QuickApply() (string, string, string, string) {
	var acc, pas, fir, las string

	fir = os.Args[3]
	las = os.Args[4]
	acc = os.Args[5]
	pas = os.Args[6]

	return acc, pas, fir, las
}
