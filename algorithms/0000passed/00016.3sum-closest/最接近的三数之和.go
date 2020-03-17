package code

import (
	"math"
	"sort"
)

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
