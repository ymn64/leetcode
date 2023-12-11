package main

import (
	"testing"
)

var tests = []struct {
	name string
	n    int
	want []string
}{
	{"test1", 1, []string{
		"()",
	}},
	{"test2", 2, []string{
		"(())",
		"()()",
	}},
	{"test3", 3, []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}},
	{"test4", 4, []string{
		"(((())))",
		"((()()))",
		"((())())",
		"((()))()",
		"(()(()))",
		"(()()())",
		"(()())()",
		"(())(())",
		"(())()()",
		"()((()))",
		"()(()())",
		"()(())()",
		"()()(())",
		"()()()()",
	}},
}

func equal(s1, s2 []string) bool {
	type Set map[string]bool
	a := make(Set)
	for _, v := range s1 {
		a[v] = true
	}
	b := make(Set)
	for _, v := range s2 {
		b[v] = true
	}
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if !b[k] {
			return false
		}
	}
	return true
}

func TestGenerateParenthesis(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.n)
			if !equal(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
