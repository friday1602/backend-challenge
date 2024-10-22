package main

import (
	"math"
)

// Function to calculate the maximum path sum
func maxPathSum(triangle [][]int) int {
	// Copy the last row to the row's slice
	row := make([]int, len(triangle[len(triangle)-1]))
	copy(row, triangle[len(triangle)-1])

	// Traverse the triangle from the second-last row upwards
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			// bring two numbers below sum with number on current row then reassgin to the array at that index
			row[j] = triangle[i][j] + int(math.Max(float64(row[j]), float64(row[j+1])))
		}
	}
	// return row at index 0 (top)
	return row[0]
}
