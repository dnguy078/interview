package main

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}
	var intersection []int
	for _, v := range nums2 {
		if count, found := m[v]; found && count > 0 {
			m[v]--
			intersection = append(intersection, v)
		}
	}
	return intersection
}
