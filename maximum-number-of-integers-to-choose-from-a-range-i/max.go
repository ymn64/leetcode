package main

import "fmt"

func maxCount(banned []int, n int, maxSum int) int {
	bannedSet := map[int]bool{}

	for i := 0; i < len(banned); i++ {
		if banned[i] <= n {
			bannedSet[banned[i]] = true
		}
	}

	sum := 0
	maxNums := 0

	for i := 1; i <= n; i++ {
		if !bannedSet[i] {
			sum += i
			if sum <= maxSum {
				maxNums++
			} else {
				break
			}

		}
	}

	return maxNums
}

func main() {
	fmt.Println(maxCount([]int{1, 6, 5}, 5, 6))
}
