package code

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	return max(robbing(nums[1:]), robbing(nums[:n-1]))
}

func robbing(nums []int) int {
	curr, prev := 0, 0
	for _, num := range nums {
		curr, prev = max(curr, prev+num), curr
	}
	return curr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
