#### 题目
<p>给定一个无序的整数数组，找到其中最长上升子序列的长度。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> <code>[10,9,2,5,3,7,101,18]
</code><strong>输出: </strong>4 
<strong>解释: </strong>最长的上升子序列是&nbsp;<code>[2,3,7,101]，</code>它的长度是 <code>4</code>。</pre>

<p><strong>说明:</strong></p>

<ul>
	<li>可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。</li>
	<li>你算法的时间复杂度应该为&nbsp;O(<em>n<sup>2</sup></em>) 。</li>
</ul>

<p><strong>进阶:</strong> 你能将算法的时间复杂度降低到&nbsp;O(<em>n</em> log <em>n</em>) 吗?</p>


 #### 题解
 1、动态规划
 dp[i]=max(dp[j])+1,0<=j<i且nums[j]<nums[i]
 ```go
func lengthOfLIS(nums []int) (ans int) {
	n := len(nums)
	if n == 0 {
		return 0
	}
	var dp = make([]int,n)
	dp[0] = 1
	ans = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i],dp[j]+1)
			}
		}
		ans = max(ans,dp[i])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
 时间复杂度O(n^2^),空间复杂度O(n)
 
 2、贪心+二分查找
 如果我们要使上升序列尽可能的长，那么我们需要让上升序列尽可能慢，因此我们希望每次在上升序列最后加上的那个数尽可能的小。
 
 维护一个递增序列，将元素使用二分法插入到序列中
 ```go
func lengthOfLIS(nums []int) (ans int) {
	n := len(nums)
	ans = 1
	if n == 0 {
		return 0
	}
	var dp = make([]int,n+1)
	dp[n] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > dp[n] {
			ans++
			dp[ans] = nums[i]
		} else {
			l,r,pos := 1,n,0
			for l <= r {
				mid := l + (r-l)/2
				if dp[mid] < nums[i] {
					pos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			dp[pos+1] = nums[i]
		}
	}
	return
}
```
 时间复杂度O(nlogn),空间复杂度O(n)