#### 题目
<p>给定一个数组，它的第<em> i</em> 个元素是一支给定的股票在第 <em>i </em>天的价格。</p>

<p>设计一个算法来计算你所能获取的最大利润。你最多可以完成 <strong>k</strong> 笔交易。</p>

<p><strong>注意:</strong>&nbsp;你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> [2,4,1], k = 2
<strong>输出:</strong> 2
<strong>解释:</strong> 在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> [3,2,6,5,0,3], k = 2
<strong>输出:</strong> 7
<strong>解释:</strong> 在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
&nbsp;    随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。
</pre>


 #### 题解
 动态规划
 
 用buy[i][j]表示恰好进行j笔交易，并且当前手上持有一只股票，这种情况下的最大利润;
 用sell[i][j]表示恰好进行j笔交易，并且手上没有股票，这种情况下的最大利润。
 
 我们可以推导状态转移方程。
 
 对于buy[i][j]，如果是第i天买的，那么在第i-1天时，我们手上不持有股票，对应sell[i-1][j],此时需要扣除买入prices[0]的费用;
 如果不是i天买入，那么第i-1天，我们手上持有股票，
 buy[i][j]=max{buy[i-1][j],sell[i-1][j]-prices[i]}
 
 对于sell[i][j],如果是第i天卖的，那么在第i-1天时，我们卖出股票，对应buy[i-1][j-1]+prices[i];如果不是第i天卖的，那么手上不持有股票。
 sell[i][j]=max{sell[i-1][j],buy[i-1][j-1]+prices[i]}
 
 确定边界条件
 
 buy[0][0]=-prices[0]
 
 buy[0][j],sell[0][j]设置最小值
 ```go
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	k = min(k,n/2) // 做多只有k笔交易
	// 第i天进行第k笔交易，手上持有股票
	buy := make([][]int,n)
	// 第i天进行第k笔交易，手上没有股票
	sell := make([][]int,n)
	// 初始化
	for i := range buy {
		buy[i] = make([]int,k+1)
		sell[i] = make([]int,k+1)
	}
	buy[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt64/2
		sell[0][i] = math.MinInt64/2
	}
	for i := 1; i < n; i++ {
		buy[i][0] = max(buy[i-1][0],sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = max(buy[i-1][j],sell[i-1][j]-prices[i])
			sell[i][j] = max(sell[i-1][j],buy[i-1][j-1]+prices[i])
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
```
 时间复杂度O(n*min(n,k)),空间复杂度O(n*min(n,k))
 
 简化
 ```go
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	k = min(k,n/2)
	buy,sell := make([]int,k+1),make([]int,k+1)
	buy[0] = -prices[0]

	for i := 1; i <= k; i++ {
		buy[i] = math.MinInt64/2
		sell[i] = math.MinInt64/2
	}
	for i := 1; i < n; i++ {
		buy[0] = max(buy[0],sell[0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[j] = max(buy[j],sell[j]-prices[i])
			sell[j] = max(sell[j],buy[j-1]+prices[i])
		}
	}
	return max(sell...)
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
```
 时间复杂度O(n*min(n,k)),空间复杂度O(n*min(n,k))