package main

import "fmt"

func main() {
	// + - * / % ++ --
	var a = 3
	var b = 4
	//1100
	//011
	fmt.Println(a + b)  //7
	fmt.Println(a - b)  //-1
	fmt.Println(a * b)  // 12
	fmt.Println(a / b)  // 0
	fmt.Println(a % b)  // 3
	fmt.Println(a << 2) // VD: 3: 0011 ->  2 bit => 1100 -> 12
	fmt.Println(a >> 1) // VD: 3: 0011 -> 1 bit => 0001 -> 1, or 2 bit => 0000 -> 0

	// >, <, >=, <=, !=
	// and -> &&
	// or -> ||
	// not -> !
}
