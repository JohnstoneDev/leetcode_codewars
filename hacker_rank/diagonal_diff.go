package main

import (
	"fmt"
	"math"
)

// Function that calculates the diagonal difference of a square
// matrix
func diagonalDifference(arr [][]int32) int32 {
	length := len(arr)
	mainDiagonal := make([]int32, length)
	secondaryDiagonal := make([]int32, length)

	// Load up the diagonal values
	for i := 0; i < length; i++ {
		mainDiagonal[i] = arr[i][i]
		secondaryDiagonal[i] = arr[i][length-i-1]
	}

	// Calculate the sums of the diagonals
	var sumMain, sumSecondary int32

	for _, val := range mainDiagonal {
		sumMain += val
	}

	for _, val := range secondaryDiagonal {
		sumSecondary += val
	}

	// Calculate the Absolute Difference
	absoluteDifference := math.Abs(float64(sumMain) - float64(sumSecondary))

	return int32(absoluteDifference)
}

// Entry Point
func main(){
	fmt.Println(diagonalDifference([][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}), "<= expecting 15")
}