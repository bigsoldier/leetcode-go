#### 题目
<p>给定一个字符串 <em><strong>s</strong></em>，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入: </strong><code>&quot;aacecaaa&quot;</code>
<strong>输出:</strong> <code>&quot;aaacecaaa&quot;</code>
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入: </strong><code>&quot;abcd&quot;</code>
<strong>输出:</strong> <code>&quot;dcbabcd&quot;</code></pre>


 #### 题解
 KMP算法
 ```go
func shortestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	i := getIndex(s+"#"+reverse(s))
	return reverse(s[i:]) + s
}

// 反转字符串
func reverse(s string) string {
	n := len(s)
	bytes := make([]byte,n)
	for i := 0; i < n; i++ {
		bytes[i] = s[n-i-1]
	}
	return string(bytes)
}

func getIndex(pattern string) int {
	n := len(pattern)
	next := make([]int,n)
	k,i := 0,1
	for i < n {
		if pattern[i] == pattern[k] {
			next[i] = k+1
			k = next[i]
			i++
		} else if k == 0{
			next[i] = k
			i++
		} else {
			k = next[k-1]
		}
	}
	return next[n-1]
}
```
 时间复杂度O(n),空间复杂度O(n)