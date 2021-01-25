#### 题目
<p>将非负整数转换为其对应的英文表示。可以保证给定输入小于&nbsp;2<sup>31</sup> - 1 。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> 123
<strong>输出:</strong> &quot;One Hundred Twenty Three&quot;
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> 12345
<strong>输出:</strong> &quot;Twelve Thousand Three Hundred Forty Five&quot;</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入:</strong> 1234567
<strong>输出:</strong> &quot;One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven&quot;</pre>

<p><strong>示例 4:</strong></p>

<pre><strong>输入:</strong> 1234567891
<strong>输出:</strong> &quot;One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One&quot;</pre>


 #### 题解
 英文的数字都是3位加逗号","
 
 ```go
var thousand = []string{"","Thousand","Million","Billion"} //
var ten = []string{"","","Twenty","Thirty","Forty","Fifty","Sixty","Seventy","Eighty","Ninety"}
var digit = []string{"","One","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten",
	"Eleven","Twelve","Thirteen","Fourteen","Fifteen","Sixteen","Seventeen","Eighteen","Nineteen","Twenty"}

func numberToWords(num int) (ans string) {
	if num == 0 {
		return "Zero"
	}
	var i int
	for num > 0 {
		res := num%1000
		if res != 0 {
			ans = calc(res) + thousand[i] + " " + ans
		}
		num = num/1000
		i++
	}
	return strings.TrimRight(ans," ")
}

// 处理小于1000的数
func calc(num int) string {
	if num == 0 {
		return ""
	}
	if num <= 20 {
		return digit[num] + " "
	}
	if num < 100 {
		return ten[num/10] + " " + calc(num%10)
	}
	return digit[num/100] + " Hundred " + calc(num%100)
}
```