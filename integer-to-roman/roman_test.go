package main

import "testing"

var tests = []struct {
	name string
	num  int
	want string
}{
	{"test 1", 3, "III"},
	{"test 2", 4, "IV"},
	{"test 3", 9, "IX"},
	{"test 4", 58, "LVIII"},
	{"test 5", 1994, "MCMXCIV"},
	{"test 6", 1995, "MCMXCV"},
}

func TestIntToRoman(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intToRoman(tt.num)
			if got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
