package code

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// 初始状态
	dp0 := 0
	// 第一次买入
	dp1 := -prices[0]
	// 第一次卖出
	dp2 := 0
	// 第二次买入
	dp3 := -prices[0]
	// 第二次卖出
	dp4 := 0
	for i := 1; i < len(prices); i++ {
		dp1 = max(dp1, dp0-prices[i])
		dp2 = max(dp2, dp1+prices[i])
		dp3 = max(dp3, dp2-prices[i])
		dp4 = max(dp4, dp3+prices[i])
	}
	return max(max(max(max(dp0, dp1), dp2), dp3), dp4)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
