
# DSA Evaluation Task 2

This repository contains the solution code and explanation for DSA evaluation question #2.

## Getting Started

To get started with this project, follow these steps:

1. Clone the repository
```bash
git clone https://github.com/ukhanseecs/dsa-evaluation-task-2.git
cd dsa-evaluation-task-2
```

2. Run the test cases
```bash
go run .
```
or
```bash
go run main.go wave_pattern.go
```

The program will execute all test cases and display the results.

## Implementation Details

### Problem Breakdown

Given an array of integers and an integer x, we rearrange the array into a wave pattern where:
- The array is divided into blocks of size 2x + 1
- The element at index x must be the maximum in the block
- Elements to the left of x must be in non-decreasing order
- Elements to the right of x must be in non-increasing order

### Solution Approach

#### Block Partitioning
- Process array in chunks of 2x + 1
- Handle smaller end blocks separately

#### Sorting Strategy
- Sort left part (0 to x-1) in ascending order
- Place maximum element at index x
- Sort right part (x+1 to end) in descending order

### Example Output
```
Original Array: [9, 1, 5, 3, 7, 2, 8, 4, 6], x: 1, Block Size: 3
Rearranged Array: [9 5 1 3 7 2 8 6 4]
```

## License
This project is open-source and free to use.

