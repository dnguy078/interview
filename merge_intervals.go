package main

import "sort"

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	return sortInterval(intervals)
}

func sortInterval(interval [][]int) [][]int {
	// sort the interval
	sort.Slice(interval, func(i, j int) bool {
		return interval[i][0] < interval[j][0]
	})

	results := make([][]int, 0)
	// add the first to result
	currentInterval := interval[0]
	results = append(results, currentInterval)

	for _, nextInterval := range interval {
		currentIntervalEnd := currentInterval[1]
		nextIntervalStart, nextIntervalEnd := nextInterval[0], nextInterval[1]
		if currentIntervalEnd >= nextIntervalStart {
			// take the max of the two for the end
			currentInterval[1] = max(currentIntervalEnd, nextIntervalEnd)
		} else {
			currentInterval = nextInterval
			results = append(results, currentInterval)
		}
	}

	return results
}
