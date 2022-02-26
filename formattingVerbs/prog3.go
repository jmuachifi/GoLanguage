package main

import "fmt"

func main() {
	var txt = "Hello"
	var i, j = true, false

	fmt.Printf("%s\n", txt)   //Prints the value as plain string
	fmt.Printf("%q\n", txt)   //Prints the value as a double-quoted string
	fmt.Printf("%8s\n", txt)  //Prints the value as plain string (width 8, right justified)
	fmt.Printf("%-8s\n", txt) //Prints the value as plain string (width 8, left justified)
	fmt.Printf("%x\n", txt)   //Prints the value as hex dump of byte values
	fmt.Printf("% x\n", txt)  //Prints the value as hex dump with spaces

	fmt.Println("-----------")
	fmt.Printf("%t\n", i)
	fmt.Printf("%t\n", j)

}
