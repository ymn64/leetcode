package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	reversed := 0

	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	}

	return reversed
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(10))
}
