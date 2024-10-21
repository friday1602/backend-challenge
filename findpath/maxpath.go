package main

import "math"

// Function to calculate the maximum path sum
func maxPathSum(triangle [][]int) int {
	// Copy the last row to the dp array (to store max sums)
	dp := make([]int, len(triangle[len(triangle)-1]))
	copy(dp, triangle[len(triangle)-1])

	// Traverse the triangle from the second-last row upwards
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			// Choose the maximum adjacent number from the row below
			dp[j] = triangle[i][j] + int(math.Max(float64(dp[j]), float64(dp[j+1])))
		}
	}

	// The top of the triangle (dp[0]) will contain the maximum path sum
	return dp[0]
}
