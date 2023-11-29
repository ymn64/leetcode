package main

import (
	"fmt"
	"slices"
)

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	bu := slices.Clone(nums[n-k:])

	for i := n - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}

	copy(nums, bu)
}

func try(nums []int, k int, want []int) {
	rotate(nums, k)
	fmt.Println("Got :", nums)
	fmt.Println("Want:", want)
	fmt.Println("----------------")
}

func main() {
	try([]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4})
	try([]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100})
	try([]int{-1}, 2, []int{-1})
}
