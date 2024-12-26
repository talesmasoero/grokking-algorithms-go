//  4.2 Write a recursive function to count the number of items in a list

package main

import "fmt"

func printCounter() {
	fmt.Println("===================================================")
	fmt.Printf("4.2 Write a recursive function to count the number \nof items in a list\n")
	fmt.Println("===================================================")

	arr := []int{5, 4, 3, 2, 1}

	fmt.Println("Array:", arr)
	fmt.Println("Counter:", count(arr))
	fmt.Println()
}

func count(arr []int) int {
	// Base cases
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return 1
	}

	var tempArr []int

	tempArr = append(tempArr, arr[1:]...)
	counter := 1

	// Recursive case
	return counter + count(tempArr)
}
