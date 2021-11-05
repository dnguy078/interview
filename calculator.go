package main

import (
	"strconv"
)

// https://leetcode.com/problems/basic-calculator/
// uses stack
func calculate(s string) int {
	sign, sum := 1, 0
	var stack []int
	for i := 0; i < len(s); i++ {
		// read the number and add to current sum
		if isNumber(s[i]) {
			number := 0
			for i < len(s) && isNumber(s[i]) {
				val, _ := strconv.Atoi(string(s[i]))
				number = (number * 10) + val
				i++
			}
			sum += number * sign
			i--
		} else if s[i] == '+' {
			sign = 1
		} else if s[i] == '-' {
			sign = -1
		} else if s[i] == '(' {
			// push the sum and sign to stack
			// reset sum and sign
			stack = append(stack, sum)
			stack = append(stack, sign)
			sum = 0
			sign = 1
		} else if s[i] == ')' {
			sum = stack[len(stack)-1] * sum
			// pop the stack
			stack = stack[:len(stack)-1]

			sum += stack[len(stack)-1]
			// pop the stack
			stack = stack[:len(stack)-1]
		}
	}
	return sum
}

func isNumber(input byte) bool {
	return input >= '0' && input <= '9'
}
