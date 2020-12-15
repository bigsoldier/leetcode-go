#### 题目
<p>给定一个非负整数&nbsp;<code>N</code>，找出小于或等于&nbsp;<code>N</code>&nbsp;的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。</p>

<p>（当且仅当每个相邻位数上的数字&nbsp;<code>x</code>&nbsp;和&nbsp;<code>y</code>&nbsp;满足&nbsp;<code>x &lt;= y</code>&nbsp;时，我们称这个整数是单调递增的。）</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> N = 10
<strong>输出:</strong> 9
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> N = 1234
<strong>输出:</strong> 1234
</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入:</strong> N = 332
<strong>输出:</strong> 299
</pre>

<p><strong>说明:</strong> <code>N</code>&nbsp;是在&nbsp;<code>[0, 10^9]</code>&nbsp;范围内的一个整数。</p>


 #### 题解
 暴力：
 判断当前数字是否是递增的，然后N--
 贪心：
 从高到低位，第一个满足str[i]>str[i+1]，那么把str[i]-1，再把后面的位置置9.
 但是由于减少了str[i]后，可能不满足str[i-1]<str[i].
 ```text
 N = 2332
 输出：2299
 ```
 这样需要遍历之前的位置
 ```go
func monotoneIncreasingDigits(N int) int {
	s := []byte(strconv.Itoa(N))
	i := 1
	for i < len(s) && s[i-1] <= s[i] {
		i++
	}

	if i < len(s) {
		for i > 0 && s[i-1] > s[i] {
			s[i-1]--
			i--
		}
		for i++;i<len(s);i++ {
			s[i] = '9'
		}
	}
	ret,_ := strconv.Atoi(string(s))
	return ret
}
```