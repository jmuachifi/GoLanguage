package main

import "fmt"

func main() {
	var i = 3.14

	fmt.Printf("%e\n", i)    //Scientific notation with 'e' as exponent
	fmt.Printf("%f\n", i)    //Decimal point, no exponent
	fmt.Printf("%.2f\n", i)  //	Default width, precision 2
	fmt.Printf("%6.2f\n", i) //	Width 6, precision 2
	fmt.Printf("%g\n", i)    //Exponent as needed, only necessary digits
}
