package main

func moveZeroes(nums []int) {
	for lastZeroIdx, current := 0, 0; current < len(nums); current++ {
		if nums[current] != 0 {
			nums[current], nums[lastZeroIdx] = nums[lastZeroIdx], nums[current]
			lastZeroIdx++
		}
	}
}
