package main

import (
	"strings"
)

// O(n)
func IsPalindrome(str string) bool {
	j := len(str) - 1
	i := 0

	for i < j {
		for i < j && !isAlphanumeric(str[i]) {
			i++
		}
		for i < j && !isAlphanumeric(str[j]) {
			j--
		}

		if !strings.EqualFold(string(str[i]), string(str[j])) {
			return false
		}

		j--
		i++
	}

	return true
}

func isAlphanumeric(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}
