#### 题目
<p>给定一个由<strong>整数</strong>组成的<strong>非空</strong>数组所表示的非负整数，在该数的基础上加一。</p>

<p>最高位数字存放在数组的首位， 数组中每个元素只存储<strong>单个</strong>数字。</p>

<p>你可以假设除了整数 0 之外，这个整数不会以零开头。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> [1,2,3]
<strong>输出:</strong> [1,2,4]
<strong>解释:</strong> 输入数组表示数字 123。
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [4,3,2,1]
<strong>输出:</strong> [4,3,2,2]
<strong>解释:</strong> 输入数组表示数字 4321。
</pre>


 #### 题解
 1、数组转为数字
 ```go
func plusOne(digits []int) []int {
	var multi = 1
	for i := 0; i < len(digits); i++ {
		multi *= 10
	}
	var ret int
	for i := 0; i < len(digits); i++ {
		multi /= 10
		ret += digits[i] * multi
	}
	ret += 1
	var result []int
	for ret > 0 {
		result = append(result, ret%10)
		ret /= 10
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}
```
但是这种方法，数组转数字会大于int和int64的范围

2、进位加一
```go
func plusOne(digits []int) []int {
	length := len(digits)
	if length == 0 {
		return nil
	}
	digits[length-1]++

	// 处理进位
	for i := length-1; i > 0; i-- {
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10
		digits[i-1]++
	}

	// 处理首位
	if digits[0] > 9 {
		digits[0] -= 10
		digits = append([]int{1},digits...)
	}
	return digits
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-21/006601.png)
时间复杂度O(n),空间复杂度O(1)