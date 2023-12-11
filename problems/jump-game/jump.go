package main

import "fmt"

// Inspiration: https://github.com/devkvlt/aoc/blob/main/2022/day15/main.go#L49
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	type Interval struct{ start, end int }

	merged := []Interval{{0, nums[0]}}

	for i := 1; i < len(nums); i++ {
		prev := merged[len(merged)-1]
		curr := Interval{i, i + nums[i]}

		if curr.start-1 >= prev.end {
			merged = append(merged, curr)
		} else {
			merged[len(merged)-1].end = max(curr.end, prev.end)
		}

	}

	return merged[0].end >= len(nums)-1
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}
