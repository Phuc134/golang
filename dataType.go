package main

import "fmt"

func main() {
	//default of bool is false
	var a bool = true // boolean
	//int default is 0
	//range of int is -2^31 -> 2^31 if 32 bit
	// -2^64 -> 2^64 if 64 bit
	var b int = 3 // Integer
	//uint just have positive number
	var b1 uint = 500
	var c float32 = 3.12 // Floating point number
	//default is ''
	var d string = "Hi!" // String
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(b1)
	fmt.Println(c)
	fmt.Println(d)
}
