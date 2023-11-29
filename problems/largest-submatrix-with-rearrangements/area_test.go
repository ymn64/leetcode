package main

import "testing"

var tests = []struct {
	name   string
	matrix [][]int
	want   int
}{
	{"test 1", [][]int{
		{0, 0, 1},
		{1, 1, 1},
		{1, 0, 1},
	}, 4},
	{"test 2", [][]int{
		{1, 0, 1, 0, 1},
	}, 3},
	{"test 3", [][]int{
		{1, 1, 0},
		{1, 0, 1},
	}, 2},
}

func TestLargestSubmatrix(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestSubmatrix(tt.matrix)
			if got != tt.want {
				t.Errorf("largestSubmatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
