package code

import "math"

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	var sum int
	ans := math.MaxInt32
	left, right := 0, 0
	for right < n {
		sum += nums[right]
		for sum >= s {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
