package main

import "fmt"

func myFunction1(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}
func main() {
	fmt.Println(myFunction1(14, "Hello"))

	a, b := myFunction1(8, "Hello")
	fmt.Println(a, b)

	/*
		If we (for some reason) do not want to use some of the returned values,
		 we can add an underscore (_), to omit this value.
	*/

	_, w := myFunction1(10, "Hello")
	fmt.Println(w)

	t, _ := myFunction1(12, "Hello")
	fmt.Println(t)
}
