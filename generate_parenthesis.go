package main

import (
	"strings"
)

// https://leetcode.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	// only add open parenthesis if open < n
	// only add closing parenthesis if close < n
	// valid if and only if open == close == n

	stack := []string{}
	res := []string{}

	backtrackParenthesis(&res, &stack, 0, 0, n)
	return res
}

func backtrackParenthesis(result, stack *[]string, openN, closeN, n int) {
	if (openN == closeN) && (openN == n) {
		*result = append(*result, strings.Join(*stack, ""))
		return
	}

	if openN < n {
		*stack = append(*stack, "(")
		backtrackParenthesis(result, stack, openN+1, closeN, n)
		pop := *stack
		pop = pop[:len(pop)-1]
		*stack = pop
	}

	if closeN < openN {
		*stack = append(*stack, ")")
		backtrackParenthesis(result, stack, openN, closeN+1, n)
		pop := *stack
		pop = pop[:len(pop)-1]
		*stack = pop

	}

}
