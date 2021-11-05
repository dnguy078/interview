package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lastNum := 0
	for i := 1; i < len(nums); i++ {
		if nums[lastNum] != nums[i] {
			lastNum++
			nums[lastNum] = nums[i]
		}
	}
	return lastNum + 1
}
