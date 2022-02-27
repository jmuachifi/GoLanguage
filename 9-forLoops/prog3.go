package main

import "fmt"

func main() {
	//nested loops
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	/*
		The range keyword is used to more easily iterate over an array, slice or map.
		It returns both the index and the value
	*/

	fmt.Println("-------------")

	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	fmt.Println("-------")

	/*
		Tip: To only show the value or the index,
		you can omit the other output using an underscore (_).
	*/

	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}

	fmt.Println("-------")

	/*
		Here, we want to omit the values
		(idx stores the index, val stores the value):
	*/
	for idx, _ := range fruits {
		fmt.Printf("%v\n", idx)
	}
}
