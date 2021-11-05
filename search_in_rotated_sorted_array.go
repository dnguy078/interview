package main

func searchInRotated(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if target == nums[mid] {
			return mid
		}
		// left portion
		if nums[l] <= nums[mid] {
			if target > nums[mid] || target < nums[l] {
				// check the right portion
				l = mid + 1
			} else {
				// its inside the left
				r = mid - 1
			}
		} else {
			// this handles the right portion
			if target < nums[mid] || target > nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
