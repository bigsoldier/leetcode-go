package _033_Search_In_Rotated_Sorted_Array

//Suppose an array sorted in ascending order is rotated at some pivot unknown to
// you beforehand.
//
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
//
// You are given a target value to search. If found in the array return its inde
//x, otherwise return -1.
//
// You may assume no duplicate exists in the array.
//
// Your algorithm's runtime complexity must be in the order of O(log n).
//
// Example 1:
//
//
//Input: nums = [4,5,6,7,0,1,2], target = 0
//Output: 4
//
//
// Example 2:
//
//
//Input: nums = [4,5,6,7,0,1,2], target = 3
//Output: -1
// Related Topics Array Binary Search

//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	// 二分法
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}

		// 前半部分有序:[left:mid]递增
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target <= nums[right] && target > nums[mid] { // 在后半部分有序中
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
