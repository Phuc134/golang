package main

import "fmt"

func test() {
	fmt.Println("day la test")
}

// paramater
func test1(number1 int, number2 int) {
	fmt.Printf("number1 is %v, number2 is %v", number1, number2)
	fmt.Println()
}

func sum(number1 int, number2 int) int {
	result := number1 + number2
	return result
}

// return multi value
func myFunc(x int, y string) (result int, text string) {
	result = x + x
	text = y + " World"
	return
}
func main() {
	test()
	test1(2, 3)
	fmt.Println(myFunc(2, "Hello"))
	//u can ignore value by use "_"
	a, _ := myFunc(2, "Hello")
	fmt.Println(a)

}
