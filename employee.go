package main

import(
	"github.com/project-0/Customer"
	"fmt"
)

type Employee struct {
	name String
}

func test(employee Employee){
	fmt.Println("Employee is: ", Employee.name)
}

