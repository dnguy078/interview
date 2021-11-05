package main

func IsValidSubsequence(array []int, sequence []int) bool {
	arrIndex, seqIndex := 0, 0
	for arrIndex < len(array) && seqIndex < len(sequence) {
		if array[arrIndex] == sequence[seqIndex] {
			seqIndex++
		}
		arrIndex++
	}
	return seqIndex == len(sequence)
}
