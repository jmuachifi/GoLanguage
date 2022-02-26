package main

import "fmt"

func main() {
	//Find the length of an array
	arr1 := [4]string{"Volvo", "BMW", "Ford", "Mzda"}
	arr2 := [...]int{4, 5, 6, 7, 8, 16}

	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
}
