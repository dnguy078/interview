package main

func search(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	} else if target > nums[mid] {
		return binarySearch(nums, target, mid+1, right)
	}
	return binarySearch(nums, target, left, mid-1)
}
