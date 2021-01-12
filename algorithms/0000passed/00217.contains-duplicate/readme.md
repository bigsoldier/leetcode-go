#### 题目
<p>给定一个整数数组，判断是否存在重复元素。</p>

<p>如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [1,2,3,1]
<strong>输出:</strong> true</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入: </strong>[1,2,3,4]
<strong>输出:</strong> false</pre>

<p><strong>示例&nbsp;3:</strong></p>

<pre><strong>输入: </strong>[1,1,1,3,3,4,3,2,4,2]
<strong>输出:</strong> true</pre>


 #### 题解
 1、排序
 使用快速排序将数组按序排列，然后比较相邻元素是否相等
 ```go
func containsDuplicate(nums []int) bool {
	quickSort(nums,0,len(nums))
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func quickSort(nums []int, left, right int) {
	if left < right {
		mid := partition(nums,left,right)
		quickSort(nums,left,mid)
		quickSort(nums,mid+1,right)
	}
}

func partition(nums []int, left, right int) int {
	x,i := nums[left],left
	for j := left+1; j < right; j++ {
		if x > nums[j] {
			i++
			nums[i],nums[j] = nums[j],nums[i]
		}
	}
	nums[i],nums[left] = nums[left],nums[i]
	return i
}
```
 时间复杂度O(nlogn),空间复杂度O(logn)
 
 2、哈希表
 ```go
func containsDuplicate(nums []int) bool {
	m := map[int]struct{}{}
	for _, num := range nums {
		if _, has := m[num]; has {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}
```
 时间复杂度O(n),空间复杂度O(n)