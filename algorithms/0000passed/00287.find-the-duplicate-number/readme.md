#### 题目
<p>给定一个包含&nbsp;<em>n</em> + 1 个整数的数组&nbsp;<em>nums</em>，其数字都在 1 到 <em>n&nbsp;</em>之间（包括 1 和 <em>n</em>），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> <code>[1,3,4,2,2]</code>
<strong>输出:</strong> 2
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> [3,1,3,4,2]
<strong>输出:</strong> 3
</pre>

<p><strong>说明：</strong></p>

<ol>
	<li><strong>不能</strong>更改原数组（假设数组是只读的）。</li>
	<li>只能使用额外的 <em>O</em>(1) 的空间。</li>
	<li>时间复杂度小于 <em>O</em>(<em>n</em><sup>2</sup>) 。</li>
	<li>数组中只有一个重复的数字，但它可能不止重复出现一次。</li>
</ol>


 #### 题解
 1、二分法
 由于只有一个数是重复的，又要求空间复杂度为O(1)，所以不能使用哈希表。
 
 假设target是重复值，那么target左边的数量是小于target值的。
 
 ```go
func findDuplicate(nums []int) (ans int) {
	n := len(nums)
	left,right := 0,n-1
	for left <= right {
		mid := left + (right-left)/2
		var cnt int
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			left = mid + 1
		} else {
			right = mid - 1
			ans = mid
		}
	}
	return
}
```
 时间复杂度O(nlogn),空间复杂度O(1)
 
 2、快慢指针
 例：[1,2,3,4,5,6,4]按照元素查找会出现[1,2,3,4,5,6,5,6..]
 那么使用快指针走两步，慢指针走一步，必定会相遇。
 ```go
func findDuplicate(nums []int) int {
	slow,fast := nums[0],nums[nums[0]]
	for slow != fast {
		slow,fast = nums[slow],nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow,fast = nums[slow],nums[fast]
	}
	return slow
}
```
 时间复杂度O(n),空间复杂度O(1)