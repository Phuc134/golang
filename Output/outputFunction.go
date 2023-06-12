package main

import "fmt"

func main() {
	var i, j string = "hello", "world"
	fmt.Print(i, "\n") //in xong xuong hang
	fmt.Print(j)
	fmt.Print("test", "\n")
	var a, b = 10, 20
	fmt.Println(a, b)

	//Printf
	//%v prints the value
	//%T prints the type
	fmt.Printf("i has value: %v and type: %T ", i, i)
}
