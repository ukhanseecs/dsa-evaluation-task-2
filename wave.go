package main

import (
	"fmt"
)

// Function to swap two elements
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// Function to sort an array using Bubble Sort (since Go does not allow inbuilt functions)
func bubbleSort(arr []int, start, end bool) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if (start && arr[j] > arr[j+1]) || (!start && arr[j] < arr[j+1]) {
				swap(arr, j, j+1)
			}
		}
	}
}

// Function to rearrange array into wave pattern
func wavePattern(arr []int, x int) {
	size := 2*x + 1
	if len(arr) != size {
		fmt.Println("Array length must be exactly 2x+1")
		return
	}

	// Sort the block to easily identify the maximum element
	bubbleSort(arr, true, true) // Sort entire block in increasing order
	
	// Extract the max element (last element) and put it at index x
	maxIndex := size - 1
	maxElement := arr[maxIndex]
	
	// Shift elements right from index x to size-1
	for i := maxIndex; i > x; i-- {
		arr[i] = arr[i-1]
	}
	arr[x] = maxElement

	// Sort left part in non-decreasing order
	bubbleSort(arr[:x], true, true)

	// Sort right part in non-increasing order
	bubbleSort(arr[x+1:], false, false)
}

func main() {
	arr := []int{3, 6, 5, 10, 7} // Example array with x = 3
	x := 2
	wavePattern(arr, x)
	fmt.Println("Wave Pattern Array:", arr)
}
