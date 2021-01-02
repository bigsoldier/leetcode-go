package code

func maxProduct(nums []int) int {
	dpMax, dpMin, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := dpMax, dpMin
		dpMax = max(mx*nums[i], max(nums[i], nums[i]*mn))
		dpMin = min(mn*nums[i], min(nums[i], nums[i]*mx))
		ans = max(dpMax, ans)
	}
	return ans
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
