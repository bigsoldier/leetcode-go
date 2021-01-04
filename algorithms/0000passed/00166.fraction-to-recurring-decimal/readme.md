#### 题目
<p>给定两个整数，分别表示分数的分子&nbsp;numerator 和分母 denominator，以字符串形式返回小数。</p>

<p>如果小数部分为循环小数，则将循环的部分括在括号内。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> numerator = 1, denominator = 2
<strong>输出:</strong> &quot;0.5&quot;
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> numerator = 2, denominator = 1
<strong>输出:</strong> &quot;2&quot;</pre>

<p><strong>示例&nbsp;3:</strong></p>

<pre><strong>输入:</strong> numerator = 2, denominator = 3
<strong>输出: </strong>&quot;0.(6)&quot;
</pre>


 #### 题解
 按照常用的除法的过程转化成程序
 
 正负号通过位运算更合理
 ```go
func fractionToDecimal(numerator int, denominator int) string {
	sig := ""
	if numerator != 0 && (numerator&(1<<31) ^ (denominator&(1<<31)) > 0) {
		// int 默认是32位的,通过位运算计算符号大小
		// 如果通过乘积比较符号大小可能会导致越界
		sig = "-"
	}
	numerator,denominator = abs(numerator),abs(denominator)
	integer := strconv.Itoa(numerator / denominator) // 整数部分
	numerator = numerator % denominator *10
	if numerator == 0 { // 没有小数部分
		return sig + integer
	}
	demical := "" // 小数部分
	m := map[string]int{}
	for i := 0; numerator > 0; i++ {
		key := strconv.Itoa(numerator) + "/" + strconv.Itoa(denominator)
		if v, ok := m[key]; ok { // 进入循环小数
			demical = demical[:v]+"("+demical[v:i] +")"
			break
		}
		m[key] = i
		demical += strconv.Itoa(numerator / denominator)
		numerator = numerator % denominator*10
	}
	return sig + integer + "." + demical
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
```