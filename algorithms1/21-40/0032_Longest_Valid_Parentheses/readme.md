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
dp[i] 表示以 i 结尾的最长有效括号长度。一个有效括号最后以右括号结尾。
一个有效括号会有两种情况：
- ...() ： 左右括号相邻 (s[i]==')' && s[i-1]=='(')，那么 dp[i] = dp[i-2]+2
- ...)) ： 如果存在一个左括号使得 s[i-dp[i-1]-1]=='('，那么 dp[i]=dp[i-1]+dp[i-dp[i-1]-2]+2
```go
func longestValidParentheses(s string) int {
	var dp = make([]int,len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' { // ...()
				if i >= 2 {
					dp[i] = dp[i-2]+2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] +dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
	}

	var max int
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-03/003201.png)
时间复杂度O(n),空间复杂度O(n)

3、栈
栈内记录左括号的索引，每当出现右括号，移除栈顶元素，然后再用索引减去上一个左括号的位置 i-stack[len(stack)-1]
```go
func longestValidParentheses(s string) int {
	var stack = make([]int,0)
	stack = append(stack, -1)
	var max int
	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				num := stack[len(stack)-1]
				if i - num > max {
					max = i - num
				}
			} else {
				stack = append(stack,i)
			}
		}
	}
	return max
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-03/003202.png)
时间复杂度O(n),空间复杂度O(n)

4、左右横跳
一个有效的括号，它的左括号和右括号必然相等。（如果在有效括号左边多左括号，那么从右遍历时能读取出来，同理在有效括号右边多右括号，从左遍历时能读取出来）
从左向右遍历时，如果右括号多，那么必然不是有效括号，数据清零，继续遍历。
```go
func longestValidParentheses(s string) int {
	// 从左向右遍历，从右向左遍历
	var left, right int
	var max int
	for _, v := range s {
		switch v {
		case '(':
			left++
		case ')':
			right++
		}

		if right > left {
			left,right = 0,0
		} else if left == right {
			if left+right > max {
				max = left + right
			}
		}
	}

	left,right = 0,0
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '(':
			left++
		case ')':
			right++
		}
		if left > right {
			left,right = 0,0
		} else if left == right {
			if left+right > max {
				max = left + right
			}
		}
	}

	return max
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-04/003203.png)
时间复杂度O(n)，空间复杂度O(1)