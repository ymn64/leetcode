package main

import (
	"slices"
	"testing"
)

var tests = []struct {
	name string
	nums []int
	k    int
	want []int
}{
	{"test 1", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
	{"test 2", []int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	{"test 3", []int{-1}, 2, []int{-1}},
}

func TestRotate(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.nums, tt.k)
			if !slices.Equal(tt.nums, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
