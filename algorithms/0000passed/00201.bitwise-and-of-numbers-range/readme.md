#### 题目
<p>给定范围 [m, n]，其中 0 &lt;= m &lt;= n &lt;= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。</p>

<p><strong>示例 1:&nbsp;</strong></p>

<pre><strong>输入:</strong> [5,7]
<strong>输出:</strong> 4</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> [0,1]
<strong>输出:</strong> 0</pre>


 #### 题解
 1、暴力法
 ```go
func rangeBitwiseAnd(m int, n int) (ans int) {
	ans = m
	for i := m+1; i <= n; i++ {
		ans = ans & i
	}
	return
}
```
 时间复杂度O(n),空间复杂度O(1)
 
 本质上这题求的是公共前缀，因为 `1101 & 1110 = 1100`
 
 2、位移
 ```go
func rangeBitwiseAnd(m int, n int) int {
	var ans int
	for m < n {
		m,n = m >> 1,n>>1
		ans++
	}
	return m<<ans
}
```
 时间复杂度O(logn),空间复杂度O(1)
 
 3、Brian Kerngighan算法
 n和n-1做按位与运算后，n最右边的1会变成0
 ```go
func rangeBitwiseAnd(m int, n int) int {
	for m < n {
		n &= n-1
	}
	return n
}
```
 时间复杂度O(logn),空间复杂度O(1)