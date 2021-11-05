package main

func LongestPalindromicSubstring(str string) string {
	currentLongest := str[0:1]
	for i := 1; i < len(str); i++ {
		odd := getPandromeFrom(str, i-1, i+1)
		even := getPandromeFrom(str, i-1, i)
		longest := even
		if len(even) < len(odd) {
			longest = odd
		}
		if len(longest) > len(currentLongest) {
			currentLongest = longest
		}
	}
	return currentLongest
}

func getPandromeFrom(input string, left, right int) string {
	for left >= 0 && right < len(input) {
		if input[left] != input[right] {
			break
		}
		left--
		right++
	}

	return input[left+1 : right]
}
