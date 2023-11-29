package main

import "testing"

var tests = []struct {
	name     string
	dividend int
	divisor  int
	want     int
}{
	{"test 1", 10, 3, 3},
	{"test 2", 7, -3, -2},
	{"test 3", 1, 1, 1},
	{"test 4", -2147483648, -1, 2147483647},
	{"test 5", -2147483648, 1, -2147483648},
}

func TestDivide(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := divide(tt.dividend, tt.divisor)
			if got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
