package main

import (
	"fmt"
)

func max (a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// function that flips a matrix to get the maximum sum of the elements in the
// top left sub matrix
func flippingMatrix(matrix [][]int32) int32 {
	rows := len(matrix)
	cols := len(matrix[0])
	maxSum := int32(0)

	for rowFlip := 0; rowFlip < (1 << uint(rows)); rowFlip++ {
		for colFlip := 0; colFlip < (1 << uint(cols)); colFlip++ {
			flippedMatrix := make([][]int32, rows)
			for i := 0; i < rows; i++ {
				flippedMatrix[i] = make([]int32, cols)
				copy(flippedMatrix[i], matrix[i])
			}

			// Apply row swaps
			for i := 0; i < rows; i++ {
				if rowFlip >> uint(i)&1 == 1 {
					for j := 0; j < cols / 2; j++ {
						flippedMatrix[i][j], flippedMatrix[i][cols - j - 1] = flippedMatrix[i][cols-j-1], flippedMatrix[i][j]
					}
				}
			}

			// Apply column swaps
			for j := 0; j < cols; j++ {
				if (colFlip>>uint(j))&1 == 1 {
					for i := 0; i < rows/2; i++ {
						flippedMatrix[i][j], flippedMatrix[rows-i-1][j] = flippedMatrix[rows-i-1][j], flippedMatrix[i][j]
					}
				}
			}

			// calculate the sum of elements in the upper left corner
			currentSum := int32(0)
			for i := 0; i < rows/2; i++ {
				for j := 0; j < cols/2; j++ {
					currentSum += flippedMatrix[i][j]
				}
			}

			maxSum = max(maxSum, currentSum)
		}

	}

	return maxSum
}

func main() {
	fmt.Println(flippingMatrix([][]int32 {
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}