package main

import (
	"fmt"
	"math"
)

/*
func main() {

	var a int = 5
	var b int = 6
	var name string = "manaa"
	const pi float64 = 3.5466544

	fmt.Println("Hello, World!", +(b + a))
	fmt.Println(&a)
	fmt.Println(len(name))
	fmt.Printf("%.3f\n", pi)
	for i := 1; i <= 10; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	age := 18
	/*if age > 18 {
		fmt.Println("yes you can vote")
	} else {
		fmt.Println("no you can not vote")
	}*/
/*switch age {
case 15:
	fmt.Println("don't run after girls")
case 18:
	fmt.Println(" blabla")
}*/
/* var EvenNum [5]int
EvenNum[0] = 0
EvenNum[1] = 2
EvenNum[3] = 3
fmt.Println(EvenNum[3]) */

//EvenNum := [5]int{0, 2, 4, 6, 8}
/* for _, value := range EvenNum {
	fmt.Println(value)
} */
/*
	for i, value := range EvenNum {
		fmt.Println(value, i)
	}

	numSlice := []int{5, 55, 4, 7, 5}
	sliced := numSlice[:5]
	fmt.Println(sliced)
	//
	slices2 := make([]int, 5, 10)
	copy(slices2, numSlice)
	fmt.Println(numSlice)

	//Maps
	studendAge := make(map[string]int)
	studendAge["arry"] = 23
	studendAge["arry1"] = 24
	studendAge["arry2"] = 22
	studendAge["aryy"] = 22
	fmt.Println(studendAge)
	fmt.Println(studendAge["arry"])
	fmt.Println(len(studendAge))
*/

//************************************
//************************************

//function

/* func main() {
	x, y := 5, 6
	fmt.Println(add(x, y))
	//fmt.Println("hello")

}
func add(x int, y int) int {
	return x + y
} */

/* func main() {
	num := 5
	fmt.Println(factorial(num))
}
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
} */

//**************************************
//*********************************

// struct type

/* func main() {
	rect1 := Rectangle{10, 5}
	fmt.Println(rect1.height)
	fmt.Println(rect1.width)
	fmt.Println("the area is:", rect1.area())
}

type Rectangle struct {
	height float64
	width  float64
}

func (rect *Rectangle) area() float64 {
	return rect.height * rect.width

} */

//******************************************************
//*******************************************

//interface

func main() {
	rect := Rectangle{50, 60}
	circ := Circle{7}
	fmt.Println("the area of rec is ", getArea(rect))
	fmt.Println("the area of circle is ", getArea(circ))
}

type Shape interface {
	area() float64
}
type Rectangle struct {
	height float64
	width  float64
}
type Circle struct {
	radius float64
}

func (r1 Rectangle) area() float64 {
	return r1.width * r1.height
}
func (c1 Circle) area() float64 {
	return math.Pow(c1.radius, 2) * math.Pi
}
func getArea(shape Shape) float64 {
	return shape.area()
}

//*****************************************
//******************************************

// files writting and reading

/* func main() {
	file, err := os.Create("samp.txt")

	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("this is a sample text file")
	file.Close()
	//to read the file
	stream, err := ioutil.ReadFile("samp.txt")

	if err != nil {
		log.Fatal(err)
	}
	s1 := string(stream)
	fmt.Println(s1)
} */
//*************************************
//********************************
// simple webserver

/*
 import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/Hello", handler2)
	http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to my homepage\n")
}
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world\n")
} */
