package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

type State struct {
	current string
	open    int
	closed  int
}

func generateParenthesis(n int) []string {
	parens := []string{}
	queue := []State{{"(", 1, 0}}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		current := state.current
		open := state.open
		closed := state.closed

		if open < n {
			queue = append(queue, State{current + "(", open + 1, closed})
		}

		if closed < open {
			queue = append(queue, State{current + ")", open, closed + 1})
		}

		if open == closed && open == n {
			parens = append(parens, current)
		}
	}

	return parens
}
