package main

import (
	"fmt"
	"slices"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	slices.Reverse(words)
	return strings.Join(words, " ")
}

func main() {
	fmt.Println("the sky is blue", "-", reverseWords("the sky is blue"))
}
