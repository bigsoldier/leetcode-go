#### 题目
<p>当 <code>A</code>&nbsp;的子数组&nbsp;<code>A[i], A[i+1], ..., A[j]</code>&nbsp;满足下列条件时，我们称其为<em>湍流子数组</em>：</p>

<ul>
	<li>若&nbsp;<code>i &lt;= k &lt; j</code>，当 <code>k</code>&nbsp;为奇数时，&nbsp;<code>A[k] &gt; A[k+1]</code>，且当 <code>k</code> 为偶数时，<code>A[k] &lt; A[k+1]</code>；</li>
	<li><strong>或 </strong>若&nbsp;<code>i &lt;= k &lt; j</code>，当 <code>k</code> 为偶数时，<code>A[k] &gt; A[k+1]</code>&nbsp;，且当 <code>k</code>&nbsp;为奇数时，&nbsp;<code>A[k] &lt; A[k+1]</code>。</li>
</ul>

<p>也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。</p>

<p>返回 <code>A</code> 的最大湍流子数组的<strong>长度</strong>。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>[9,4,2,10,7,8,8,1,9]
<strong>输出：</strong>5
<strong>解释：</strong>(A[1] &gt; A[2] &lt; A[3] &gt; A[4] &lt; A[5])
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>[4,8,12,16]
<strong>输出：</strong>2
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>[100]
<strong>输出：</strong>1
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ol>
	<li><code>1 &lt;= A.length &lt;= 40000</code></li>
	<li><code>0 &lt;= A[i] &lt;= 10^9</code></li>
</ol>


 #### 题解
 1、双指针法
 ```go
func maxTurbulenceSize(arr []int) (ans int) {
	n := len(arr)
	ans = 1
	var left,right int
	for right < n-1 {
		if left == right {
			if arr[left] == arr[left+1] {
				left++
			}
			right++
		} else {
			if arr[right-1] < arr[right] && arr[right] > arr[right+1] {
				right++
			} else if arr[right-1] > arr[right] && arr[right] < arr[right+1] {
				right++
			} else {
				left = right
			}
		}
		ans = max(ans,right-left+1)
	}
	return
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
```
 时间复杂度O(n),空间复杂度O(1)
 
 2、动态规划
 我们能够得到一个动态规划方程
 ```go
func maxTurbulenceSize(arr []int) (ans int) {
	n := len(arr)
	// dp[i][0]表示以arr[i]结尾，且arr[i-1]<arr[i]的湍流子数组的最大长度
	// dp[i][1]表示以arr[i]结尾，且arr[i-1]>arr[i]的湍流子数组的最大长度
	dp := make([][2]int,n)
	// 边界
	dp[0] = [2]int{1,1}
	for i := 1; i < n; i++ {
		dp[i] = [2]int{1,1}
		if arr[i] > arr[i-1] {
			dp[i][0] = dp[i-1][1]+1
		} else if arr[i] < arr[i-1] {
			dp[i][1] = dp[i-1][0]+1
		}
	}
	ans = 1
	for i := 0; i < n; i++ {
		ans = max(ans,dp[i][0])
		ans = max(ans,dp[i][1])
	}
	return
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 3、优化的动态规划
 ```go
func maxTurbulenceSize(arr []int) (ans int) {
	n := len(arr)
	ans = 1
	dp0,dp1 := 1,1
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			dp0,dp1 = dp1+1,1
		} else if arr[i] < arr[i-1] {
			dp0,dp1 = 1,dp0+1
		} else {
			dp0,dp1 = 1,1
		}
		ans = max(ans,max(dp0,dp1))
	}
	return
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
```
 时间复杂度O(n),空间复杂度O(1)