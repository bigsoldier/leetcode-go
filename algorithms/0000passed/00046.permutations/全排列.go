package code

func permute(nums []int) [][]int {
	var output [][]int
	n := len(nums)
	backtrack(n, 0, nums, &output)
	return output
}

func backtrack(n, index int, nums []int, output *[][]int) {
	if n == index {
		var tmp = make([]int, len(nums))
		copy(tmp, nums)
		*output = append(*output, tmp)
	}
	for i := index; i < n; i++ {
		nums[i], nums[index] = nums[index], nums[i]
		backtrack(n, index+1, nums, output)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
