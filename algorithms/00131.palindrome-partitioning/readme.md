#### 题目
<p>给定一个字符串 <em>s</em>，将<em> s </em>分割成一些子串，使每个子串都是回文串。</p>

<p>返回 <em>s</em> 所有可能的分割方案。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong>&nbsp;&quot;aab&quot;
<strong>输出:</strong>
[
  [&quot;aa&quot;,&quot;b&quot;],
  [&quot;a&quot;,&quot;a&quot;,&quot;b&quot;]
]</pre>


 #### 题解
 ```go
func partition(s string) [][]string {
	var result [][]string
	part(s,0,len(s),nil,&result)
	return result
}

func part(s string, start, end int, ret []string, result *[][]string) {
	if start == end {
		*result = append(*result, ret)
		return
	}
	for i := start; i < end; i++ {
		if !isParam(s, start, i) {
			continue
		}
		part(s,i+1,end, append(ret, s[start:i+1]),result)
	}
}

// 是否是回文字符
func isParam(str string, start, end int) bool {
	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}
```
 时间复杂度O(n^2^)，空间复杂度O(1)