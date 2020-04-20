#### 题目
<p>一个机器人位于一个 <em>m x n </em>网格的左上角 （起始点在下图中标记为&ldquo;Start&rdquo; ）。</p>

<p>机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为&ldquo;Finish&rdquo;）。</p>

<p>问总共有多少条不同的路径？</p>

<p><img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/robot_maze.png"></p>

<p><small>例如，上图是一个7 x 3 的网格。有多少可能的路径？</small></p>

<p>&nbsp;</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> m = 3, n = 2
<strong>输出:</strong> 3
<strong>解释:</strong>
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -&gt; 向右 -&gt; 向下
2. 向右 -&gt; 向下 -&gt; 向右
3. 向下 -&gt; 向右 -&gt; 向右
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> m = 7, n = 3
<strong>输出:</strong> 28</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= m, n &lt;= 100</code></li>
	<li>题目数据保证答案小于等于 <code>2 * 10 ^ 9</code></li>
</ul>


 #### 题解
 动态规划
 我们可以发现，每个格子的路径可能性等于它的上面的格子的路径可能性加上左边的格子的路径可能性，即dp[i][j]=dp[i-1][j]+dp[i][j-1]
 ```go
func uniquePaths(m int, n int) int {
	dp := make([][]int,n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int,m)
		dp[i][0] = 1
	}
	for i := 0; i < m; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n-1][m-1]
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-20/006201.png)
时间复杂度O(n^2^),空间复杂度O(n^2^)