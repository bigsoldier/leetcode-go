#### 题目
<p>给定一个字符串 <em>s</em>，将<em> s </em>分割成一些子串，使每个子串都是回文串。</p>

<p>返回 <em>s</em> 所有可能的分割方案。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong>&nbsp;&quot;aab&quot;
<strong>输出:</strong>
[
  [&quot;aa&quot;,&quot;b&quot;],
  [&quot;a&quot;,&quot;a&quot;,&quot;b&quot;]
]</pre>


 #### 题解
 回溯
 
 将字符串分为两部分，第一部分判断是否是回文串，第二部分递归重复第一部分判断。
 ```go
func partition(s string) [][]string {
	var result [][]string
	backtracking(s,0,len(s),[]string{},&result)
	return result
}

func backtracking(str string, start, end int, part []string, result *[][]string) {
	if start == end {
		s := make([]string,len(part))
		copy(s,part)
		*result = append(*result, s)
		return
	}
	for i := start; i < end; i++ {
		if !checkPalindrome(str, start, i) {
			continue
		}
		backtracking(str,i+1,end, append(part, str[start:i+1]),result)
	}
}

func checkPalindrome(str string, left, right int) bool {
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
```
 时间复杂度O(n^2^)，空间复杂度O(1)
 
 参考第0005题最长回文串，使用动态规划来判断[i][j]是否是回文。
 如果[i][j]是回文，那么[i+1][j-1]也是回文。
 将重复计算是否是回文这部分提出用动态规划处理
 ```go
func partition(s string) [][]string {
	var result [][]string
	var dp = make([][]bool,len(s))
	for right := 0; right < len(s); right++ {
		dp[right] = make([]bool,len(s))
		for left := 0; left <= right; left++ {
			if s[left] == s[right] && (right-left <= 2 || dp[left+1][right-1]) {
				dp[left][right] = true
			}
		}
	}
	backtracking(s,0,len(s),dp,nil,&result)
	return result
}

func backtracking(str string, start, end int, dp [][]bool, path []string, result *[][]string) {
	if start == end {
		s := make([]string,len(path))
		copy(s,path)
		*result = append(*result, s)
		return
	}
	for i := start; i < end; i++ {
		if !dp[start][i] {
			continue
		}
		backtracking(str,i+1,end,dp, append(path, str[start:i+1]),result)
	}
}
```
 时间复杂度O(n^2^)，空间复杂度O(1)
 