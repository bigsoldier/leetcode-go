#### 题目
<p>实现&nbsp;<code>int sqrt(int x)</code>&nbsp;函数。</p>
<p>计算并返回&nbsp;<em>x</em>&nbsp;的平方根，其中&nbsp;<em>x </em>是非负整数。</p>
<p>由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。</p>
<p><strong>示例 1:</strong></p>
<pre><strong>输入:</strong> 4
<strong>输出:</strong> 2
</pre>

<p><strong>示例 2:</strong></p>
<pre><strong>输入:</strong> 8
<strong>输出:</strong> 2
<strong>说明:</strong> 8 的平方根是 2.82842..., 
&nbsp;    由于返回类型是整数，小数部分将被舍去。
</pre>


 #### 题解
 1、乘法
 从1开始，累加相乘知道找到数字
 ```go
func mySqrt(x int) int {
	for i := 1; i <= x; i++ {
		if i*i > x {
			return i - 1
		} else if i * i == x {
			return i
		}
	}
	return 0
}
 ```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-22/006901.png)
时间复杂度O(n),空间复杂度O(1)

 2、![袖珍计算器算法](https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Exponential_identity)
$$
\sqrt{x} = e^\frac{\ln x}{2}
$$
```go
func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left := int(math.Pow(math.E,math.Log(float64(x))*0.5))
	right := left+1
	if right * right > x {
		return left
	}
	return right
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-22/006902.png)
时间复杂度O(1),空间复杂度O(1)

3、二叉查找
左边界和有边界逼近
```go
func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left,right := 2,x/2
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid > x {
			right = mid - 1
		} else if mid*mid < x {
			left = mid + 1
		} else {
			return mid
		}
	}
	return right
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-22/006903.png)
时间复杂度O(logn),空间复杂度O(1)

4、牛顿法
x~k+1~ = 1/2*[x~k~ + x/x~k~]
如果 x~0~ = x，逐步收敛到 sqrt(x)
```go
func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	x0 := float64(x)
	x1 := (x0+float64(x)/x0)/2.0
	for math.Abs(float64(x0-x1))>= 1 {
		x0 = x1
		x1 = (x0+float64(x)/x0)/2.0
	}
	return int(x1)
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-22/006904.png)
时间复杂度O(logn),空间复杂度O(1)