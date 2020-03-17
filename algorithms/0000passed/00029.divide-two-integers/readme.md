#### 题目
<p>给定两个整数，被除数&nbsp;<code>dividend</code>&nbsp;和除数&nbsp;<code>divisor</code>。将两数相除，要求不使用乘法、除法和 mod 运算符。</p>

<p>返回被除数&nbsp;<code>dividend</code>&nbsp;除以除数&nbsp;<code>divisor</code>&nbsp;得到的商。</p>

<p>整数除法的结果应当截去（<code>truncate</code>）其小数部分，例如：<code>truncate(8.345) = 8</code> 以及 <code>truncate(-2.7335) = -2</code></p>

<p>&nbsp;</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> dividend = 10, divisor = 3
<strong>输出:</strong> 3
<strong>解释: </strong>10/3 = truncate(3.33333..) = truncate(3) = 3</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> dividend = 7, divisor = -3
<strong>输出:</strong> -2
<strong>解释:</strong> 7/-3 = truncate(-2.33333..) = truncate(-2) = 3</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>被除数和除数均为 32 位有符号整数。</li>
	<li>除数不为&nbsp;0。</li>
	<li>假设我们的环境只能存储 32 位有符号整数，其数值范围是 [&minus;2<sup>31</sup>,&nbsp; 2<sup>31&nbsp;</sup>&minus; 1]。本题中，如果除法结果溢出，则返回 2<sup>31&nbsp;</sup>&minus; 1。</li>
</ul>


 #### 题解
 1、全都转化为正数计算
 ```go
 func divide(dividend int, divisor int) int {
 	if dividend == 0 {
 		return 0
 	}
 	if divisor == 1 {
 		return dividend
 	}
 	if divisor == -1 {
 		if dividend > math.MinInt32 {
 			return -dividend
 		}
 		return math.MaxInt32	// 是最小的那个数
 	}
 
 	signal := 1 	// 正符号
 	if (dividend > 0 && divisor < 0) || (dividend < 0 &&divisor > 0) {
 		signal = -1
 	}
 	if dividend < 0 {
 		dividend = -dividend
 	}
 	if divisor < 0 {
 		divisor = -divisor
 	}
 
 	result := div(dividend,divisor)
 	if signal > 0 {
 		if result > math.MaxInt32 {
 			return math.MaxInt32
 		}
 		return result
 	}
 	return -result
 }
 
 func div(dividend, divisor int) int {
 	if dividend < divisor {
 		return 0
 	}
 	var count = 1
 	tb := divisor
 	for tb + tb <= dividend {
 		tb += tb
 		count += count
 	}
 	return count + div(dividend - tb,divisor)
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-19/002901.png)
 
 2、全都转化为负数计算
 正数需要考虑边界问题，使用负数就不用考虑边界问题了。
 ```go
 func divide(dividend int, divisor int) int {
 	signal := -1
 	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
 		signal = 1
 	}
 	if dividend > 0 {
 		dividend = -dividend
 	}
 	if divisor > 0 {
 		divisor = -divisor
 	}
 
 	var result int
 	for dividend <= divisor {
 		var tmpDivisor = divisor
 		var tmpResult = -1
 		for dividend <= (tmpDivisor << 1) {
 			if tmpDivisor <= (math.MinInt32>>1) {
 				break
 			}
 			tmpDivisor = tmpDivisor<<1
 			tmpResult = tmpResult<<1
 		}
 		dividend -= tmpDivisor
 		result += tmpResult
 	}
 	if signal < 0 {
 		if result <= math.MinInt32 {
 			return math.MaxInt32
 		}
 		result = -result
 	}
 	return result
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-19/002902.png)