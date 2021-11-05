package main

func lengthOfLongestSubstring(s string) int {
	left, right, result := 0, 0, 0
	seen := make(map[byte]bool)
	for right < len(s) {
		if !seen[s[right]] {
			result = max(result, right-left+1)
			seen[s[right]] = true
			right++
		} else {
			seen[s[left]] = false
			left++
		}
	}

	return result
}

/*
m = 1


*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
