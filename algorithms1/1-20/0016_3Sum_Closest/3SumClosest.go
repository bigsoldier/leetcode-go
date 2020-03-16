/**
 *Created by Xie Jian on 2020/2/12 15:28
 */
package _016_3Sum_Closest

import (
	"math"
	"sort"
)

//Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.
//
// Example:
//
//
//Given array nums = [-1, 2, 1, -4], and target = 1.
//
//The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
//
// Related Topics Array Two Pointers

//leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	// 排序
	sort.Ints(nums)

	var gap = math.MaxInt32
	var result int
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > target {
				if sum-target < gap {
					gap = sum - target
					result = sum
				}
				right--
			} else if sum == target {
				return target
			} else {
				if target-sum < gap {
					gap = target - sum
					result = sum
				}
				left++
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
