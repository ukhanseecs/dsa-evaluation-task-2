# Wave Pattern : DSA Evaluation Task 2

This repository contains the solution code and explanation for DSA evaluation question #2.

## Getting Started

To get started with this project, follow these steps:

1. Clone the repository
    ```bash
    git clone https://github.com/ukhanseecs/dsa-evaluation-task-2.git
    cd dsa-evaluation-task-2
    go mod tidy
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

## Project Structure

```
dsa-evaluation-task-2/
├── main.go              # Entry point of the program
├── wave_pattern.go      # Contains wave pattern implementation
├── wave_pattern_test.go # Test cases for wave pattern
└── README.md           # Project documentation
```

### File Descriptions

- **main.go**: Contains the main function that demonstrates the wave pattern algorithm with example cases
- **wave_pattern.go**: Implements the `waveRearrange` function and related helper functions
- **wave_pattern_test.go**: Contains unit tests to verify the correctness of the implementation


## Algorithm Explanation

The `waveRearrange` function organizes an array into a wave-like pattern by processing it in blocks of size (2x + 1). Here's how it works:

### Dividing the Array into Blocks
- The array is split into consecutive blocks, each containing (2x + 1) elements
- If the remaining elements at the end are fewer than (2x + 1), they form a smaller block

### Sorting Each Block
- Each block is sorted in ascending order
- Sorting ensures a structured arrangement where the largest element moves to the end

### Swapping the Middle and Largest Elements
- The largest element (which is at the last index after sorting) is swapped with the middle element of the block
- This creates a peak-like effect where the middle of each block has the highest value, forming a wave pattern

### Example
```go
arr := []int{10, 5, 6, 3, 2, 20, 100, 80, 90}
x := 1  // Block size = (2*1 + 1) = 3
```

#### Process:
1. Blocks: `[10, 5, 6]`, `[3, 2, 20]`, `[100, 80, 90]`
2. After sorting: `[5, 6, 10]`, `[2, 3, 20]`, `[80, 90, 100]`
3. After swapping: `[5, 10, 6]`, `[2, 20, 3]`, `[80, 90, 100]`

#### Final Output:
```go
[5, 10, 6, 2, 20, 3, 80, 100, 90] 
```

This approach ensures that within each block, the middle value forms a peak, resulting in a wave-like pattern across the array.