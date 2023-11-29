package main

import (
	"slices"
	"testing"
)

var tests = []struct {
	name     string
	nums     []int
	wantK    int
	wantNums []int
}{
	{"test 1", []int{1, 1, 1, 2, 2, 3}, 5, []int{1, 1, 2, 2, 3}},
	{"test 2", []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7, []int{0, 0, 1, 1, 2, 3, 3}},
	{"test 3", []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
	{"test 4", []int{}, 0, []int{}},
	{"test 5", []int{1}, 1, []int{1}},
	{"test 6", []int{1}, 1, []int{1}},
	{"test 7", []int{1, 1, 1, 1}, 2, []int{1, 1}},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)

			if got != tt.wantK {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.wantK)
			}

			if !slices.Equal(tt.nums[:got], tt.wantNums) {
				t.Errorf("removeDuplicates() -> %v, want %v", tt.nums[:got], tt.wantNums)
			}
		})
	}
}
