package main

import "testing"

var tests = []struct {
	name   string
	prices []int
	want   int
}{
	{"test 1", []int{7, 1, 5, 3, 6, 4}, 7},
	{"test 2", []int{1, 2, 3, 4, 5}, 4},
	{"test 3", []int{7, 6, 4, 3, 1}, 0},
	{"test 4", []int{7}, 0},
}

func TestMaxProfit(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)

			if got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
