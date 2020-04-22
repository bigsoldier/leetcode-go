#### 题目
<p>假设你正在爬楼梯。需要 <em>n</em>&nbsp;阶你才能到达楼顶。</p>

<p>每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？</p>

<p><strong>注意：</strong>给定 <em>n</em> 是一个正整数。</p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong> 2
<strong>输出：</strong> 2
<strong>解释：</strong> 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong> 3
<strong>输出：</strong> 3
<strong>解释：</strong> 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
</pre>


 #### 题解
 1、动态规划
 dp[i] = dp[i-1]+dp[i-2]
 ```go
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var climb = make([]int,n)
	climb[0] = 1
	climb[1] = 2
	for i := 2; i < n; i++ {
		climb[i] = climb[i-2] + climb[i-1]
	}
	return climb[n-1]
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-22/007001.png)
时间复杂度O(n),空间复杂度O(n)
 
 2、斐波那契数列
 只需要保存最近的两个数
 ```go
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	first,second := 1,2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-22/007002.png)
时间复杂度O(n),空间复杂度O(1)