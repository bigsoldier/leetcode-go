package code

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var result = make([][]int, 1)
	sort.Ints(nums)
	var start = len(result)
	for i := 0; i < len(nums); i++ {
		var ret [][]int
		for j := 0; j < len(result); j++ {
			if i > 0 && nums[i] == nums[i-1] && j < start {
				continue
			}
			tmp := append(result[j], nums[i])
			ret = append(ret, tmp)
		}
		start = len(result)
		result = append(result, ret...)
	}
	return result
}
