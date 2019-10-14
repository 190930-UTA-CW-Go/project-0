package app

import "fmt"

/*
Tech documentation
*/
type Tech struct {
	account   string
	password  string
	company   string
	firstname string
	lastname  string
}

func (te Tech) print() {
	fmt.Println("Welcome", te.firstname)
}
