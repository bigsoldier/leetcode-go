#### 题目
实现 ![strStr函数](https://baike.baidu.com/item/strstr/811469)
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找到 needle 字符串出现的第一个位置（从0开始）。如果不存在，则返回 -1。
**示例1：**
```
输入：haystack="hello",needle="ll"
输出：2
```
**示例2：**
```
输入：haystack="aaaaa",needle="bba"
输出：-1
```

#### 题解
暴力法
移动 haystack 的坐标值，比较坐标值和 needle的值。
```go
func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		var j int
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-18/002801.png)
空间复杂度O(n^2^)，时间复杂度O(1)
