package main

import (
	"slices"
	"testing"
)

var tests = []struct {
	name     string
	words    []string
	maxWidth int
	want     []string
}{
	{
		"test 1",
		[]string{"This", "is", "an", "example", "of", "text", "justification."},
		16,
		[]string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		},
	},
	{
		"test 2",
		[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
		16,
		[]string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		},
	},
	{
		"test 3",
		[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
		20,
		[]string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  ",
		},
	},
}

func TestFullJustify(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fullJustify(tt.words, tt.maxWidth)
			if !slices.Equal(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
