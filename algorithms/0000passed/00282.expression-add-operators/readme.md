#### 题目
<p>给定一个仅包含数字&nbsp;<code>0-9</code>&nbsp;的字符串和一个目标值，在数字之间添加<strong>二元</strong>运算符（不是一元）<code>+</code>、<code>-</code>&nbsp;或&nbsp;<code>*</code>&nbsp;，返回所有能够得到目标值的表达式。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> <code><em>num</em> = </code>&quot;123&quot;, <em>target</em> = 6
<strong>输出: </strong>[&quot;1+2+3&quot;, &quot;1*2*3&quot;] 
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> <code><em>num</em> = </code>&quot;232&quot;, <em>target</em> = 8
<strong>输出: </strong>[&quot;2*3+2&quot;, &quot;2+3*2&quot;]</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入:</strong> <code><em>num</em> = </code>&quot;105&quot;, <em>target</em> = 5
<strong>输出: </strong>[&quot;1*0+5&quot;,&quot;10-5&quot;]</pre>

<p><strong>示例&nbsp;4:</strong></p>

<pre><strong>输入:</strong> <code><em>num</em> = </code>&quot;00&quot;, <em>target</em> = 0
<strong>输出: </strong>[&quot;0+0&quot;, &quot;0-0&quot;, &quot;0*0&quot;]
</pre>

<p><strong>示例 5:</strong></p>

<pre><strong>输入:</strong> <code><em>num</em> = </code>&quot;3456237490&quot;, <em>target</em> = 9191
<strong>输出: </strong>[]
</pre>


 #### 题解
 回溯
 ```go
func addOperators(num string, target int) (ans []string) {
	var dfs func(s,currStr string,curr,prefix int)
	dfs = func(s, currStr string, curr, prefix int) {
		if len(s) == 0 {
			if curr == target {
				ans = append(ans, currStr)
			}
			return
		}
		for i := 1; i <= len(s); i++ {
			left := s[:i]
			if left[0] == '0' && len(left) > 1 { // 以0开头的字符串无法变成数字
				return
			}
			leftNum := integer(left)
			right := s[i:]
			if len(currStr) == 0 {
				dfs(right,left,leftNum,leftNum)
			} else {
				dfs(right,currStr+"+"+left,curr+leftNum,leftNum)
				dfs(right,currStr+"-"+left,curr-leftNum,-leftNum)
				dfs(right,currStr+"*"+left,curr-prefix+prefix*leftNum,prefix*leftNum)
			}
		}
	}
	dfs(num,"",0,0)
	return
}

func integer(s string) int {
	res := int(s[0] - '0')
	for i := 1; i < len(s); i++ {
		res = res * 10 + int(s[i]-'0')
	}
	return res
}
```