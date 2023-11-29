package main

import "testing"

var tests = []struct {
	name string
	s    string
	want string
}{
	{"test 1", "the sky is blue", "blue is sky the"},
	{"test 2", "  hello world  ", "world hello"},
	{"test 3", "a good   example", "example good a"},
}

func TestReverseWords(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWords(tt.s)
			if got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
