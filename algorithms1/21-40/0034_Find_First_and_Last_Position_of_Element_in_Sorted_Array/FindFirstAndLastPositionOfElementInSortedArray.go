package _034_Find_First_and_Last_Position_of_Element_in_Sorted_Array

//Given an array of integers nums sorted in ascending order, find the starting a
//nd ending position of a given target value.
//
// Your algorithm's runtime complexity must be in the order of O(log n).
//
// If the target is not found in the array, return [-1, -1].
//
// Example 1:
//
//
//Input: nums = [5,7,7,8,8,10], target = 8
//Output: [3,4]
//
// Example 2:
//
//
//Input: nums = [5,7,7,8,8,10], target = 6
//Output: [-1,-1]
// Related Topics Array Binary Search

//leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	var index = []int{-1, -1}
	var leftIndex = insertIndex(nums, target, true)
	if leftIndex == len(nums) || nums[leftIndex] != target {
		return index
	}
	index[0] = leftIndex
	index[1] = insertIndex(nums, target, false) - 1
	return index
}

// 遍历列表，haw为true时，找左边界，haw为false时，找右边界
func insertIndex(nums []int, target int, haw bool) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if target < nums[mid] || (haw && target == nums[mid]) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

//leetcode submit region end(Prohibit modification and deletion)
