package main

import (
	"slices"
	"testing"
)

var tests = []struct {
	name string
	nums []int
	want []int
}{
	{"test 1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	{"test 2", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
}

func TestProductExceptSelf(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productExceptSelf(tt.nums)
			if !slices.Equal(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
