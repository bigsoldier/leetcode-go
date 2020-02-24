#### 题目

给定一个字符串，请你找出其中不含有重复字符的 **最长子串** 的长度。

**示例 1:**
```$xslt
输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```

**示例 2:**
```$xslt
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
```

**示例 3:**
```$xslt
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
```

#### 题解
1. 暴力法
简单来说就是列出字符串中所有的子串
```go
func lengthOfLongestSubstring(s string) int {
	var maxLen int
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isUnique(s[i:j])  {
				if len(s[i:j]) > maxLen {
					maxLen = len(s[i:j])
				}
			} else { // 如果该子串重复，继续添加后续字符也还是不重复
				break
			}
		}
	}
	return maxLen
}

func isUnique(s string) bool {
	var m = make(map[int32]bool)
	for _, vChar := range s {
		if _, ok := m[vChar]; ok {
			return false
		}
		m[vChar] = true
	}
	return true
}
```

![暴力法](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-01-07/000301.png)

时间复杂度 O(n^3^)
空间复杂度O(min(n, m))，我们需要 O(k) 的空间来检查子字符串中是否有重复字符，其中 k 表示 Set 的大小。而 Set 的大小取决于字符串 n 的大小以及字符集/字母 m 的大小。

2. 滑动窗口
暴力法很简单，可是太慢了，浪费的时间在会检查重复的字符。如果索引 i 到 j-1 之间的字符串 s~ij~  已经检查没有重复字符了，那么就只需要检查 s~j~ 是否在 s~ij~ 。
滑动窗口的右边界不断的右移，只要没有重复的字符，就不用的向右扩大窗口边界。一旦出现了重复字符，此时先计算一下滑动窗口的大小，记录下来。再需要缩小左边界。
直到重复的字符移出了左边界。接着又可以开始移动滑动窗口的右边界。以此类推，不断的刷新记录的窗口大小。

```go
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int // 字符出现的次数
	var maxLen int
	left,right := 0,-1 // 左右下标
	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 { // 右标如果没有出现过，记录下来
			freq[s[right+1]] ++
			right ++
		} else { // 右标出现过，移动左标
			freq[s[left]] --
			left ++
		}
		if maxLen < right-left+1 {
			maxLen = right-left+1
		}
	}
	return maxLen
}
```

![滑动窗口法](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-01-07/000302.png)

时间复杂度 O(2n) = O(n) ,最糟糕的情况，每个字符都被访问一次，即所有相同字符串。
空间复杂度 O(min(n,m))

3. 优化的滑动窗口
方法2最多需要执行2n个步骤，可以考虑记录每个字符上次出现的位置。如果 s~j~ 在 s~ij~ 中有重复，找到 s~j'~ 的索引，左标移到 s~j'~ ,这样就跳过了重复字符覆盖的区域
```go
func lengthOfLongestSubstring(s string) int {
	var count = make(map[rune]int) // 存放字符的位置
	var max int
	left := 0
	for k, v := range s {
		if val, ok := count[v]; ok {
			if left < val {
				left = val
			}
		}
		if max < k-left+1 {
			max = k - left + 1
		}
		count[v] = k + 1 // 更新位置
	}
	return max
}
``` 

时间复杂度 O(n)，右索引会遍历字符串
空间复杂度 O(m)，字符串的长度