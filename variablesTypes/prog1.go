package main

import "fmt"

// declaring variables outsides of a function
var a int
var b int = 8
var c = 16

func main() {
	a = 1

	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	/*
		Variable Declaration without initial value
		var a string
		var b int
		var c bool

		fmt.Println(a) //result nathing or ""
		fmt.Println(b) // result zero
		fmt.Println(c) // result false
	*/

	/*
		The value can be assign after declaration
		Example:
		var student1 string
		student1 = "John"
		fmt.Println(student1)
	*/
}
