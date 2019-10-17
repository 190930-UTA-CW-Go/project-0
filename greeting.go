package main

import "fmt"

//Greeting is a slice that is used to print out different greeting options
func greeting() {
	var grting []string
	grting[0] = "Hi there! Welcome to Piggy Bank! \n"
	grting[1] = "Piggy Bank is here to help you keep track of your savings."
	grting[2] = "Let's get started!"
	grting[3] = "Welcome back!"

	fmt.Println(grting[0], grting[1])
}
