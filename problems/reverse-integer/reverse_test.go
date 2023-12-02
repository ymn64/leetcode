package main

import "testing"

var tests = []struct {
	name string
	x    int
	want int
}{
	{"test 1", 123, 321},
	{"test 2", -123, -321},
	{"test 3", 120, 21},
}

func TestReverse(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse(tt.x)

			if got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
