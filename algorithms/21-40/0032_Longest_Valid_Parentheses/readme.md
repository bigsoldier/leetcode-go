#### 题目
给定一个只包含 '(' 和 ')'的字符串，找出最长的包含有效括号的子串的长度。
**示例1：**
```
输入："(()"
输出：2
解释：最长有效括号子串为“（）”
```
**示例2：**
```
输入：")()())"
输出：4
解释：最长有效括号子串为“（）（）”
```

#### 题解
1、暴力法
获取字符串里每个可能的短字符串，然后判断是否是有效的括号
```go
func longestValidParentheses(s string) int {
	var max int
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s)+1; j++ {
			str := s[i:j]
			//fmt.Println(i,j,str)
			if isValid(str) && j-i > max {
				max = j-i
			}
		}
	}
	return max
}

// 判断字符串是否是有效的括号
func isValid(s string) bool {
	var stack int
	for _, str := range s {
		bt := string(str)
		switch bt {
		case "(":
			stack++
		case ")":
			if stack > 0 {
				stack--
			} else {
				return false
			}
		}
	}
	if stack == 0 {
		return true
	}
	return false
}
```
时间复杂度O(n^3^)，空间复杂度O(1)

2、动态规划
