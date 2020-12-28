#### 题目
<p>给定一个字符串 <em>s</em>，将 <em>s</em> 分割成一些子串，使每个子串都是回文串。</p>

<p>返回符合要求的最少分割次数。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong>&nbsp;&quot;aab&quot;
<strong>输出:</strong> 1
<strong>解释: </strong>进行一次分割就可将&nbsp;<em>s </em>分割成 [&quot;aa&quot;,&quot;b&quot;] 这样两个回文子串。
</pre>


 #### 题解
 动态规划
 
 dp[i]表示从0-i字符需要切割成回文串的次数。
 
 状态转移方程：
 如果s[0:i]是一个回文串，那么dp[i]=0.
 如果s[0:i]不是一个回文串，那么尝试切割边界j，
 如果s[j+1:i]不是回文串，尝试下一个；
 如果s[j+1:i]是回文串，那么dp[i]就是在dp[j]的基础上多一个分割。
 ```go
func minCut(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}
	var dp = make([]int,n)
	for i := 0; i < n; i++ {
		dp[i] = i
	}
	for i := 1; i < n; i++ {
		if checkPalindrome(s, 0, i) {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if checkPalindrome(s, j+1, i) {
				dp[i] = min(dp[i],dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

func checkPalindrome(str string, left, right int) bool {
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++;right--
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
 时间复杂度O(n^2^),空间复杂度O(n)
 
 优化的动态规划
 
 我们发现会重复计算[j:i]是否是回文串,
 那么参考第5题，动态规划回文串的方法
 ```go
func minCut(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}
	var dp = make([]int,n)
	for i := 0; i < n; i++ {
		dp[i] = i
	}
	
	// 动态规划回文串
	checkPalindrome := make([][]bool,n)
	for right := 0; right < n; right++ {
		checkPalindrome[right] = make([]bool,n)
		for left := 0; left <= right; left++ {
			if s[left] == s[right] && (right-left <= 2 || checkPalindrome[left+1][right-1]) {
				checkPalindrome[left][right] = true
			}
		}
	}
	for i := 1; i < n; i++ {
		if checkPalindrome[0][i] {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if checkPalindrome[j+1][i] {
				dp[i] = min(dp[i],dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
 时间复杂度O(n^2^),空间复杂度O(n^2^)