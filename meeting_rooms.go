package main

import (
	"sort"
)

// meeting room I
func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) == 0 {
		return true
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	currentInterval := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < currentInterval[1] {
			return false
		}
		currentInterval = intervals[i]
	}

	return true
}

// meeting room II
// https://leetcode.com/problems/meeting-rooms-ii/
func minMeetingRooms(intervals [][]int) int {
	startSorted, endSorted := intervals, intervals

	sort.Slice(startSorted, func(i, j int) bool {
		return startSorted[i][0] < startSorted[j][0]
	})

	sort.Slice(endSorted, func(i, j int) bool {
		return endSorted[i][1] < endSorted[j][1]
	})

	x, y := 0, 0
	count := 0
	result := 0
	for x < len(startSorted) && y < len(endSorted) {
		if startSorted[x][0] < endSorted[y][1] {
			count++
			x++
		} else {
			y++
			count--
		}
		result = max(result, count)
	}

	return result
}
