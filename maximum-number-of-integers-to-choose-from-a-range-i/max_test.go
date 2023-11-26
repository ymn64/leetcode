package main

import "testing"

var tests = []struct {
	name   string
	banned []int
	n      int
	maxSum int
	want   int
}{
	{"test 1", []int{1, 6, 5}, 5, 6, 2},
	{"test 2", []int{1, 2, 3, 4, 5, 6, 7}, 8, 1, 0},
	{"test 3", []int{11}, 7, 50, 7},
}

func Test_maxCount(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCount(tt.banned, tt.n, tt.maxSum); got != tt.want {
				t.Errorf("maxCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
