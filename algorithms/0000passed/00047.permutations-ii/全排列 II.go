package code

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var output [][]int
	used := make([]bool, len(nums))
	var p []int
	backtrack(nums, 0, p, &output, &used)
	return output
}

func backtrack(nums []int, index int, p []int, output *[][]int, used *[]bool) {
	if len(nums) == index {
		var tmp = make([]int, len(nums))
		copy(tmp, p)
		*output = append(*output, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] {
				continue
			}
			(*used)[i] = true
			backtrack(nums, index+1, append(p, nums[i]), output, used)
			(*used)[i] = false
		}
	}
}
