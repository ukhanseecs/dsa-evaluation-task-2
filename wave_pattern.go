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

	// Sort left part (0 to x) in ascending order
	for i := 0; i < x; i++ {
		minIndex := i
		for j := i + 1; j <= x && j < size; j++ {
			if block[j] < block[minIndex] {
				minIndex = j
			}
		}
		block[i], block[minIndex] = block[minIndex], block[i]
	}

	// Find max element in the block and place it at index x
	maxIndex := x
	for i := x; i < size; i++ {
		if block[i] > block[maxIndex] {
			maxIndex = i
		}
	}
	block[x], block[maxIndex] = block[maxIndex], block[x]

	// Sort right part (x+1 to size-1) in descending order
	for i := x + 1; i < size; i++ {
		maxIndex := i
		for j := i + 1; j < size; j++ {
			if block[j] > block[maxIndex] {
				maxIndex = j
			}
		}
		block[i], block[maxIndex] = block[maxIndex], block[i]
	}
}
