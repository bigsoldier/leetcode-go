package code

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 { // 如果数组小于3，返回[]
		return nil
	}

	// 对数组排序(快排，时间复杂度O(nlogn))
	sort.Ints(nums)

	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 { // 如果当前元素大于0，后续肯定不会存在符合的三元组
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // 对于出现的重复解，跳过
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] { // 去除左指针的重复元素
					left++
				}
				for left < right && nums[right] == nums[right-1] { // 去除右指针的重复元素
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] > 0 { // 三数之和大于0，移动右指针
				right--
			} else { // 三数之和大于0，移动左指针
				left++
			}
		}
	}
	return result
}
