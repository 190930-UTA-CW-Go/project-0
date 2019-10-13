package main

import (
	"fmt"
)

type employee struct {
	customerlist []account
}

func (e *employee) List() {
	for _, account := range e.customerlist {
		fmt.Println(account.Firstname, account.Lastname, account.Username, account.Password, account.Balance)
	}
}
