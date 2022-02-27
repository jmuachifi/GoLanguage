package main

import "fmt"

func main() {
	//Access elements of an array
	prices := [3]int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	fmt.Println("-----")

	//Change elements of an array
	prices[2] = 50
	fmt.Println(prices)
}
