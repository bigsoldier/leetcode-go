#### 题目
<p>实现一个基本的计算器来计算一个简单的字符串表达式的值。</p>

<p>字符串表达式仅包含非负整数，<code>+</code>， <code>-</code> ，<code>*</code>，<code>/</code> 四种运算符和空格&nbsp;<code>&nbsp;</code>。 整数除法仅保留整数部分。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入: </strong>&quot;3+2*2&quot;
<strong>输出:</strong> 7
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> &quot; 3/2 &quot;
<strong>输出:</strong> 1</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入:</strong> &quot; 3+5 / 2 &quot;
<strong>输出:</strong> 5
</pre>

<p><strong>说明：</strong></p>

<ul>
	<li>你可以假设所给定的表达式都是有效的。</li>
	<li>请<strong>不要</strong>使用内置的库函数 <code>eval</code>。</li>
</ul>


 #### 题解
 栈
 
 由于乘和除具有更高的优先级，所以只需要在遇到乘除时优先计算
 ```go
func calculate(s string) (sum int) {
	stack := []int{}
	num := 0
	var lastSign byte = '+'
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num *10 + int(s[i]-'0')
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i]=='/'||i==len(s)-1{
			if lastSign == '+' {
				stack = append(stack, num)
			} else if lastSign == '-' {
				stack = append(stack, -num)
			} else if lastSign == '*' {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top*num)
			} else if lastSign == '/' {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top/num)
			}
			num = 0
			lastSign = s[i]
		}
	}
	for _, i := range stack {
		sum += i
	}
	return sum
}
```
 时间复杂度O(n),空间复杂度O(n)