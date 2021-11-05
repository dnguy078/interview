package main

func rob(nums []int) int {
	rob1, rob2 := 0, 0

	for _, n := range nums {
		temp := max(n+rob1, rob2)
		rob1 = rob2
		rob2 = temp
	}
	return rob2
}
