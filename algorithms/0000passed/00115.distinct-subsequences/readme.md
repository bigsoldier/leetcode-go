#### 题目
<p>给定一个字符串&nbsp;<strong>S&nbsp;</strong>和一个字符串&nbsp;<strong>T</strong>，计算在 <strong>S</strong> 的子序列中 <strong>T</strong> 出现的个数。</p>

<p>一个字符串的一个子序列是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，<code>&quot;ACE&quot;</code>&nbsp;是&nbsp;<code>&quot;ABCDE&quot;</code>&nbsp;的一个子序列，而&nbsp;<code>&quot;AEC&quot;</code>&nbsp;不是）</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入: </strong>S = <code>&quot;rabbbit&quot;</code>, T = <code>&quot;rabbit&quot;
<strong>输出:</strong>&nbsp;3
</code><strong>解释:
</strong>
如下图所示, 有 3 种可以从 S 中得到 <code>&quot;rabbit&quot; 的方案</code>。
(上箭头符号 ^ 表示选取的字母)

<code>rabbbit</code>
^^^^ ^^
<code>rabbbit</code>
^^ ^^^^
<code>rabbbit</code>
^^^ ^^^
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入: </strong>S = <code>&quot;babgbag&quot;</code>, T = <code>&quot;bag&quot;
<strong>输出:</strong>&nbsp;5
</code><strong>解释:
</strong>
如下图所示, 有 5 种可以从 S 中得到 <code>&quot;bag&quot; 的方案</code>。 
(上箭头符号 ^ 表示选取的字母)

<code>babgbag</code>
^^ ^
<code>babgbag</code>
^^    ^
<code>babgbag</code>
^    ^^
<code>babgbag</code>
  ^  ^^
<code>babgbag</code>
    ^^^</pre>


 #### 题解
 迭代
 ```go
func numDistinct(s string, t string) int {
	if len(t) == 0 {
		return 0
	}
	return distinct(s,t)
}

func distinct(s, t string) int {
	var count int
	if len(t) == 0 {
		count+=1
		return count
	}
	if len(s) < len(t) {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if s[i] == t[0] {
			count += distinct(s[i+1:],t[1:])
		}
	}
	return count
}
```
 时间复杂度O(n^m^),空间复杂度O(1)
 
 动态规划
 
 dp[i][j]表示t的第i个字符和s的第j个字符符合
 ```go
func numDistinct(s string, t string) int {
	var dp = make([][]int,len(t)+1)
	for i := 0; i < len(t)+1; i++ {
		dp[i] = make([]int,len(s)+1)
	}
	for j := 0; j < len(s)+1; j++ {
		dp[0][j] = 1
	}
	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]
}
```
 时间复杂度O(n*m),空间复杂度O(n*m)