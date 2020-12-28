package code

import "math"

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	k = min(k, n/2) // 做多只有k笔交易
	// 第i天进行第k笔交易，手上持有股票
	buy := make([][]int, n)
	// 第i天进行第k笔交易，手上没有股票
	sell := make([][]int, n)
	// 初始化
	for i := range buy {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}
	buy[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt64 / 2
		sell[0][i] = math.MinInt64 / 2
	}
	for i := 1; i < n; i++ {
		buy[i][0] = max(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}
	return max(sell[n-1]...)
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
