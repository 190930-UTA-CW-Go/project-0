package main

import "fmt"

type employee struct {
	customerlist []customers
}

func (e *employee) List() {
	for _, customers := range e.customerlist {
		fmt.Println(customers.Firstname, customers.Lastname, customers.Username, customers.Password, customers.Balance)
	}
}
