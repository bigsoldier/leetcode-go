package code

import "sort"

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		} else if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		} else if nums[i]+nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3] < target {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			// 去重
			if j-i > 1 && nums[j] == nums[j-1] {
				continue
			} else if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			} else if nums[i]+nums[j]+nums[len(nums)-1]+nums[len(nums)-2] < target {
				continue
			}

			left, right := j+1, len(nums)-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return result
}
