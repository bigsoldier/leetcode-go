#### 题目
<p>给定两个以字符串形式表示的非负整数&nbsp;<code>num1</code>&nbsp;和&nbsp;<code>num2</code>，返回&nbsp;<code>num1</code>&nbsp;和&nbsp;<code>num2</code>&nbsp;的乘积，它们的乘积也表示为字符串形式。</p>
<p><strong>示例 1:</strong></p>
<pre><strong>输入:</strong> num1 = &quot;2&quot;, num2 = &quot;3&quot;
<strong>输出:</strong> &quot;6&quot;</pre>

<p><strong>示例&nbsp;2:</strong></p>
<pre><strong>输入:</strong> num1 = &quot;123&quot;, num2 = &quot;456&quot;
<strong>输出:</strong> &quot;56088&quot;</pre>

<p><strong>说明：</strong></p>
<ol>
	<li><code>num1</code>&nbsp;和&nbsp;<code>num2</code>&nbsp;的长度小于110。</li>
	<li><code>num1</code> 和&nbsp;<code>num2</code> 只包含数字&nbsp;<code>0-9</code>。</li>
	<li><code>num1</code> 和&nbsp;<code>num2</code>&nbsp;均不以零开头，除非是数字 0 本身。</li>
	<li><strong>不能使用任何标准库的大数类型（比如 BigInteger）</strong>或<strong>直接将输入转换为整数来处理</strong>。</li>
</ol>


 #### 题解
 这里用到竖排乘法。
通常乘数num1的位数为M，被乘数num2的位数为N，那么他们之积的位数不大于M+N。
例如：14 * 5， 我们会先计算4*5=20，然后再计算1*5=5，这时个位为0，十位为5+2=7.
扩展为两位数计算
    2 4
    5 6
---------
    2 4
  2 0      
​  1 2     
1 0
---------
1 3 4 4
 ```go
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var result = make([]uint8,len(num1) + len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := num1[i] - '0'
		for j := len(num2) - 1; j >= 0; j-- {
			n2 := num2[j] - '0'
			sum := result[i+j+1] + n1*n2
			result[i+j+1] = sum % 10
			result[i+j] += sum/10   // 进位
		}
	}

	var ret string
	for i := 0; i < len(result); i++ {
		if i == 0 && result[i] == 0 {
			continue
		}
		ret += string(result[i] + '0')
	}
	return ret
}
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-19/004301.png)