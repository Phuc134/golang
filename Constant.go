package main

import "fmt"

const PI = 3.14 // declare without defined type
const A int = 1 //declare with defined type

// declare with multiple constant
const (
	A1 = 1
	B1 = 3.14
	C1 = "Hi!"
)

func main() {
	fmt.Println(PI)
	fmt.Println(A)

	//multiple constant
	fmt.Println("---------")
	fmt.Println(A1)
	fmt.Println(B1)
	fmt.Println(C1)

}
