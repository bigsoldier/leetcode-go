#### 题目
<p>给定一个整数数组 <code>nums</code>&nbsp;，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [2,3,-2,4]
<strong>输出:</strong> <code>6</code>
<strong>解释:</strong>&nbsp;子数组 [2,3] 有最大乘积 6。
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> [-2,0,-1]
<strong>输出:</strong> 0
<strong>解释:</strong>&nbsp;结果不能为 2, 因为 [-2,-1] 不是子数组。</pre>


 #### 题解
 动态规划
 
 因为连续的负数相乘可以得到正数，所以也需要计算最小值
 ```go
func maxProduct(nums []int) (ans int) {
	if len(nums) == 0 {
		return 0
	}
	var dpMax,dpMin = make([]int,len(nums)),make([]int,len(nums))
	dpMax[0],dpMin[0] = nums[0],nums[0]
	for i := 1; i < len(nums); i++ {
		dpMax[i] = max(dpMax[i-1]*nums[i],max(nums[i],nums[i]*dpMin[i-1]))
		dpMin[i] = min(dpMin[i-1]*nums[i],min(nums[i],nums[i]*dpMax[i-1]))
	}
	ans = dpMax[0]
	for i := 1; i < len(nums); i++ {
		ans = max(dpMax[i],ans)
	}
	return
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 空间优化
 ```go
func maxProduct(nums []int) int {
	dpMax,dpMin,ans := nums[0],nums[0],nums[0]
	for i := 1; i < len(nums); i++ {
		mx,mn := dpMax,dpMin
		dpMax = max(mx*nums[i],max(nums[i],nums[i]*mn))
		dpMin = min(mn*nums[i],min(nums[i],nums[i]*mx))
		ans = max(dpMax,ans)
	}
	return ans
}
```
 时间复杂度O(n),空间复杂度O(1)