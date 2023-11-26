package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prods := make([]int, n)

	left := 1

	for i := 0; i < n; i++ {
		prods[i] = left
		left *= nums[i]
	}

	right := 1

	for i := n - 1; i >= 0; i-- {
		prods[i] *= right
		right *= nums[i]
	}

	return prods
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println([]int{24, 12, 8, 6})
}
