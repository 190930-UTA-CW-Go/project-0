package main

import (
	"fmt"

	_ "github.com/Tony-Moon/project-0/gentest"
	_ "github.com/Tony-Moon/project-0/gen"
	_ "github.com/Tony-Moon/project-0/vend"
)

func main() {
	s := gentest.TestGenerate()
	fmt.Println("Test ran successfully")
	fmt.Println(s)
}
