package main

import "fmt"

func main() {
	fmt.Println("---memory efficiency---")
	// Memory efficiency
	fmt.Println("\noriginal slice")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", len(numbers))

	// Create copy with ony needed numbers
	fmt.Println("\nthe new slice only with the numbers needed")
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", len(numbersCopy))
}
