#### 题目
<p>给定一个数组，它的第<em> i</em> 个元素是一支给定的股票在第 <em>i </em>天的价格。</p>

<p>设计一个算法来计算你所能获取的最大利润。你最多可以完成&nbsp;<em>两笔&nbsp;</em>交易。</p>

<p><strong>注意:</strong>&nbsp;你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> [3,3,5,0,0,3,1,4]
<strong>输出:</strong> 6
<strong>解释:</strong> 在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
&nbsp;    随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> [1,2,3,4,5]
<strong>输出:</strong> 4
<strong>解释:</strong> 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。 &nbsp; 
&nbsp;    注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。 &nbsp; 
&nbsp;    因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入:</strong> [7,6,4,3,1] 
<strong>输出:</strong> 0 
<strong>解释:</strong> 在这个情况下, 没有交易完成, 所以最大利润为 0。</pre>


 #### 题解
 暴力法
 
 有0<=l1<r1<=l2<r2<n来获取两次的最大利润
 ```go
func maxProfit(prices []int) int {
	var profit int
	for l1 := 0; l1 < len(prices); l1++ {
		for r1 := l1+1; r1 < len(prices); r1++ {
			profit = max(profit,prices[r1]-prices[l1])
			for l2 := r1; l2 < len(prices); l2++ {
				for r2 := l2+1; r2 < len(prices); r2++ {
					profit = max(profit,prices[r1]-prices[l1]+prices[r2]-prices[l2])
				}
			}
		}
	}
	return profit
}
```
 时间复杂度O(n^4^),空间复杂度O(1)
 
 同样的思路，分成两个区间
 ```go
func maxProfit(prices []int) int {
	var profit int
	var min1,max1 = math.MaxInt32,0
	for i := 0; i < len(prices); i++ {
		// 求[0,i]的最大差值
		min1 = min(min1,prices[i])
		max1 = max(max1,prices[i]-min1)
		var min2,max2 = math.MaxInt32,0
		for j := i; j < len(prices); j++ {
			min2 = min(min2,prices[j])
			max2 = max(max2,prices[j]-min2)
		}
		profit = max(profit,max1+max2)
	}
	return profit
}

```
 时间复杂度O(n^2^),空间复杂度O(1)
 
 两遍遍历
 
 正向遍历获取最低点和当前的差值，反向遍历获取最大值与当前的差值，本质上还是在找中间的那次操作。
 ```go
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// 正向遍历
	var profit = make([]int,len(prices))
	profit[0] = 0
	var minVal = prices[0]
	for i := 1; i < len(prices); i++ {
		if minVal > prices[i] {
			minVal = prices[i]
		}
		profit[i] = prices[i] - minVal
	}
	// 反向遍历
	var ans = profit[len(prices)-1] // 总利润
	var value = 0
	var maxVal = 0
	for i := len(prices)-1; i > 0; i-- {
		if maxVal < prices[i] {
			maxVal = prices[i]
		}
		value = max(value,maxVal-prices[i])
		ans = max(ans,value+profit[i-1])
	}
	return ans
}
```
 时间复杂度O(n),空间复杂度O(1)
 
 动态规划
 ```go
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var dp = make([][5]int,len(prices))
	// 初始状态
	dp[0][0] = 0
	// 第一次买入
	dp[0][1] = -prices[0]
	// 第一次卖出
	dp[0][2] = 0
	// 第二次买入
	dp[0][3] = -prices[0]
	// 第二次卖出
	dp[0][4] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1],dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2],dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3],dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4],dp[i-1][3]+prices[i])
	}
	return max(max(max(max(dp[len(prices)-1][0],dp[len(prices)-1][1]),dp[len(prices)-1][2]),dp[len(prices)-1][3]),dp[len(prices)-1][4])
}
```
 时间复杂度O(n),空间复杂度O(1)
 
 优化后的动态规划
 ```go
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
		dp1 = max(dp1,dp0-prices[i])
		dp2 = max(dp2,dp1+prices[i])
		dp3 = max(dp3,dp2-prices[i])
		dp4 = max(dp4,dp3+prices[i])
	}
	return max(max(max(max(dp0,dp1),dp2),dp3),dp4)
}
```