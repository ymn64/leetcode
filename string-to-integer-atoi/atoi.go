package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")

	if len(s) == 0 {
		return 0
	}

	sign := 1
	result := 0

	switch s[0] {
	case '-':
		sign = -1
		s = s[1:]
	case '+':
		s = s[1:]
	}

	for _, c := range s {
		if c >= '0' && c <= '9' {
			d := int(c - '0')
			result = result*10 + d
			if sign*result > math.MaxInt32 {
				return math.MaxInt32
			} else if sign*result < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			break
		}
	}

	return sign * result
}

func main() {
	s := "42"
	result := myAtoi(s)
	fmt.Println(result)
}
