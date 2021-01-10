#### 题目
<p>给定一个含有&nbsp;<strong>n&nbsp;</strong>个正整数的数组和一个正整数&nbsp;<strong>s ，</strong>找出该数组中满足其和<strong> &ge; s </strong>的长度最小的连续子数组<strong>。</strong>如果不存在符合条件的连续子数组，返回 0。</p>

<p><strong>示例:&nbsp;</strong></p>

<pre><strong>输入:</strong> <code>s = 7, nums = [2,3,1,2,4,3]</code>
<strong>输出:</strong> 2
<strong>解释: </strong>子数组&nbsp;<code>[4,3]</code>&nbsp;是该条件下的长度最小的连续子数组。
</pre>

<p><strong>进阶:</strong></p>

<p>如果你已经完成了<em>O</em>(<em>n</em>) 时间复杂度的解法, 请尝试&nbsp;<em>O</em>(<em>n</em> log <em>n</em>) 时间复杂度的解法。</p>


 #### 题解
 1、暴力法
 ```go
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		var sum int
		for j := i; j < n; j++ {
			sum += nums[i]
			if sum >= s {
				ans = min(ans,j-i+1)
				break
			}
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
 时间复杂度O(n^2^),空间复杂度O(1)
 
 2、二分查找
 方法一在确定每个子数组的下标后，找到长度最小的子数组需要O(n),
 如果使用二分法，可以优化到O(logn)
 ```go
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	sums := make([]int,n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	for i := 1; i <= n; i++ {
		target := s + sums[i-1]
		bound := sort.SearchInts(sums,target)
		if bound < 0 {
			bound = -bound-1
		}
		if bound <= n {
			ans = min(ans,bound-(i-1))
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
 时间复杂度O(nlogn),空间复杂度O(n)
 
 3、双指针
 ```go
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	var sum int
	ans := math.MaxInt32
	left,right := 0,0
	for right < n {
		sum += nums[right]
		for sum >= s {
			ans = min(ans,right-left+1)
			sum-=nums[left]
			left++
		}
		right++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
 时间复杂度O(n),空间复杂度O(n)