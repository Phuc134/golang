package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	fmt.Println(arr1)
	fmt.Println(arr2)
	//DECLARE WITH INFERRED LENGTH
	var _arr1 = [...]int{1, 2, 3}
	_arr2 := [...]int{4, 5, 6, 7, 8}
	fmt.Println(_arr1)
	fmt.Println(_arr2)
	//Access element
	fmt.Println(arr1[0])
	//Change element
	arr1[0] = 10
	fmt.Println(arr1)
	//Initialize only specific element
	arr := [...]int{1: 20, 2: 30, 6: 3}
	fmt.Println(arr)
	//get length of array
	fmt.Println(len(arr))
	//Copy an array into another array
	arrCopy := arr1
	arrCopy[1] = 11
	fmt.Println(arrCopy)
}
