package main

import "fmt"

func main() {

	/*
		The Printf() function first formats its
		argument based on the given formatting verb and then prints them

	*/
	var i string = "Hello"
	var j int = 25

	fmt.Printf("i has value: %v and type: %T \n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)

}
