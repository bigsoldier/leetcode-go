#### 题目
<p>判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> 121
<strong>输出:</strong> true
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> -121
<strong>输出:</strong> false
<strong>解释:</strong> 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入:</strong> 10
<strong>输出:</strong> false
<strong>解释:</strong> 从右向左读, 为 01 。因此它不是一个回文数。
</pre>

<p><strong>进阶:</strong></p>

<p>你能不将整数转为字符串来解决这个问题吗？</p>


 #### 题解
 1、暴力法
 翻转整个数字，与原数字比较
 ```go
 func isPalindrome(x int) bool {
 	if x < 0 {
 		return false
 	}
 
 	var temp int
 	var y = x
 	for y != 0 {
 		temp = temp *10 + y%10
 		y/=10
 	}
 	if x == temp {
 		return true
 	}
 	return false
 }
 ```
 ![]([https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-01-16/000901.png)
 时间复杂度：O(n)
 空间复杂度：O(1)
 
 2、比较首末位数字
 如果是回文数字，那么 **1**22**1**
 1221%10=1,1221/1000=1，进行比较
 ```go
 func isPalindrome(x int) bool {
 	if x < 0 {
 		return false
 	}
 
 	div := 1
 	for x/div >= 10 {
 		div *= 10
 	}
 
 	for x > 0 {
 		left := x / div
 		right := x % 10
 		if left != right {
 			return false
 		}
 		x = (x%div)/10 // 取出去除首末位的数字
 		div/=100
 	}
 	return true
 }
 ```
 
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-01-16/000902.png)
 时间复杂度：O(n)
 空间复杂度：O(1)
 
 3、翻转一半数字
 ```go
 func isPalindrome(x int) bool {
 	if x < 0 || (x%10 == 0 && x != 0) {
 		return false
 	}
 
 	var reverseNum int
 	for x > reverseNum {
 		reverseNum = reverseNum*10 + x%10
 		x /= 10
 	}
 	return x == reverseNum || x == reverseNum/10
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-01-16/000903.png)
 时间复杂度：O(n)
 空间复杂度：O(1)