package main

import "testing"

var tests = []struct {
	name string
	num  int
	want string
}{
	{"test 1", 123, "One Hundred Twenty Three"},
	{"test 2", 12345, "Twelve Thousand Three Hundred Forty Five"},
	{"test 3", 1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
	{"test 4", 0, "Zero"},
}

func TestNumberToWords(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberToWords(tt.num)
			if got != tt.want {
				t.Errorf("numberToWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
