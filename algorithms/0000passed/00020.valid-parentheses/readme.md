#### 题目
<p>给定一个只包括 <code>&#39;(&#39;</code>，<code>&#39;)&#39;</code>，<code>&#39;{&#39;</code>，<code>&#39;}&#39;</code>，<code>&#39;[&#39;</code>，<code>&#39;]&#39;</code>&nbsp;的字符串，判断字符串是否有效。</p>

<p>有效字符串需满足：</p>

<ol>
	<li>左括号必须用相同类型的右括号闭合。</li>
	<li>左括号必须以正确的顺序闭合。</li>
</ol>

<p>注意空字符串可被认为是有效字符串。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> &quot;()&quot;
<strong>输出:</strong> true
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> &quot;()[]{}&quot;
<strong>输出:</strong> true
</pre>

<p><strong>示例&nbsp;3:</strong></p>

<pre><strong>输入:</strong> &quot;(]&quot;
<strong>输出:</strong> false
</pre>

<p><strong>示例&nbsp;4:</strong></p>

<pre><strong>输入:</strong> &quot;([)]&quot;
<strong>输出:</strong> false
</pre>

<p><strong>示例&nbsp;5:</strong></p>

<pre><strong>输入:</strong> &quot;{[]}&quot;
<strong>输出:</strong> true</pre>


 #### 题解
 很显然，这是道数据结构中经典的括号匹配问题，用栈的数据结构处理
 ```go
 func isValid(s string) bool {
 	if len(s) == 0 {
 		return true
 	}
 
 	stack := make([]rune,0)
 	for _, v := range s {
 		if v == '(' || v == '{' || v == '[' {
 			stack = append(stack, v)
 		} else if v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' ||
 			v == '}' && len(stack) > 0 && stack[len(stack)-1] == '{' ||
 			v == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' {
 			stack = stack[:len(stack)-1]
 		} else {
 			return false
 		}
 	}
 	return len(stack) == 0
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-14/002001.png)
 时间复杂度O(n)，空间复杂度O(1)