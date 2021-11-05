package main

func lengthOfLIS(nums []int) int {
	longestAt := make([]int, len(nums))
	for i := range longestAt {
		longestAt[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if longestAt[j]+1 > longestAt[i] {
					longestAt[i] = longestAt[j] + 1
				}
			}
		}
	}

	return maxInList(longestAt...)
}

func maxInList(a ...int) int {
	m := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > m {
			m = a[i]
		}
	}
	return m
}
