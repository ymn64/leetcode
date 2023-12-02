package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	reversed := 0

	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	}

	return reversed
}

func countDistinctIntegers(nums []int) int {
	set := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = struct{}{}
		set[reverse(nums[i])] = struct{}{}
	}
	return len(set)
}

func main() {
	fmt.Println(countDistinctIntegers([]int{1, 13, 10, 12, 31}))
	fmt.Println(countDistinctIntegers([]int{2, 2, 2}))
}
