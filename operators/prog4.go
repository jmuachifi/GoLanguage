package main

import "fmt"

func main() {
	var x = 9
	var y = 8
	fmt.Printf("x = %b\n", x)
	fmt.Printf("x & y is %b\n", x&y)
	fmt.Printf("x | y is %b\n", x|y)
	fmt.Printf("x ^ y is %b\n", x^y)
	fmt.Printf("x << 2 is %b\n", x<<2)
	fmt.Printf("x >> 2 is %b\n", x>>2)
}
