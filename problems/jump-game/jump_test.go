package main

import "testing"

var tests = []struct {
	name string
	nums []int
	want bool
}{
	{"test 1", []int{2, 3, 1, 1, 4}, true},
	{"test 2", []int{3, 2, 1, 0, 4}, false},
	{"test 3", []int{2, 5, 0, 0}, true},
	{"test 4", []int{1, 2, 3}, true},
	{"test 5", []int{2, 0, 0}, true},
	{"test 6", []int{0}, true},
}

func TestCanJump(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canJump(tt.nums)
			if got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
