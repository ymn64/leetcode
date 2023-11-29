package main

import (
	"testing"
)

func BenchmarkRemoveDuplicates(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4}

	for i := 0; i < b.N; i++ {
		removeDuplicates(nums)
	}
}

func BenchmarkRemoveDuplicates2(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4}

	for i := 0; i < b.N; i++ {
		removeDuplicates2(nums)
	}
}
