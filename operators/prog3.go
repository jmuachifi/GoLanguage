package main

import "fmt"

func main() {
	var x = 5
	var y = 3
	fmt.Println(x > y)
	fmt.Println(x != y)
	fmt.Println(x < 5 && x < 1)
	fmt.Println(x < 5 || x < 1)
	fmt.Println(!(x < 5 && x < 1))
}
