#### 题目
给定一个只包括 **() {} []** 的字符串，判断字符串是否有效。
有效字符串需满足：
1、左括号必须用相同类型的右括号闭合
2、左括号必须以正确的顺序闭合。
注意 空字符串可被认为是有效字符串
**示例1：**
```
输入："()"
输出：true
```
**示例2：**
```
输入："()[]{}"
输出：true
```
**示例3：**
```
输入："(]"
输出：false
```
**示例4：**
```
输入："([)]"
输出：false
```
**示例5：**
```
输入："{[]}"
输出：true
```

#### 题解
很显然，这是道数据结构中的括号匹配问题，用栈的数据结构处理
```go
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]rune,0)
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else if v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' ||
			v == '}' && len(stack) > 0 && stack[len(stack)-1] == '{' ||
			v == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-14/002001.png)
时间复杂度O(n)，空间复杂度O(1)