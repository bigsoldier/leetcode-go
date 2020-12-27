#### 题目
<p>给定两个字符串&nbsp;<em><strong>s&nbsp;</strong></em>和&nbsp;<strong><em>t</em></strong>，判断它们是否是同构的。</p>

<p>如果&nbsp;<em><strong>s&nbsp;</strong></em>中的字符可以被替换得到&nbsp;<strong><em>t&nbsp;</em></strong>，那么这两个字符串是同构的。</p>

<p>所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> <strong><em>s</em></strong> = <code>&quot;egg&quot;, </code><strong><em>t = </em></strong><code>&quot;add&quot;</code>
<strong>输出:</strong> true
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> <strong><em>s</em></strong> = <code>&quot;foo&quot;, </code><strong><em>t = </em></strong><code>&quot;bar&quot;</code>
<strong>输出:</strong> false</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入:</strong> <strong><em>s</em></strong> = <code>&quot;paper&quot;, </code><strong><em>t = </em></strong><code>&quot;title&quot;</code>
<strong>输出:</strong> true</pre>

<p><strong>说明:</strong><br>
你可以假设&nbsp;<em><strong>s&nbsp;</strong></em>和 <strong><em>t </em></strong>具有相同的长度。</p>


 #### 题解
 类似于双射，由s的map记录s->t的映射，由t的map记录t->s的映射
 ```go
func isIsomorphic(s string, t string) bool {
	var s2t,t2s = map[byte]byte{},map[byte]byte{}
	for i := 0; i < len(s); i++ {
		x,y := s[i],t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] =y
		t2s[y] = x
	}
	return true
}
```
 时间复杂度O(n),空间复杂度O(n)