#### 题目
<p>给定一个整数数组 <code>nums</code>&nbsp;，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> [-2,1,-3,4,-1,2,1,-5,4],
<strong>输出:</strong> 6
<strong>解释:</strong>&nbsp;连续子数组&nbsp;[4,-1,2,1] 的和最大，为&nbsp;6。
</pre>

<p><strong>进阶:</strong></p>

<p>如果你已经实现复杂度为 O(<em>n</em>) 的解法，尝试使用更为精妙的分治法求解。</p>


 #### 题解
 1、暴力法
 找出所有的子数组，然后累加和最大值进行比较
 ```go
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// 找到多个子数组
	var max = math.MinInt64
	for i := 0; i < len(nums); i ++ {
		for j := i; j < len(nums); j++ {
			var sum int
			for k := i; k <= j; k++ {
				sum+=nums[k]
			}
			if sum > max {
				max = sum
			}
		}
	}
	return max
}
```
时间复杂度O(n^3^),空间复杂度O(1),结果超时

2、贪心算法
每次记录当前元素位置的最大和
```go
func maxSubArray(nums []int) int {
	var cur,max = nums[0],nums[0]
	for i := 1; i < len(nums); i++ {
		sum := nums[i] + cur
		if nums[i] < sum {
			cur = sum
		} else {
			cur = nums[i]
		}

		if max < cur {
			max = cur
		}
	}
	return max
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-07/005302.png)
时间复杂度O(n),空间复杂度O(1)

3、动态规划
```go
func maxSubArray(nums []int) int {
	var sum = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if  sum < nums[i] {
			sum = nums[i]
		}
	}
	return sum
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-07/005303.png)
时间复杂度O(n),空间复杂度O(1)