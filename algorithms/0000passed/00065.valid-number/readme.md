#### 题目
<p>验证给定的字符串是否可以解释为十进制数字。</p>

<p>例如:</p>

<p><code>&quot;0&quot;</code>&nbsp;=&gt;&nbsp;<code>true</code><br>
<code>&quot; 0.1 &quot;</code>&nbsp;=&gt;&nbsp;<code>true</code><br>
<code>&quot;abc&quot;</code>&nbsp;=&gt;&nbsp;<code>false</code><br>
<code>&quot;1 a&quot;</code>&nbsp;=&gt;&nbsp;<code>false</code><br>
<code>&quot;2e10&quot;</code>&nbsp;=&gt;&nbsp;<code>true</code><br>
<code>&quot; -90e3&nbsp; &nbsp;&quot;</code>&nbsp;=&gt;&nbsp;<code>true</code><br>
<code>&quot; 1e&quot;</code>&nbsp;=&gt;&nbsp;<code>false</code><br>
<code>&quot;e3&quot;</code>&nbsp;=&gt;&nbsp;<code>false</code><br>
<code>&quot; 6e-1&quot;</code>&nbsp;=&gt;&nbsp;<code>true</code><br>
<code>&quot; 99e2.5&nbsp;&quot;</code>&nbsp;=&gt;&nbsp;<code>false</code><br>
<code>&quot;53.5e93&quot;</code>&nbsp;=&gt;&nbsp;<code>true</code><br>
<code>&quot; --6 &quot;</code>&nbsp;=&gt;&nbsp;<code>false</code><br>
<code>&quot;-+3&quot;</code>&nbsp;=&gt;&nbsp;<code>false</code><br>
<code>&quot;95a54e53&quot;</code>&nbsp;=&gt;&nbsp;<code>false</code></p>

<p><strong>说明:</strong>&nbsp;我们有意将问题陈述地比较模糊。在实现代码之前，你应当事先思考所有可能的情况。这里给出一份可能存在于有效十进制数字中的字符列表：</p>

<ul>
	<li>数字 0-9</li>
	<li>指数 - &quot;e&quot;</li>
	<li>正/负号 - &quot;+&quot;/&quot;-&quot;</li>
	<li>小数点 - &quot;.&quot;</li>
</ul>

<p>当然，在输入中，这些字符的上下文也很重要。</p>

<p><strong>更新于 2015-02-10:</strong><br>
<code>C++</code>函数的形式已经更新了。如果你仍然看见你的函数接收&nbsp;<code>const char *</code> 类型的参数，请点击重载按钮重置你的代码。</p>


 #### 题解
 一位位判断是否存在数字,'e','.'之外的字符
 ```go
func isNumber(s string) bool {
	s = trim(s) // 去除首尾的空格
	if len(s) == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		return isUnsignedNumber(s[1:], false)
	}
	return isUnsignedNumber(s, false)
}

// 是否是无符号实数，最多只能有一个逗号
func isUnsignedNumber(s string, hasDot bool) bool {
	if len(s) == 0 {
		return false
	}

	for i, c := range s {
		switch {
		case '0' <= c && c <= '9':
			continue
		case c == '.':
			if hasDot {
				return false
			}
			if i == len(s)-1 && i != 0 {
				// 以'.'结尾
				return true
			}
			if i+1 < len(s) && s[i+1] == 'e' {
				// 2.e3是正确的
				return i != 0 && isInterger(s[i+2:])
			}
			return isUnsignedNumber(s[i+1:], true)
		case c == 'e':
			if i == 0 {
				// e不能开头
				return false
			}
			return isInterger(s[i+1:])
		default:
			return false
		}
	}
	return true
}

func isInterger(s string) bool {
	if len(s) == 0 {
		return false
	}
	// 去符号
	if s[0] == '+' || s[0] == '-' {
		return isUnsignedInterger(s[1:])
	}
	return isUnsignedInterger(s)
}

func isUnsignedInterger(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		switch {
		case c < '0' || c > '9':
			return false
		}
	}
	return true
}

func trim(s string) string {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	j := len(s) - 1
	for i <= j && s[j] == ' ' {
		j--
	}

	return s[i : j+1]
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-21/006501.png)
时间复杂度O(n),空间复杂度O(1)