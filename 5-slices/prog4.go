package main

import "fmt"

func main() {
	//Acces elements of a slice
	prices := []int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	fmt.Println("---")

	//Change elements of a slice
	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	fmt.Println("---append elements to a slice---")

	//Append elements to a slice
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice2 := append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))

	fmt.Println("---append one slice to another slide---")

	// Append one slice to another slice
	myslice3 := []int{1, 2, 3}
	myslice4 := []int{4, 5, 6}
	myslice5 := append(myslice3, myslice4...)

	fmt.Printf("myslice5 = %v\n", myslice5)
	fmt.Printf("length = %d\n", len(myslice5))
	fmt.Printf("capacity = %d\n", cap(myslice5))

	fmt.Println("---change the length of a slice---")

	//Change the length of a slice
	arr1 := [6]int{9, 10, 11, 12, 13, 14}
	myslice6 := arr1[1:5]
	fmt.Printf("myslice6 = %v\n", myslice6)
	fmt.Printf("length = %d\n", len(myslice6))
	fmt.Printf("capacity = %d\n", cap(myslice6))

	myslice6 = arr1[1:3] //change length by re-slicing the array
	fmt.Printf("myslice6 = %v\n", myslice6)
	fmt.Printf("length = %d\n", len(myslice6))
	fmt.Printf("capacity = %d\n", cap(myslice6))

	myslice6 = append(myslice6, 20, 21, 22, 23) //change length by appending items
	fmt.Printf("myslice6 = %v\n", myslice6)
	fmt.Printf("length = %d\n", len(myslice6))
	fmt.Printf("capacity = %d\n", cap(myslice6))

}
