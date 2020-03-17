package code

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n { // 保证 nums1 最短
		return findMedianSortedArrays(nums2, nums1)
	}

	var nums1Left, nums1Right, nums2Left, nums2Right int
	var right = m * 2
	var left = 0
	for left <= right {
		c1 := (left + right) / 2
		c2 := m + n - c1

		if c1 != 0 {
			nums1Left = nums1[(c1-1)/2]
		} else {
			nums1Left = math.MinInt64
		}
		if c1 != 2*m {
			nums1Right = nums1[c1/2]
		} else {
			nums1Right = math.MaxInt64
		}
		if c2 != 0 {
			nums2Left = nums2[(c2-1)/2]
		} else {
			nums2Left = math.MinInt64
		}
		if c2 != 2*n {
			nums2Right = nums2[c2/2]
		} else {
			nums2Right = math.MaxInt64
		}

		if nums1Left > nums2Right {
			right = c1 - 1
		} else if nums2Left > nums1Right {
			left = c1 + 1
		} else {
			break
		}
	}

	return (math.Max(float64(nums1Left), float64(nums2Left)) + math.Min(float64(nums1Right), float64(nums2Right))) / 2.0
}
