package main

import "fmt"

func main() {
	//Iterating Over Maps
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	/*for k, v := range a {
		fmt.Printf("%v : %v, ", k, v)
	}*/

	fmt.Println()

	//Iterate Over Maps in a Specific Orde

	var b []string //defining the order
	b = append(b, "one", "two", "three", "four")

	for k, v := range a { //loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range b { //loop with the defined order
		fmt.Printf("%v : %v, ", element, a[element])
	}

}
