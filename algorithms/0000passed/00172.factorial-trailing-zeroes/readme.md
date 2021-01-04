#### 题目
<p>给定一个整数 <em>n</em>，返回 <em>n</em>! 结果尾数中零的数量。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> 3
<strong>输出:</strong> 0
<strong>解释:</strong>&nbsp;3! = 6, 尾数中没有零。</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> 5
<strong>输出:</strong> 1
<strong>解释:</strong>&nbsp;5! = 120, 尾数中有 1 个零.</pre>

<p><strong>说明: </strong>你算法的时间复杂度应为&nbsp;<em>O</em>(log&nbsp;<em>n</em>)<em>&nbsp;</em>。</p>


 #### 题解
 1、计算阶乘
 速度太慢，同时可能会有溢出
 
 2、计算5的个数
 观察可知，5的倍数能生成1个0,25的倍数能生成2个0,125的倍数能生成3个0
 ```go
func trailingZeroes(n int) int {
	var zeroCount int
	for i := 5; i <= n; i += 5 {
		powerOf5 := 5
		for i % powerOf5 == 0 {
			zeroCount += 1
			powerOf5 *= 5
		}
	}
	return zeroCount
}
```

 3、更高效的方法
 ```go
func trailingZeroes(n int) int {
	var zeroCount int
	powerOf5 := 5
	for n >= powerOf5 {
		zeroCount += n/powerOf5
		powerOf5 *= 5
	}
	return zeroCount
}
```
```go
func trailingZeroes(n int) int {
	var zeroCount int
	for n > 0 {
		n /= 5
		zeroCount += n
	}
	return zeroCount
}
```
 时间复杂度O(logn),空间复杂度O(1)