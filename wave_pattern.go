package main

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
		processBlock(arr[i:end], x)
	}
}

// Process a single block of size (2x+1)
func processBlock(block []int, x int) {
	size := len(block)
	if size <= 1 {
		return
	}

	// Ensure x is not larger than the block size
	effectiveX := x
	if effectiveX >= size {
		effectiveX = size - 1
	}

	// Simple selection sort to sort left part (0 to x-1)
	for i := 0; i < effectiveX; i++ {
		minIndex := i
		for j := i + 1; j < effectiveX; j++ {
			if block[j] < block[minIndex] {
				minIndex = j
			}
		}
		block[i], block[minIndex] = block[minIndex], block[i]
	}

	// Find max element in the block and place it at index x
	if effectiveX < size {
		maxIndex := effectiveX
		for i := effectiveX; i < size; i++ {
			if block[i] > block[maxIndex] {
				maxIndex = i
			}
		}
		block[effectiveX], block[maxIndex] = block[maxIndex], block[effectiveX]

		// Simple selection sort to sort right part (x+1 to size-1) in decreasing order
		for i := effectiveX + 1; i < size; i++ {
			maxIndex := i
			for j := i + 1; j < size; j++ {
				if block[j] > block[maxIndex] {
					maxIndex = j
				}
			}
			block[i], block[maxIndex] = block[maxIndex], block[i]
		}
	}
}
