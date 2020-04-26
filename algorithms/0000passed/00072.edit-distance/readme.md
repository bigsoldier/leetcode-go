#### 题目
<p>给定两个单词&nbsp;<em>word1</em> 和&nbsp;<em>word2</em>，计算出将&nbsp;<em>word1</em>&nbsp;转换成&nbsp;<em>word2 </em>所使用的最少操作数&nbsp;。</p>

<p>你可以对一个单词进行如下三种操作：</p>

<ol>
	<li>插入一个字符</li>
	<li>删除一个字符</li>
	<li>替换一个字符</li>
</ol>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> word1 = &quot;horse&quot;, word2 = &quot;ros&quot;
<strong>输出:</strong> 3
<strong>解释:</strong> 
horse -&gt; rorse (将 &#39;h&#39; 替换为 &#39;r&#39;)
rorse -&gt; rose (删除 &#39;r&#39;)
rose -&gt; ros (删除 &#39;e&#39;)
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> word1 = &quot;intention&quot;, word2 = &quot;execution&quot;
<strong>输出:</strong> 5
<strong>解释:</strong> 
intention -&gt; inention (删除 &#39;t&#39;)
inention -&gt; enention (将 &#39;i&#39; 替换为 &#39;e&#39;)
enention -&gt; exention (将 &#39;n&#39; 替换为 &#39;x&#39;)
exention -&gt; exection (将 &#39;n&#39; 替换为 &#39;c&#39;)
exection -&gt; execution (插入 &#39;u&#39;)
</pre>


 #### 题解
 对于字符串ab的插入、删除、修改操作可以分解成3个操作：
 - 对a插入(等价于对b删除)
 - 对b插入(等价于对a删除)
 - 对a修改(等价于对b修改)
 建立状态转移方程,dp[i][j]表示到a的前i个字母和b的前j个字母之间的编辑距离。
 - dp[i-1][j]: 对于b的第j个字母，在a的末尾添加一个字符，即dp[i][j]=dp[i-1][j]+1
 - dp[i][j-1]: 对于a的第i个字母，在b的末尾添加一个字符，即dp[i][j]=dp[i][j-1]+1
 - dp[i-1][j-1]: 对于a的第i个字母和b的第j个字母，如果字母相同，dp[i][j]=dp[i-1][j-1]，如果不同需要加一个修改操作，dp[i][j]=dp[i-1][j-1]+1
 所以，
 最后字母相同时, dp[i][j] = min(dp[i-1][j]+1,dp[i][j-1]+1,dp[i-1][j-1])
 最后字母不同时, dp[i][j] = min(dp[i-1][j]+1,dp[i][j-1]+1,dp[i-1][j-1]+1)
 处理一下边界：如果a的长度为0,返回b的长度;如果b的长度为0,返回a的长度
 ```go
func minDistance(word1 string, word2 string) int {
	m,n := len(word1),len(word2)
	if m*n == 0 {
		return m+n
	}

	var dp = make([][]int,m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int,n+1)
	}

	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			left := dp[i-1][j]+1
			down := dp[i][j-1]+1
			leftDown := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftDown += 1
			}
			var min = left
			if min > down {
				min = down
			}
			if min > leftDown {
				min = leftDown
			}
			dp[i][j] = min
		}
	}
	return dp[m][n]
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-26/007201.png)
 时间复杂度O(n^2^),空间复杂度O(n^2^)