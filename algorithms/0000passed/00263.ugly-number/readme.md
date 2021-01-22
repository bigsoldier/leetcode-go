#### 题目
<p>编写一个程序判断给定的数是否为丑数。</p>

<p>丑数就是只包含质因数&nbsp;<code>2, 3, 5</code>&nbsp;的<strong>正整数</strong>。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> 6
<strong>输出:</strong> true
<strong>解释: </strong>6 = 2 &times;&nbsp;3</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> 8
<strong>输出:</strong> true
<strong>解释: </strong>8 = 2 &times; 2 &times;&nbsp;2
</pre>

<p><strong>示例&nbsp;3:</strong></p>

<pre><strong>输入:</strong> 14
<strong>输出:</strong> false 
<strong>解释: </strong><code>14</code> 不是丑数，因为它包含了另外一个质因数&nbsp;<code>7</code>。</pre>

<p><strong>说明：</strong></p>

<ol>
	<li><code>1</code>&nbsp;是丑数。</li>
	<li>输入不会超过 32 位有符号整数的范围:&nbsp;[&minus;2<sup>31</sup>,&nbsp; 2<sup>31&nbsp;</sup>&minus; 1]。</li>
</ol>


 #### 题解
 1、
 ```go
func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else if num%3 == 0 {
			num /= 3
		} else if num%5 == 0 {
			num /= 5
		} else {
			return false
		}
	}
	return true
}
``` 
 时间复杂度O(n),空间复杂度O(1)