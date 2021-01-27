package code

func lengthOfLIS(nums []int) (ans int) {
	n := len(nums)
	if n == 0 {
		return 0
	}
	var dp = make([]int, n)
	dp[0] = 1
	ans = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
