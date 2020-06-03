#### 题目
<p>一条包含字母&nbsp;<code>A-Z</code> 的消息通过以下方式进行了编码：</p>

<pre>&#39;A&#39; -&gt; 1
&#39;B&#39; -&gt; 2
...
&#39;Z&#39; -&gt; 26
</pre>

<p>给定一个只包含数字的<strong>非空</strong>字符串，请计算解码方法的总数。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> &quot;12&quot;
<strong>输出:</strong> 2
<strong>解释:</strong>&nbsp;它可以解码为 &quot;AB&quot;（1 2）或者 &quot;L&quot;（12）。
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> &quot;226&quot;
<strong>输出:</strong> 3
<strong>解释:</strong>&nbsp;它可以解码为 &quot;BZ&quot; (2 26), &quot;VF&quot; (22 6), 或者 &quot;BBF&quot; (2 2 6) 。
</pre>


 #### 题解
 动态规划
 dp[i+1] = dp[i]+dp[i-1]
 需要考虑字符为0的情况，如果是10或20可以获得字母，其他的都不成立
 ```go
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	var dp = make([]int,len(s)+1)
	dp[0] = 1
	dp[1] = 1
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				dp[i+1] = dp[i-1]
			} else {
				return 0
			}
		} else if (s[i-1] == '1' ) || s[i-1] == '2' && (s[i] >= '1' && s[i] <= '6') {
			dp[i+1] = dp[i]+dp[i-1]
		} else {
			dp[i+1] = dp[i]
		}
	}
	return dp[len(s)]
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-06-03/009101.png)
 时间复杂度O(n),空间复杂度O(n)