package main

import (
	"fmt"
	"strings"
)

func justifyLastLine(line []string, maxWidth int) string {
	n := -1

	for i := 0; i < len(line); i++ {
		n += len(line[i]) + 1
	}

	return strings.Join(line, " ") + strings.Repeat(" ", maxWidth-n)
}

func justifyLine(line []string, maxWidth int) string {
	n := len(line)

	if n == 1 {
		return line[0] + strings.Repeat(" ", maxWidth-len(line[0]))
	}

	totalText := 0

	for i := 0; i < n; i++ {
		totalText += len(line[i])
	}

	totalSpaces := maxWidth - totalText
	normalSpaces := totalSpaces / (n - 1)
	extraSpaces := totalSpaces % (n - 1)

	justified := ""

	for i := 0; i < n-1; i++ {
		justified += line[i] + strings.Repeat(" ", normalSpaces)

		if extraSpaces > 0 {
			justified += " "
			extraSpaces--
		}
	}

	justified += line[n-1]

	return justified
}

func fullJustify(words []string, maxWidth int) []string {
	justified := []string{}

	line := []string{words[0]}
	width := len(words[0])

	for i := 1; i < len(words); i++ {
		width += len(words[i]) + 1

		if width > maxWidth {
			justified = append(justified, justifyLine(line, maxWidth))
			line = []string{words[i]}
			width = len(words[i])
		} else {
			line = append(line, words[i])
		}

	}

	justified = append(justified, justifyLastLine(line, maxWidth))

	return justified
}

func debug(justified []string) {
	for _, line := range justified {
		fmt.Printf("|%s|\n", line)
	}
}

func try(words []string, maxWidth int, want []string) {
	debug(fullJustify(words, maxWidth))
	fmt.Println("===")
	debug(want)
	fmt.Println()
}

func main() {
	// words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	// maxWidth := 16
	// want := []string{
	// 	"This    is    an",
	// 	"example  of text",
	// 	"justification.  ",
	// }
	// try(words, maxWidth, want)

	words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth := 16
	want := []string{
		"What   must   be",
		"acknowledgment  ",
		"shall be        ",
	}
	try(words, maxWidth, want)

	// words = []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do", "di"}
	// maxWidth = 20
	// want = []string{
	// 	"Science  is  what we",
	// 	"understand      well",
	// 	"enough to explain to",
	// 	"a  computer.  Art is",
	// 	"everything  else  we",
	// 	"do                di",
	// }
	// try(words, maxWidth, want)
}
