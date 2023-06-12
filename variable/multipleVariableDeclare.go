package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//declare different type
	fmt.Println("DECLARE DIFFERENT TYPE")
	var e, f = 6, "Hello"
	g, h := 7, "World!"

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	//declare in a block
	fmt.Println("DECLARE VARIABLE IN BLOCK")
	var (
		a1 int
		b1 int    = 1
		c1 string = "hello"
	)

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
}
