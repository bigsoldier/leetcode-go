#### 题目
<p>给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> [1,2,3]
<strong>输出:</strong> 6
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> [1,2,3,4]
<strong>输出:</strong> 24
</pre>

<p><strong>注意:</strong></p>

<ol>
	<li>给定的整型数组长度范围是[3,10<sup>4</sup>]，数组中所有的元素范围是[-1000, 1000]。</li>
	<li>输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。</li>
</ol>


 #### 题解
 1、排序
 找最大的三个正整数或两个负数一个正数
 ```go
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return max(nums[0]*nums[1]*nums[n-1],nums[n-1]*nums[n-2]*nums[n-3])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
 时间复杂度O(nlogn),空间复杂度O(logn)
 
 2、一次遍历
 记录最大的三个数和最小的两个数
 ```go
func maximumProduct(nums []int) int {
	min1,min2 := math.MaxInt32,math.MaxInt32
	max1,max2,max3 := math.MinInt32,math.MinInt32,math.MinInt32
	for _, num := range nums {
		if num < min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}
		if num > max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num > max2 {
			max3 = max2
			max2 = num
		} else if num > max3 {
			max3 = num
		}
	}
	return max(min1*min2*max1, max1*max2*max3)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
 时间复杂度O(n),空间复杂度O(1)