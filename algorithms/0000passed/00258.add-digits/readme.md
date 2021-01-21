#### 题目
<p>给定一个非负整数 <code>num</code>，反复将各个位上的数字相加，直到结果为一位数。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> <code>38</code>
<strong>输出:</strong> 2 
<strong>解释: </strong>各位相加的过程为<strong>：</strong><code>3 + 8 = 11</code>, <code>1 + 1 = 2</code>。 由于&nbsp;<code>2</code> 是一位数，所以返回 2。
</pre>

<p><strong>进阶:</strong><br>
你可以不使用循环或者递归，且在 O(1) 时间复杂度内解决这个问题吗？</p>


 #### 题解
 1、递归
 ```go
func addDigits(num int) int {
	if num < 10 {
		return num
	}
	val := 0
	for num != 0 {
		val += num%10
		num/=10
	}
	return addDigits(val)
}
```
 2、迭代
 ```go
func addDigits(num int) int {
	for num >= 10 {
		val := 0
		for num != 0 {
			val += num%10
			num /= 10
		}
		num = val
	}
	return num
}
```
 3、数学方法
 假设有一个n位的十进制数，我们可以写成x=a0+a1*10+a2*100+...
 
 因为10 mod 9 =1,100mod9=1,所以xmod9
 
 现象：
  1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
  1 2 3 4 5 6 7 8 9 1  2  3  4  5  6  7  8  9  1  2 
  
  可以观察到9个为1组。
  ```go
func addDigits(num int) int {
	return (num-1)%9+1
}
```
 时间复杂度O(1),空间复杂度O(1)