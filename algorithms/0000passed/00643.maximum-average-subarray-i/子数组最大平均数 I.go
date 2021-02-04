package code

func findMaxAverage(nums []int, k int) (ans float64) {
	var sum int
	for _, num := range nums[:k] {
		sum += num
	}
	ans = float64(sum) / float64(k)
	for i := k; i < len(nums); i++ {
		sum += nums[i]
		sum -= nums[i-k]
		ans = max(ans, float64(sum)/float64(k))
	}
	return
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
