#### 题目
给定一个字符串 *s*，找到 *s* 中最长的回文子串。你可以假定 *s* 的最大长度为1000.
**示例1：**
```$xslt
输入： babad
输出：bab
注意：aba也是有效答案
```
**示例2：**
```$xslt
输入：cbbd
输出：bb
```

#### 题解
1、暴力法
```go
func longestPalindrome(s string) string {
	var maxLen int
	var maxStr string
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)+1; j++ {
			len := isPalindrome(s[i:j])
			if len > maxLen {
				maxStr = s[i:j]
				maxLen = len
			}
		}
	}
	return maxStr
}

func isPalindrome(s string) int {
	var newStr string
	for i := len(s)-1; i >= 0; i--{
		newStr += s[i:i+1]
	}
	if s == newStr {
		return len(s)
	}
	return 0
}
```
很显然，超时了。
时间复杂度：O(n^3^)
空间复杂度：O(1)

2、动态规划
如果子串 P 是回文子串，那么P(i,j)=(P(i+1,j-1) 且 S(i)=S(j))。
那么倒着查询就行
```go
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	// P[i][j]表示回文，P[i+1][j-1]也是回文
	longest := s[0:1]
	for i := 1; i < len(s); i++ {
		for rightStep := 0; rightStep < 2; rightStep++ {
			for p, q := i-1, i+rightStep; p >= 0 && q < len(s) && s[p] == s[q]; {
				if q-p+1 > len(longest) {
					longest = s[p : q+1]
				}
				p--
				q++
			}
		}
	}
	return longest
}
```

