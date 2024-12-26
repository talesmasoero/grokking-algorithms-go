//  4.2 Write a recursive function to count the number of items in a list

package main

import "fmt"

func printCounter() {
	fmt.Println("===================================================")
	fmt.Printf("4.2 Write a recursive function to count the number \nof items in a list\n")
	fmt.Println("===================================================")

	arr := []int{5, 4, 3, 2, 1}

	fmt.Println("Array:", arr)
	fmt.Println("Counter:", counter(arr))
	fmt.Println()
}

func counter(arr []int) int {
	// Base case
	if len(arr) == 0 {
		return 0
	}

	// Recursive case
	return 1 + counter(arr[1:])
}
