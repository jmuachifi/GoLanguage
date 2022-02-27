package main

import "fmt"

func main() {
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Println(a)

	fmt.Println(a["brand"])

	fmt.Println("------updating-adding an element------")

	a["year"] = "1970" //Updating an element
	a["color"] = "red" //Adding an element

	fmt.Println(a)

}
