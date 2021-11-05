package main

// https://leetcode.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	frequency := make([][]int, len(nums)+1)

	for _, n := range nums {
		count[n] = 1 + count[n]
	}
	for k, v := range count {
		frequency[v] = append(frequency[v], k)
	}
	var results []int
	for i := len(frequency) - 1; i >= 0; i-- {
		if len(frequency[i]) != 0 {
			results = append(results, frequency[i]...)
			if len(results) == k {
				return results
			}
		}
	}

	return results
}
