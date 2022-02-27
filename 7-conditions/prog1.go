package main

import "fmt"

func main() {
	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}

	//we can also test variables
	x := 20
	y := 18
	if x > y {
		fmt.Println("x is greater than y")
	}

	//ussing the if else statement
	time := 20
	if time < 18 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	temperature := 14
	if temperature > 15 {
		fmt.Println("It is warm out there")
	} else {
		fmt.Println("It is cold out there")
	}
}
