#### 题目
<p>给定一个整数数组，判断数组中是否有两个不同的索引 <em>i</em> 和 <em>j</em>，使得&nbsp;<strong>nums [i]</strong> 和&nbsp;<strong>nums [j]</strong>&nbsp;的差的绝对值最大为 <em>t</em>，并且 <em>i</em> 和 <em>j</em> 之间的差的绝对值最大为 <em>ķ</em>。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> nums = [1,2,3,1], k<em> </em>= 3, t = 0
<strong>输出:</strong> true</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入: </strong>nums = [1,0,1,1], k<em> </em>=<em> </em>1, t = 2
<strong>输出:</strong> true</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入: </strong>nums = [1,5,9,1,5,9], k = 2, t = 3
<strong>输出:</strong> false</pre>


 #### 题解
 1、暴力法
 ```go
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := max(i-k,0); j < i; j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
```
 时间复杂度O(n min(n,k)),空间复杂度O(1)
 
 2、桶
 两数之差小于k，那么
 
 |a-b|<k => -k<a-b<k => -1<a/k-b/k<1 
 
 由于我们计算的是整除，那么 a/k=b/k 
 
 特殊情况: a/k=0时需要判断符号，那么我们设a<0时，a/k=a/k-1
 
 那么a=-3，b=0的情况下，a/k=-1,b/k=0,所以我们需要多判断一下 |a-b|<t
 
 ```go
func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	size := len(nums)
	if k == 0 || t < 0 || size <= 1 {
		return false
	}
	t++
	nMap := map[int]int{}
	for i, num := range nums {
		m := num/t
		if num < 0 {
			m--
		}
		if _, has := nMap[m]; has {
			return true
		} else if n, ok := nMap[m-1]; ok && abs(n, num) < t {
			return true
		} else if n, ok = nMap[m+1]; ok && abs(n, num) < t {
			return true
		}
		nMap[m] = num
		if i >= k {
			delete(nMap,nums[i-k]/t)
		}
	}

	return false
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
```
 时间复杂度O(n),空间复杂度O(min(n,k))