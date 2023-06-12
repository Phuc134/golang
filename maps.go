package main

import "fmt"

func main() {
	//map luu 1 cap key value va khong trung nhau
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4
	fmt.Println(a)
	fmt.Println(a["brand"])
	//delete element from map
	delete(a, "brand")
	fmt.Println(a)
	fmt.Println(b)
	//check element exist in map
	var _a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	val1, ok1 := _a["brand"] // Checking for existing key and its value
	val2, ok2 := _a["color"] // Checking for non-existing key and its value
	val3, ok3 := _a["day"]   // Checking for existing key and its value
	_, ok4 := _a["model"]    // Only checking for existing key and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	//maps are references
	var test1 = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	test2 := test1

	fmt.Println("test1 ", test1)
	fmt.Println("test2 ", test2)

	test2["year"] = "1970"
	fmt.Println("After change to test2:")

	fmt.Println(test1)
	fmt.Println(test2)

	//iterating over map
	a1 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for k, v := range a1 {
		fmt.Printf("%v : %v, ", k, v)
	}
	var lp = []string{}
	lp = append(lp, "one", "two", "three", "four")
	fmt.Println(lp)

}
