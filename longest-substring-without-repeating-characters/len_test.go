package main

import "testing"

var tests = []struct {
	name string
	s    string
	want int
}{
	{"test 1", "abcabcbb", 3},
	{"test 2", "bbbbb", 1},
	{"test 3", "pwwkew", 3},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
