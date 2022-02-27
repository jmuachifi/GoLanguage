package main

import "fmt"

func myFunction1(x int, y int) int {
	return x + y
}
func myFunction2(x int, y int) (result int) {
	result = x + y
	return //or rerurn result
}

func main() {
	fmt.Println(myFunction1(4, 2))
	fmt.Println(myFunction2(4, 8))

	//store the return value in a variable
	total := myFunction2(20, 25)
	fmt.Println(total)
}
