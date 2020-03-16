#### 题目
给定一个字符串 **s** 和一些长度相同的单词 **words**。找出 **s** 中恰好可以由 **words** 中所有单词串联形成的子串的起始位置。
注意子串要与 **words** 中的单词完全匹配，中间不能有其他字符，但不需要考虑 **words** 中单词串联的顺序。
**示例1：**
```
输入：s = "barfoothefoobarman",words = ["foo","bar]
输出：[0,9]
解释：输出顺序不重要，[9,0]也是有效答案
```
**示例2：**
```
输入：s = "wordgoodgoodgoodbestword",words = ["words","good","best","word"]
输出：[]
```

#### 题解
1、暴力法
列出 **words** 所有的可能性，有 n！个，然后再滑动窗口找符合条件的字符串。

2、将 **words** 放入map中，并累计出现的次数，然后从 **s** 头扫描，每次判断字符串数组里的字符串是否用完。
```go
// words的长度相同
func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return nil
	}
	var wordsMap = make(map[string]int)
	for _, word := range words {
		wordsMap[word]++
	}

	var res []int
	wordLen, totalLen := len(words[0]), len(words[0])*len(words) // 单词长度
	tmpCounter := copyMap(wordsMap)
	for i, start := 0, 0; i < len(s)-wordLen+1 && start < len(s)-wordLen+1; i++ {
		word := s[i : i+wordLen]
		if tmpCounter[word] > 0 {
			tmpCounter[word]--
			if checkWords(tmpCounter) && (i+wordLen-start == totalLen) {
				res = append(res, start)
				continue
			}
			i = i + wordLen - 1
		} else {
			start ++
			i = start -1
			tmpCounter = copyMap(wordsMap)
		}

	}
	return res
}

// 检查是否使用完
func checkWords(s map[string]int) bool {
	flag := true
	for _, v := range s {
		if v > 0 {
			flag = false
			break
		}
	}
	return flag
}

func copyMap(s map[string]int) map[string]int {
	c := make(map[string]int)
	for k, v := range s {
		c[k] = v
	}
	return c
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-24/003001.png)

3、滑动窗口
```go
func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return nil
	}
	var res []int
	wordLen := len(words[0])
	wordNum := len(words)
	if len(s) < wordLen {
		return nil
	}

	var mapWords = make(map[string]int)
	for _, word := range words {
		mapWords[word]++
	}

	for i := 0; i < wordLen; i++ {	// 每wordLen长度间隔查询
		left,right,count := i,i,0
		var tmpMap = make(map[string]int)
		for right+wordLen <= len(s) {
			word := s[right:right+wordLen]
			right += wordLen
			if _, ok := mapWords[word]; !ok {
				count = 0
				left = right
				tmpMap = make(map[string]int)
			} else {
				tmpMap[word] ++
				count ++
				for tmpMap[word] > mapWords[word] {
					count --
					tmpMap[s[left:left+wordLen]]--
					left += wordLen
				}
				if count == wordNum {
					res = append(res,left)
				}
			}
		}
	}
	return res
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-24/003002.png)