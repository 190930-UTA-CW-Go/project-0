package main

import "fmt"

type employee struct {
	customerlist []customers
}

func (e *employee) add(customer1 customers) {
	e.customerlist = append(e.customerlist, customer1)
	fmt.Println("Inside of employee")
}

func (e *employee) List() {
	for _, customers := range e.customerlist {
		fmt.Println(customers.Firstname, customers.Lastname, customers.Username, customers.Password, customers.Balance)
	}
}
