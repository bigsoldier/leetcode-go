#### 题目
<p>给你一个长度为&nbsp;<em>n</em>&nbsp;的整数数组&nbsp;<code>nums</code>，其中&nbsp;<em>n</em> &gt; 1，返回输出数组&nbsp;<code>output</code>&nbsp;，其中 <code>output[i]</code>&nbsp;等于&nbsp;<code>nums</code>&nbsp;中除&nbsp;<code>nums[i]</code>&nbsp;之外其余各元素的乘积。</p>

<p>&nbsp;</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> <code>[1,2,3,4]</code>
<strong>输出:</strong> <code>[24,12,8,6]</code></pre>

<p>&nbsp;</p>

<p><strong>提示：</strong>题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。</p>

<p><strong>说明: </strong>请<strong>不要使用除法，</strong>且在&nbsp;O(<em>n</em>) 时间复杂度内完成此题。</p>

<p><strong>进阶：</strong><br>
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组<strong>不被视为</strong>额外空间。）</p>


 #### 题解
 这似乎很简单，先数组内所有的数的乘积，然后除以每个元素。但是这样会有问题，如果有元素为0，那么就不可行。
 同时，题目规定不允许使用除法。
 
 1、记录一个数的左边的乘积和右边的乘积
  ```go
 func productExceptSelf(nums []int) []int {
 	n := len(nums)
 	left,right,ans := make([]int,n),make([]int,n),make([]int,n)
 	left[0] = 1
 	for i := 1; i < n; i++ {
 		left[i] = left[i-1] * nums[i-1]
 	}
 	right[n-1] = 1
 	for i := n-2; i >= 0; i-- {
 		right[i] = right[i+1]*nums[i+1]
 	}
 	for i := 0; i < n; i++ {
 		ans[i] = left[i] * right[i]
 	}
 	return ans
 }
 ```
 时间复杂度O(n),空间复杂度O(1)

 2、优化空间
 ```go
func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int,n)
	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	right := 1
	for i := n-1; i >= 0; i-- {
		ans[i] = ans[i] * right
		right *= nums[i]
	}
	return ans
}
```
 时间复杂度O(n),空间复杂度O(1)