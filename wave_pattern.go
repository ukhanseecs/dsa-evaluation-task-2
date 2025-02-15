package main

import (
	"sort"
)

// Function to rearrange array in wave pattern
func waveRearrange(arr []int, x int) {
	blockSize := 2*x + 1
	n := len(arr)

	// Process each block of size (2x+1)
	for i := 0; i < n; i += blockSize {
		end := i + blockSize
		if end > n {
			end = n
		}
		processBlock(arr, i, end, x)
	}
}

// Process a single block of size (2x+1) in-place
func processBlock(arr []int, start, end, x int) {
	size := end - start
	if size <= 1 {
		return
	}

	// Sort the block
	sort.Slice(arr[start:end], func(i, j int) bool {
		return arr[i] < arr[j]
	})

	// Swap the middle element with the last element
	mid := start + x
	maxIdx := end - 1
	arr[mid], arr[maxIdx] = arr[maxIdx], arr[mid]
}
