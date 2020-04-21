#### 题目
<p>给定两个二进制字符串，返回他们的和（用二进制表示）。</p>

<p>输入为<strong>非空</strong>字符串且只包含数字&nbsp;<code>1</code>&nbsp;和&nbsp;<code>0</code>。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> a = &quot;11&quot;, b = &quot;1&quot;
<strong>输出:</strong> &quot;100&quot;</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> a = &quot;1010&quot;, b = &quot;1011&quot;
<strong>输出:</strong> &quot;10101&quot;</pre>


 #### 题解
 1、内置函数法
 使用标准库string转int或int64，然后累加后再转string
 ```go
func addBinary(a string, b string) string {
	aa,_ := strconv.Atoi(a)
	bb,_ := strconv.Atoi(b)
	return strconv.Itoa(aa+bb)
}
```
 但是有个问题，如果string传入超过32位或64位，不能转为int或int64
 
 2、逐位运算
 从a，b的末位开始累加，逢2进1
 ```go
func addBinary(a string, b string) string {
	i,j := len(a)-1,len(b)-1
	c := byte(0)
	var result string
	for i >= 0 || j >= 0 || c > 0 {
		sum := c
		if i >= 0 {
			sum += a[i] - '0'
			i--
		}
		if j >= 0 {
			sum += b[j] - '0'
			j--
		}
		c = sum/2
		result = string('0'+sum%2) + result
	}
	return result
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-21/006701.png)
 时间复杂度O(max(M,N)),空间复杂度O(max(M,N))
 
 3、位运算
 将a，b转成整型数字x和y，x保存结果，y保存进位
 使用二进制方法，利用与或，x = x^y,y = (x&y)<<1,但是golang转二进制比较麻烦，此方法省略
 