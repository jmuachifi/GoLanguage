package main

import "fmt"

func main() {
	var i = 15

	fmt.Printf("%b\n", i)   //base 2
	fmt.Printf("%d\n", i)   //base 10
	fmt.Printf("%+d\n", i)  //base 10 and always show sign
	fmt.Printf("%o\n", i)   //base 8
	fmt.Printf("%O\n", i)   // base 8, with leading Oo
	fmt.Printf("%x\n", i)   //base 16, lowercase
	fmt.Printf("%X\n", i)   //base 16, uppercase
	fmt.Printf("%#x\n", i)  //base 16, with leading Ox
	fmt.Printf("%4d\n", i)  //pad with spaces(width 4, right justified)
	fmt.Printf("%-4d\n", i) //pad with spaces (width 4, left justified)
	fmt.Printf("%04d\n", i) // pad with zeros
}
