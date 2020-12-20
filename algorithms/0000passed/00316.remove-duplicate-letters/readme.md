#### 题目
<p>给你一个仅包含小写字母的字符串，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证返回结果的字典序最小（要求不能打乱其他字符的相对位置）。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> <code>&quot;bcabc&quot;</code>
<strong>输出:</strong> <code>&quot;abc&quot;</code>
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> <code>&quot;cbacdcbc&quot;</code>
<strong>输出:</strong> <code>&quot;acdb&quot;</code></pre>

<p>&nbsp;</p>

<p><strong>注意：</strong>该题与 1081 <a href="https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters">https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters</a> 相同</p>


 #### 题解
 题意：字典序最小，意味着最后获取到的不重复的字母顺序最小，例如，adcad最后的结果是acd
 
 `cbacdcbc`的结果是acdb。
 acd是增序且d只有一个。
 
 ```go
func removeDuplicateLetters(s string) string {
	var left [26]int
	for _,ch := range s {
		left[ch-'a']++
	}
	var stack []byte
	var inStack [26]bool
	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a']--
	}
	return string(stack)
}
```
 时间复杂度O(n),空间复杂度O(n)