package main

import "fmt"

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondeheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	fmt.Println("------------------------------")

	//Creating Maps using make() function

	var s1 = make(map[string]string)
	s1["brand"] = "Ford"
	s1["model"] = "Mustang"
	s1["year"] = "1964"

	var s2 = make(map[string]int)
	s2["Oslo"] = 8
	s2["Bergen"] = 4
	s2["Trondheim"] = 12
	s2["Stavanger"] = 20

	fmt.Printf("s1\t%v\n", s1)
	fmt.Printf("s2\t%v\n", s2)

	fmt.Println("---------------")

	//creating an empty map
	var x1 = make(map[string]string)
	var x2 map[string]string

	fmt.Println(x1 == nil)
	fmt.Println(x2 == nil)
}
