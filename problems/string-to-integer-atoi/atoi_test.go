package main

import "testing"

var tests = []struct {
	name string
	s    string
	want int
}{
	{"test 1", "42", 42},
	{"test 2", "   -42", -42},
	{"test 3", "4193 with words", 4193},
	{"test 4", "0032", 32},
	{"test 5", "2147483648", 2147483647},
	{"test 6", "-2147483649", -2147483648},
	{"test 7", "", 0},
	{"test 8", "-", 0},
	{"test 9", "20000000000000000000", 2147483647},
	{"test 10", "9223372036854775808", 2147483647},
}

func TestMyAtoi(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := myAtoi(tt.s)
			if got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
