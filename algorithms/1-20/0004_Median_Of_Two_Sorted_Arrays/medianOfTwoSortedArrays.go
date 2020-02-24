package _004_Median_Of_Two_Sorted_Arrays

//There are two sorted arrays nums1 and nums2 of size m and n respectively.
//
// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
//
// You may assume nums1 and nums2 cannot be both empty.
//
// Example 1:
//
//
//nums1 = [1, 3]
//nums2 = [2]
//
//The median is 2.0
//
//
// Example 2:
//
//
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//The median is (2 + 3)/2 = 2.5
//
// Related Topics Array Binary Search Divide and Conquer

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	sumLen := len1 + len2
	i, j := 0, 0
	var resNum = make([]int, len1+len2)
	for k := 0; k < sumLen; k++ {
		if i == len1 || (i < len1 && j < len2 && nums1[i] > nums2[j]) {
			resNum[k] = nums2[j]
			j++
			continue
		}
		if j == len2 || (i < len1 && j < len2 && nums1[i] <= nums2[j]) {
			resNum[k] = nums1[i]
			i++
		}
	}

	if sumLen%2 == 0 {
		return float64(resNum[sumLen/2]+resNum[sumLen/2-1]) / 2.0
	} else {
		return float64(resNum[sumLen/2])
	}
}

// 其实意思是找出两个数组合并后的有序数组 中位数
// 合并数组的时间复杂度为O(max(m,n))
// 所以问题关键在于如何高效合并两个有序数组
