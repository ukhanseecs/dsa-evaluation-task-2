package main

import (
    "fmt"
)

func main() {
	testCases := []struct {
		arr []int
		x   int
	}{
		{[]int{9, 1, 5, 3, 7, 2, 8, 4, 6}, 1},           
		{[]int{10, 20, 30, 40, 50, 60, 70, 80, 90}, 1}, 
		{[]int{3, 6, 2, 9, 1, 7, 4, 8, 5, 10}, 2},      
		{[]int{15, 10, 20, 5, 25, 30, 35, 40, 45, 50, 55, 60, 65, 34}, 3}, 
	}

    for _, testCase := range testCases {
		fmt.Printf("Original Array: %v, X value: %d, Block Size: %d\n", testCase.arr, testCase.x, 2*testCase.x+1)
        waveRearrange(testCase.arr, testCase.x)  // Changed from wavePattern to waveRearrange
        fmt.Println("Rearranged Array:", testCase.arr)
        fmt.Println("---------------------")
    }
}