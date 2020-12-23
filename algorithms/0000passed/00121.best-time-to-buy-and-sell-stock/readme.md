#### 题目
<p>给定一个数组，它的第&nbsp;<em>i</em> 个元素是一支给定股票第 <em>i</em> 天的价格。</p>

<p>如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。</p>

<p>注意你不能在买入股票前卖出股票。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [7,1,5,3,6,4]
<strong>输出:</strong> 5
<strong>解释: </strong>在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> [7,6,4,3,1]
<strong>输出:</strong> 0
<strong>解释: </strong>在这种情况下, 没有交易完成, 所以最大利润为 0。
</pre>


 #### 题解
 暴力法
 ```go
func maxProfit(prices []int) int {
	var profit int
	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			if prices[j] - prices[i] > profit {
				profit = prices[j] - prices[i]
			}
		}
	}
	return profit
}
```
 时间复杂度O(n^2^),空间复杂度O(1)
 
 一次遍历
 ```go
func maxProfit(prices []int) int {
	var minPrice, maxprofit = math.MaxInt32, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxprofit {
			maxprofit = prices[i] - minPrice
		}
	}
	return maxprofit
}

```
 时间复杂度O(n),空间复杂度O(1)