package main

import (
	"fmt"

	"github.com/Tony-Moon/project-0/gen"
	_ "github.com/Tony-Moon/project-0/vend"
)

func main() {
	index, beverage, stock := gen.Generate(3, 5, 10)
	fmt.Println(index)
	fmt.Println(beverage)
	fmt.Println(stock)
}
