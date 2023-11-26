// NOTE: This problem is a waste of time!

package main

import (
	"fmt"
	"slices"
)

func largestSubmatrix(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	for j := 0; j < cols; j++ {
		for i := 1; i < rows; i++ {
			if matrix[i][j] == 1 {
				matrix[i][j] += matrix[i-1][j]
			}
		}
	}

	for i := 0; i < rows; i++ {
		slices.Sort(matrix[i])
	}

	area := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if a := matrix[i][j] * (cols - j); a > area {
				area = a
			}
		}
	}

	return area
}

func main() {
	m := [][]int{
		{0, 0, 1},
		{1, 1, 1},
		{1, 0, 1},
	}
	fmt.Println(largestSubmatrix(m))
}
