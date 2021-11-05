package main

import "math"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	m := make(map[byte]int)
	for _, v := range []byte(t) {
		m[v]++
	}
	start, end, index := 0, 0, 0
	counter := len(t)
	minLen := math.MaxInt32
	for end < len(s) {
		if m[s[end]] > 0 {
			counter--
		}
		m[s[end]]--
		end++
		for counter == 0 {
			if minLen > end-start {
				minLen = end - start
				index = start
			}
			m[s[start]]++
			if m[s[start]] == 1 {
				counter++
			}
			start++
		}
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[index : index+minLen]
}
