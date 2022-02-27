package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

//Pass Struct as Function Arguments

func main() {
	var pers1 Person
	var pers2 Person

	//Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	//Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 45000

	//Access and print Pers1 info
	fmt.Println("Person1 Information:")
	printPerson(pers1)

	fmt.Println("--------------")

	//Access and print Pers2 info
	fmt.Println("Person2 Information:")
	printPerson(pers2)

}
func printPerson(pers Person) {
	fmt.Println("Name:", pers.name)
	fmt.Println("Age:", pers.age)
	fmt.Println("Job:", pers.job)
	fmt.Println("Salary:", pers.salary)
}
