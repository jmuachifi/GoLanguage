package main

import "fmt"

func main() {
	//Create a slice from an array
	arr1 := [6]int{10, 11, 12, 13, 14, 15} //create the array
	myslice := arr1[2:4]                   //a slice made from the array

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
}
