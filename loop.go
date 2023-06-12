package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i == 4 {
			break
		}
		fmt.Println(i)
	}

	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	//RANGE KEYWORD
	fruits = [3]string{"apple", "orange", "banana"}
	for index, value := range fruits {
		fmt.Printf("%v\t%v\n", index, value)
	}
	//if you want ignore index u can replace it into "_"
	fruits = [3]string{"apple", "orange", "banana"}
	for _, value := range fruits {
		fmt.Printf("%v\n", value)
	}
	//similar u can replace it with value
	//for example "for index, _ := range fruits"
}
