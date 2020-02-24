#### 题目
编写一个函数来查找字符串数组中的最长公共前缀
如果不存在公共前缀，返回空字符串 ""
**实例1：**
```
输入：["flower","flow","flight"]
输出："fl"
```
**实例2：**
```
输入：["dog","racecar","car"]
输出：""
```

#### 题解
1、垂直扫描
比较每个字符相同的位置，判断如果当前位置是该字符串的末尾或当前字符串的元素不相等，返回索引位置的字符串
```go
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var index int
	for index < len(strs[0]) {
		ch := strs[0][index]
		for j := 1; j < len(strs); j++ {
			if index >= len(strs[j]) || strs[j][index] != ch {
				return strs[0][:index]
			}
		}
		index++
	}
	return strs[0][:index]
}
```
![https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-12/001401.png]()
时间复杂度：O(n^2^)，空间复杂度：O(1)

2、横向扫描
比较第一个和第二个字符串的相同前缀，然后拿这个前缀和第三个字符串比较
```go
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefix = strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}
```
![https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-12/001402.png]()
时间复杂度：O(n^2^)，空间复杂度：O(1)