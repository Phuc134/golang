package main

import "fmt"

func main() {
	// so sanh bang -> ==
	// <=
	// >=
	// !=

	a := 14
	b := 15
	if a < b {
		fmt.Println("a is less than b.")
	} else if a > b {
		fmt.Println("a is more than b.")
	} else {
		fmt.Println("a and b are equal.")
	}
}
