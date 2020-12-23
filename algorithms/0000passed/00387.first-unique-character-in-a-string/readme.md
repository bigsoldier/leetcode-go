#### 题目
<p>给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。</p>

<p><strong>案例:</strong></p>

<pre>
s = &quot;leetcode&quot;
返回 0.

s = &quot;loveleetcode&quot;,
返回 2.
</pre>

<p>&nbsp;</p>

<p><strong>注意事项：</strong>您可以假定该字符串只包含小写字母。</p>


 #### 题解
 ```go
func firstUniqChar(s string) int {
	var slice = make([]int,26)
	for i := 0; i < len(s); i++ {
		slice[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if slice[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
```
 时间复杂度O(n),空间复杂度O(n)