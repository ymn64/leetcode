package main

import (
	"fmt"
	"strconv"
)

func digitCount(num int) int {
	n := 0
	for num != 0 {
		num /= 10
		n++
	}
	return n
}

// digitCount(n) == 1: n
// digitCount(n) == 2: (9-1+1)*1 + (n-10+1)*2
// digitCount(n) == 3: (9-1+1)*1 + (99-10+1)*2 + (n-100+1)*3
// digitCount(n) == 4: (9-1+1)*1 + (99-10+1)*2 + (999-100+1)*3 + (n-1000+1)*4
// ...
func digitCountSum(n int) int {
	d := digitCount(n)
	sum := 0

	for i := 1; i <= d; i++ {
		p := 1
		for j := 1; j < i; j++ {
			p *= 10
		}
		sum += (min(n, 10*p-1) - p + 1) * i
	}

	return sum
}

func splitMessage(message string, limit int) []string {
	if limit <= 5 {
		return []string{}
	}

	n := 1

	for {
		textLen := n*(limit-3-digitCount(n)) - digitCountSum(n)
		// textLen := n * (limit - 3 - digitCount(n))
		// for i := 0; i < n; i++ {
		// 	textLen -= digitCount(i + 1)
		// }

		if textLen < 0 {
			return []string{}
		}

		if len(message) <= textLen {
			break
		}

		n++
	}

	parts := make([]string, n)
	start := 0
	k := limit - 3 - digitCount(n)
	a := "/" + strconv.Itoa(n) + ">"

	for i := 0; i < n-1; i++ {
		end := start + k - digitCount(i+1)
		parts[i] = message[start:end] + "<" + strconv.Itoa(i+1) + a
		start = end
	}

	parts[n-1] = message[start:] + "<" + strconv.Itoa(n) + a
	return parts
}

func main() {
	s := splitMessage("this is really a very awesome message", 9)
	for _, v := range s {
		fmt.Println(v)
	}
}
