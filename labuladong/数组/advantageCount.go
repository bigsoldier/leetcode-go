package main

import "sort"

// 田忌赛马
func advantageCount(nums1, nums2 []int) []int {
	n := len(nums1)
	// nums2 降序
	m := make([][2]int, n)
	for i, val := range nums2 {
		m[i][0], m[i][1] = val, i
	}
	sort.Slice(m, func(i, j int) bool {
		return m[i][0] < m[j][0]
	})
	// nums1 升序
	sort.Ints(nums1)
	// 双指针
	left, right := 0, n-1
	result := make([]int, n)
	for _, v := range nums1 {
		if m[left][0] < v {
			result[m[left][1]] = v
			left++
		} else {
			result[m[right][1]] = v
			right--
		}
	}
	return result
}
