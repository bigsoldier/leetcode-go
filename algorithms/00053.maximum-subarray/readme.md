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

2、