package main

func SortedSquaredArray(array []int) []int {
	sortedSquares := make([]int, len(array))
	left, right := 0, len(array)-1
	for pos := len(array) -1; pos >= 0; pos-- {
		if absoluteValue(array[right]) > absoluteValue(array[left]) {
			sortedSquares[pos] = array[right] * array[right]
			right--
		} else {
			sortedSquares[pos] = array[left] * array[left]
			left++
		}
	}
	return sortedSquares
}

func absoluteValue(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
