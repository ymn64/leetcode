package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	sign := 1

	if (dividend ^ divisor) < 0 {
		sign = -1
	}

	dividend = max(dividend, -dividend)
	divisor = max(divisor, -divisor)

	quotient := 0

	for dividend >= divisor {
		dividend -= divisor
		quotient++
	}

	quotient = sign * quotient

	if quotient < math.MinInt32 {
		return math.MinInt32
	} else if quotient > math.MaxInt32 {
		return math.MaxInt32
	}

	return quotient
}

func main() {
	fmt.Println(divide(10, 3))
}
