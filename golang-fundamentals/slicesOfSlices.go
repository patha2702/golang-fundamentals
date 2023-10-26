package main

import (
	"fmt"
)


// Slices of slices:
// - Slices can hold other slices creating a structure like 2D matrix.

// Create a 2D matrix of tables

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++{
		for j := 0; j < cols; j++ {
			matrix[i][j] = i * j
		}
	}

	return matrix
}


// Tricky slices:
// - Always assign the result of the append() function to the same slice
// - slice = append(slice, elements)


// Range:
// -Go provides an easy way to iterate over a slice.


func main() {
	matrix := createMatrix(10,10)
	fmt.Println(matrix)

	// Range
	mySlice := []int{1,2,3,4,5}
	for index, element := range mySlice {
		fmt.Println(index, element)
	}
}
