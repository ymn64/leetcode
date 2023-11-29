package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)

	if n <= 2 {
		return n
	}

	k := 2

	for i := 2; i < n; i++ {
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func removeDuplicates2(nums []int) int {
	n := len(nums)
	k := n
	mustRemove := false

	for i := 1; i < n; i++ {
		if nums[i-1-n+k] == nums[i] {
			if mustRemove {
				k--
			} else {
				mustRemove = true
				nums[i-n+k] = nums[i]
			}
		} else {
			nums[i-n+k] = nums[i]
			mustRemove = false
		}
	}

	return k
}

func main() {
	s := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(s)
	k := removeDuplicates(s)
	fmt.Println(s[:k])
	fmt.Println([]int{0, 0, 1, 1, 2, 3, 3})
}
