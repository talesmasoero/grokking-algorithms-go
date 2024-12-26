//  4.1 Write out the code for the earlier sum function

package main

import "fmt"

func printSum() {
	fmt.Println("===================================================")
	fmt.Println("4.1 Write out the code for the earlier sum function")
	fmt.Println("===================================================")

	arr := []int{1, 2, 3, 4, 5}

	fmt.Println("Array:", arr)
	fmt.Println("Sum:", sum(arr))
	fmt.Println()
}

func sum(arr []int) int {
	// Base cases
	switch len(arr) {
	case 0:
		return 0
	case 1:
		return arr[0]
	}

	var tempArr []int

	first := arr[0]
	tempArr = append(tempArr, arr[1:]...)

	// Recursive case
	return first + sum(tempArr)
}
