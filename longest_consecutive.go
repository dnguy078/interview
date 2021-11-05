package main

// https://leetcode.com/problems/longest-consecutive-sequence/
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numMap := make(map[int]int)
	for _, number := range nums {
		numMap[number]++
	}

	maxSequence := 0
	for k, _ := range numMap {
		isStart := isStartOfSequence(k, numMap)
		if isStart {
			sequenceCount := findConsecutiveCount(k, numMap)
			maxSequence = max(sequenceCount, maxSequence)
		}
	}

	return maxSequence + 1
}

func isStartOfSequence(input int, numMap map[int]int) bool {
	left := input - 1
	_, found := numMap[left]
	return !found
}

func findConsecutiveCount(input int, numMap map[int]int) int {
	consecutiveCount := 0
	for k := input + 1; k <= +input+len(numMap); k++ {
		if _, found := numMap[k]; found {
			consecutiveCount++
		} else {
			return consecutiveCount
		}
	}
	return consecutiveCount
}
