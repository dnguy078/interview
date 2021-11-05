package main

// https://leetcode.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	answer := make([]int, len(nums))
	calculatePrefix(nums, answer)
	calculatePostfix(nums, answer)

	return answer
}

func calculatePrefix(input []int, output []int) {
	defaultPrefix := 1
	for i, n := range input {
		output[i] = defaultPrefix
		defaultPrefix *= n
	}
}

func calculatePostfix(input []int, output []int) {
	defaultPostfix := 1
	for i := len(input) - 1; i >= 0; i-- {
		output[i] *= defaultPostfix
		defaultPostfix *= input[i]
	}
}
