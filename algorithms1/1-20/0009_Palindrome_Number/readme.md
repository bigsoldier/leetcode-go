#### 题目
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

**示例1：**
```$xslt
输入：121
输出：true
```
**示例2：**
```$xslt
输入：-121
输出：false
解释：从左向右读为121-，不是回文数字
```
**示例3：**
```$xslt
输入：10
输出：false
解释：从左向右读是01，不是回文数字
```

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