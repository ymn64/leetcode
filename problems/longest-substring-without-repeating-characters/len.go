package main

import (
	"fmt"
)

func lastIndex(s string, r rune) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == byte(r) {
			return i
		}
	}
	return -1
}

func lengthOfLongestSubstring(s string) int {
	longest := ""
	current := ""

	for _, c := range s {
		i := lastIndex(current, c)
		if i >= 0 {
			current = current[i+1:]
		}

		current = current + string(c)

		if len(current) > len(longest) {
			longest = current
		}
	}

	return len(longest)
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
