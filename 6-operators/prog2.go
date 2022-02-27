package main

import "fmt"

func main() {
	//Arithmetic operators
	x := 10
	y := 2

	fmt.Printf("%d + %d = %d\n", x, y, x+y)
	fmt.Printf("%d - %d = %d\n", x, y, x-y)
	fmt.Printf("%d * %d = %d\n", x, y, x*y)
	fmt.Printf("%d / %d = %d\n", x, y, x/y)
	fmt.Printf("%d perc %d = %d\n", x, y, x%y)
	x++
	fmt.Printf("x++ = %d\n", x)
	x--
	fmt.Printf("x-- = %d\n", x)

	var w = 20
	w += 5
	fmt.Println(w)
}
