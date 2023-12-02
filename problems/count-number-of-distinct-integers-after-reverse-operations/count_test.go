package main

import "testing"

var tests = []struct {
	name string
	nums []int
	want int
}{
	{"test 1", []int{1, 13, 10, 12, 31}, 6},
	{"test 2", []int{2, 2, 2}, 1},
}

func TestCountDistinctIntegers(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countDistinctIntegers(tt.nums)

			if got != tt.want {
				t.Errorf("countDistinctIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
