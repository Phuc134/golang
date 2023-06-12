package main

import (
	"fmt"
)

func main() {

	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println("------------")
	//If you clare a variable without an initial value it will set value of its type
	var a string
	var b int
	var c bool
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//var and :=
	//var can declare inside and outside function
	// := just can use inside function

}
