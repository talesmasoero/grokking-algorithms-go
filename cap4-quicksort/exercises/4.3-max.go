//  4.3 Find the maximum number in a list

package main

import (
	"fmt"
)

func printMax() {
	fmt.Println("===================================================")
	fmt.Println("4.3 Find the maximum number in a list")
	fmt.Println()

	arr := []int{4, 1, 5, 3, 2}

	fmt.Println("Array:", arr)
	fmt.Printf("\tMax (Normal Way): %v\n", maxNormalWay(arr))
	fmt.Printf("\tMax (Best Way): %v\n", maxBestWay(arr, len(arr)))
	fmt.Println()
}

func maxNormalWay(arr []int) int {
	// Base case
	if len(arr) == 1 {
		return arr[0]
	}

	// Recursive case
	maxOfRest := maxNormalWay(arr[1:])
	if arr[0] > maxOfRest {
		return arr[0]
	}
	return maxOfRest
}

// This is probably the best way of doing it (but it's AI code, not mine)
// FindMax recursively finds the maximum number in an array
// Parameters:
//
//	arr: The input array of integers
//	n: The current size of the array being considered
//
// Returns:
//
//	The maximum value found in the array
func maxBestWay(arr []int, n int) int {
	// Base case: if we're down to one element, that's our maximum
	if n == 1 {
		return arr[0]
	}

	// Recursive case: compare the last element with the maximum of the rest
	// Get the maximum of the subarray excluding the last element
	maxOfRest := maxBestWay(arr, n-1)

	// Compare the maximum of the subarray with the last element
	if maxOfRest > arr[n-1] {
		return maxOfRest
	}
	return arr[n-1]
}
