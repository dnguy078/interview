package main

// maxSubArray returns the sum of the max subarray in an array
func maxSubArray(nums []int) int {
	maxEndingHere, maxSoFar := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		maxEndingHere = max(maxEndingHere+v, v)
		maxSoFar = max(maxEndingHere, maxSoFar)
	}
	return maxSoFar
}
