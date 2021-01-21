#### 题目
<p>给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。你需要给出所有可能的组合的结果。有效的运算符号包含 <code>+</code>,&nbsp;<code>-</code>&nbsp;以及&nbsp;<code>*</code>&nbsp;。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> <code>&quot;2-1-1&quot;</code>
<strong>输出:</strong> <code>[0, 2]</code>
<strong>解释: </strong>
((2-1)-1) = 0 
(2-(1-1)) = 2</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入: </strong><code>&quot;2*3-4*5&quot;</code>
<strong>输出:</strong> <code>[-34, -14, -10, -10, 10]</code>
<strong>解释: 
</strong>(2*(3-(4*5))) = -34 
((2*3)-(4*5)) = -14 
((2*(3-4))*5) = -10 
(2*((3-4)*5)) = -10 
(((2*3)-4)*5) = 10</pre>


 #### 题解
 1、分而治之
 - 分解：按运算符分成左右两部分
 - 计算
 - 合并
 ```go
func diffWaysToCompute(input string) (ans []int) {
	if isDigit(input) {
		t,_ := strconv.Atoi(input)
		return []int{t}
	}
	for i, ch := range input {
		if ch == '+' || ch == '-' || ch == '*' {
			left := diffWaysToCompute(input[:i])
			right := diffWaysToCompute(input[i+1:])
			for _, l := range left {
				for _, r := range right {
					ans = append(ans, operate(l,r, byte(ch)))
				}
			}
		}
	}
	return
}

func isDigit(a string) bool {
	_,err := strconv.Atoi(a)
	if err != nil {
		return false
	}
	return true
}

func operate(a, b int, opt byte) int {
	switch opt {
	case '+':
		return a+b
	case '-':
		return a-b
	default:
		return a*b
	}
}
```
 