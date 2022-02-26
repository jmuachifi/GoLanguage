package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mzda"}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(cars)

	/*
			This example declares two arrays (arr1 and arr2)
			with inferred lengths:

			var arr1 = [...]int{1,2,3}
		    arr2 := [...]int{4,5,6,7,8}

		  	fmt.Println(arr1)
		  	fmt.Println(arr2)

	*/
}
