package main

import (
	"sort"
)

func NonConstructibleChange(coins []int) int {
	// Write your code here.
	sort.Ints(coins)
	currentChange := 0
	for _, c := range coins {
		if c > currentChange+1 {
			return currentChange + 1
		}
		currentChange += c
	}
	return currentChange + 1
}
