package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	ro := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	parts := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	roman := ""

	for num > 0 {
		for _, n := range parts {
			q := num / n
			num %= n
			roman += strings.Repeat(ro[n], q)
		}
	}

	return roman
}

func main() {
	fmt.Println(intToRoman(1994))
	fmt.Println(intToRoman(9))
}
