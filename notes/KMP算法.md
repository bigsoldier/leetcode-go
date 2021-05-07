### 暴力匹配
有一个字符串S和一个模式串P，现在要查找P在S中的位置。

如果用暴力匹配的思路，假设字符串匹配到i位置，模式串匹配到j位置，现有：
- 如果当前字符串匹配成功(即S[i]=P[j])，则i++，j++，继续匹配下一个字符
- 如果失败(即S[i]!=P[j])，令i=i-(j-1)，j=0，相当于每次匹配失败，i回溯，j重置

```go
func violentMatch(s, p string) int {
	sLen,pLen := len(s),len(p)
	var i,j int
	for i < sLen && j < pLen {
		if s[i] == p[j] {
			i++
			j++
		} else {
i = i -j + 1
j = 0
}
}
if j == pLen {
return i-j
}
return -1
}
```

**举个例子**

给定文本串S：“BBC ABCDAB ABCDABCDABDE”，模式串P：“ABCDABD”。现在要拿模式串去匹配文本串

1、S[0]为B，P[0]为A，不匹配，右移

S: BBC ABCDAB ABCDABCDABDE

P: ABCDABD

2、直到S[4]=P[0]，i++，j++

S: BBC ABCDAB ABCDABCDABDE

P:     ABCDABD

3、S[10]为空，P[6]=D，不匹配，所以文本串回溯到S[5]，模式串回溯到P[0]

S: BBC ABCDAB ABCDABCDABDE

P:      ABCDABD

### KMP算法

暴力匹配有没有可以优化的地方？

每次匹配失败后，i都回溯，但是回溯的地方是我们遍历过的，这里可以被优化。

同样，假设字符串匹配到i位置，模式串匹配到j位置，现有：

- 如果当前字符串匹配成功(即S[i]=P[j]) 或 `j=-1`，则i++，j++，继续匹配下一个字符
- 如果失败(即S[i]!=P[j]) 且 `j!=-1`,则令i不变，j=next[j]。意味着匹配失败后，模式串P相对于字符串S向右移动j-next[j]位

这里的next数组意义：代表当前字符串有多大长度的相同前缀。如果next[j]=k，代表j之前的字符串中有最大长度为k的相同前缀。

也就意味着某个字符匹配失败后，该字符对应的next值会告诉你下一步匹配中，模式串应该跳到哪个位置，如果next[j]=0或-1，则跳到模式串开头。

```go
func kmpSearch(s, p string) int {
var i, j int
sLen, pLen := len(s), len(p)
next := subStringRepetition(p)
for i < sLen && j < pLen {
if j == -1 || s[i] == p[j] {
i++; j++
} else {
j = next[j]
}
}
if j == pLen {
return i - j
}
return -1
}

func subStringRepetition(str string) []int {
n := len(str)
next := make([]int, n+1)
next[0] = -1
i,j := 0, -1
for i < n {
if j == -1 || str[i] == str[j] {
i++; j++
next[i] = j
} else {
j = next[j]
}
}
return next
}
```