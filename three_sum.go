package main

import (
	"sort"
)

// does not work if there are duplicates
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var results [][]int
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		start := i + 1
		end := len(nums) - 1
		for start < end {
			currentSum := nums[i] + nums[start] + nums[end]
			if currentSum == 0 {
				for start < end && nums[start] == nums[start+1] {
					start++
				}
				for start < end && nums[end] == nums[end-1] {
					end--
				}
				results = append(results, []int{nums[i], nums[start], nums[end]})
				start, end = start+1, end-1
			} else if currentSum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return results
}
