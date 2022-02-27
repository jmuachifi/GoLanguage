package main

import "fmt"

func main() {
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Println(a)

	//Remove element from Map
	fmt.Println("-------Remove element--------")

	delete(a, "year")

	fmt.Println(a)
}
