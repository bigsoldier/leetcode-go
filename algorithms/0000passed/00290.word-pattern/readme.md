#### 题目
<p>给定一种规律 <code>pattern</code>&nbsp;和一个字符串&nbsp;<code>str</code>&nbsp;，判断 <code>str</code> 是否遵循相同的规律。</p>

<p>这里的&nbsp;<strong>遵循&nbsp;</strong>指完全匹配，例如，&nbsp;<code>pattern</code>&nbsp;里的每个字母和字符串&nbsp;<code>str</code><strong>&nbsp;</strong>中的每个非空单词之间存在着双向连接的对应规律。</p>

<p><strong>示例1:</strong></p>

<pre><strong>输入:</strong> pattern = <code>&quot;abba&quot;</code>, str = <code>&quot;dog cat cat dog&quot;</code>
<strong>输出:</strong> true</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong>pattern = <code>&quot;abba&quot;</code>, str = <code>&quot;dog cat cat fish&quot;</code>
<strong>输出:</strong> false</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入:</strong> pattern = <code>&quot;aaaa&quot;</code>, str = <code>&quot;dog cat cat dog&quot;</code>
<strong>输出:</strong> false</pre>

<p><strong>示例&nbsp;4:</strong></p>

<pre><strong>输入:</strong> pattern = <code>&quot;abba&quot;</code>, str = <code>&quot;dog dog dog dog&quot;</code>
<strong>输出:</strong> false</pre>

<p><strong>说明:</strong><br>
你可以假设&nbsp;<code>pattern</code>&nbsp;只包含小写字母，&nbsp;<code>str</code>&nbsp;包含了由单个空格分隔的小写字母。&nbsp; &nbsp;&nbsp;</p>


 #### 题解
 使用哈希表将pattern存入key，将对应位置的str存入value，然后依次比较值
 ```go
 func wordPattern(pattern string, str string) bool {
 	strs := strings.Split(str," ")
 	var m = make(map[uint8]string)
 	for i := 0; i < len(pattern); i++ {
 		if val,ok := m[pattern[i]];ok {
 			if val != strs[i] {
 				return false
 			}
 		} else {
 			m[pattern[i]] = strs[i]
 		}
 	}
 	return true
 }
```
 但是这种有坏情况
 ```text
 pattern: abba
 str: cat cat cat cat
```
 这种情况下不通过。
 那么就将哈希表的key和value反过来，把str对应位置作为key，把pattern作为value
 ```go
func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str," ")
	if len(strs) != len(pattern) {
		return false
	}
	var m = make(map[string]byte)
	for i := 0; i < len(strs); i++ {
		if val, ok := m[strs[i]]; ok {
			if val != pattern[i] {
				return false
			}
		} else {
			m[strs[i]] = pattern[i]
		}
	}
	return true
}
```
但是这种有坏情况
 ```text
 pattern: abba
 str: dog cat cat fish
```

 所以要将两种方法结合起来
 ```go
func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str," ")
	if len(strs) != len(pattern) {
		return false
	}
	var ch2word = make(map[byte]string)
	var word2ch = make(map[string]byte)
	for i, word := range strs {
		ch := pattern[i]
		if (word2ch[word] > 0 && word2ch[word] != ch) || (ch2word[ch] != "" && ch2word[ch] != word) {
			return false
		}
		word2ch[word] = ch
		ch2word[ch] = word
	}
	return true
}
```
 时间复杂度O(n+m),空间复杂度O(n+m)