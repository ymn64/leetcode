package main

import "fmt"

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	left := 0
	right := 0

	for i := 0; i < n; i++ {
		right += nums[i]
	}

	for i := 0; i < n; i++ {
		result[i] = (2*i-n)*nums[i] - left + right

		left += nums[i]
		right -= nums[i]
	}

	return result
}

func main() {
	fmt.Println(getSumAbsoluteDifferences([]int{2, 3, 5}))
}
