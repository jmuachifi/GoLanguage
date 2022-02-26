package main

import "fmt"

func main() {
	//Array initialization
	arr1 := [5]int{}              //not initialized
	arr2 := [5]int{1, 2}          //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	fmt.Println("-----------")

	//Initialize only specific elements
	arr4 := [5]int{1: 10, 2: 40}
	fmt.Println(arr4)
}
