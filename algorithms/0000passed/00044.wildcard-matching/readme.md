#### 题目
<p>给定一个字符串&nbsp;(<code>s</code>) 和一个字符模式&nbsp;(<code>p</code>) ，实现一个支持&nbsp;<code>&#39;?&#39;</code>&nbsp;和&nbsp;<code>&#39;*&#39;</code>&nbsp;的通配符匹配。</p>

<pre>&#39;?&#39; 可以匹配任何单个字符。
&#39;*&#39; 可以匹配任意字符串（包括空字符串）。
</pre>

<p>两个字符串<strong>完全匹配</strong>才算匹配成功。</p>

<p><strong>说明:</strong></p>

<ul>
	<li><code>s</code>&nbsp;可能为空，且只包含从&nbsp;<code>a-z</code>&nbsp;的小写字母。</li>
	<li><code>p</code>&nbsp;可能为空，且只包含从&nbsp;<code>a-z</code>&nbsp;的小写字母，以及字符&nbsp;<code>?</code>&nbsp;和&nbsp;<code>*</code>。</li>
</ul>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong>
s = &quot;aa&quot;
p = &quot;a&quot;
<strong>输出:</strong> false
<strong>解释:</strong> &quot;a&quot; 无法匹配 &quot;aa&quot; 整个字符串。</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong>
s = &quot;aa&quot;
p = &quot;*&quot;
<strong>输出:</strong> true
<strong>解释:</strong>&nbsp;&#39;*&#39; 可以匹配任意字符串。
</pre>

<p><strong>示例&nbsp;3:</strong></p>

<pre><strong>输入:</strong>
s = &quot;cb&quot;
p = &quot;?a&quot;
<strong>输出:</strong> false
<strong>解释:</strong>&nbsp;&#39;?&#39; 可以匹配 &#39;c&#39;, 但第二个 &#39;a&#39; 无法匹配 &#39;b&#39;。
</pre>

<p><strong>示例&nbsp;4:</strong></p>

<pre><strong>输入:</strong>
s = &quot;adceb&quot;
p = &quot;*a*b&quot;
<strong>输出:</strong> true
<strong>解释:</strong>&nbsp;第一个 &#39;*&#39; 可以匹配空字符串, 第二个 &#39;*&#39; 可以匹配字符串 &quot;dce&quot;.
</pre>

<p><strong>示例&nbsp;5:</strong></p>

<pre><strong>输入:</strong>
s = &quot;acdcb&quot;
p = &quot;a*c?b&quot;
<strong>输入:</strong> false</pre>


 #### 题解
 动态规划
 ```go
func isMatch(s string, p string) bool {
	var sLen,pLen = len(s),len(p)
	var dp = make([][]bool,sLen+1)
	for i := 0; i < sLen+1; i++ {
		dp[i] = make([]bool,pLen+1)
	}

	dp[0][0] = true
	for i := 1; i <= pLen; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if p[j-1] != '*' {
				// 单字符匹配并且前面的字符也需要匹配上
				dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1]=='?')
			} else {
				// 当 p[j-1] == '*' 时
				//   要么，dp[i-1][j] == true，意味着，
				//   当 s[:i] 与 p[:j+1] 匹配，且p[j] == '*' 的时候，
				//   s[:i] 后接任意字符串的 s[:i+1] 仍与 p[:j+1] 匹配。
				//   要么，dp[i][j-1] == true，意味着，
				//   当 s[:i+1] 与 p[:j] 匹配后
				//   在 p[:j] 后添加'*'，s[:i+1] 与 p[:j+1] 任然匹配
				//   因为， '*' 可以匹配空字符。
				dp[i][j] = dp[i-1][j] || dp[i][j-1] // *匹配上空格；*匹配上一串字符
			}
		}
	}

	return dp[sLen][pLen]
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-23/004401.png)
时间复杂度O(n^2^),空间复杂度O(n^2^)