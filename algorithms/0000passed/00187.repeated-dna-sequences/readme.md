#### 题目
<p>所有 DNA 都由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：&ldquo;ACGAATTCCG&rdquo;。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。</p>

<p>编写一个函数来查找 DNA 分子中所有出现超过一次的 10 个字母长的序列（子串）。</p>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>

<pre><strong>输入：</strong>s = &quot;AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT&quot;
<strong>输出：</strong>[&quot;AAAAACCCCC&quot;, &quot;CCCCCAAAAA&quot;]</pre>


 #### 题解
 1、哈希表
 ```go
func findRepeatedDnaSequences(s string) (ans []string) {
	const L = 10
	m := map[string]int{}
	for i := 0; i < len(s)-L+1; i++ {
		if v := m[s[i:i+L]]; v == 1 {
			ans = append(ans, s[i:i+L])
		}
		m[s[i:i+L]]++
	}
	return 
}
```
 时间复杂度O(n),空间复杂度O(n)
 