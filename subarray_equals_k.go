package main

func subarraySum(nums []int, k int) int {
	count, currentSum := 0, 0
	occurencemap := map[int]int{
		0: 1,
	}
	for _, num := range nums {
		currentSum += num
		if occurence, found := occurencemap[currentSum-k]; found {
			count += occurence
		}

		occurencemap[currentSum]++
	}
	return count
}
