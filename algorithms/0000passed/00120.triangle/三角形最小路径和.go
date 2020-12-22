package code

func minimumTotal(triangle [][]int) int {
	var dp = make([]int, len(triangle))
	dp[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		dp[len(triangle[i])-1] = dp[len(triangle[i])-2] + triangle[i][len(triangle[i])-1]

		for j := len(triangle[i]) - 2; j >= 1; j-- {
			dp[j] = min(dp[j], dp[j-1]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}
	var minVal = dp[0]
	for _, val := range dp {
		minVal = min(minVal, val)
	}
	return minVal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
