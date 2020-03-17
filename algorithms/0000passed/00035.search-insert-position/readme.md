#### 题目
<p>给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。</p>

<p>你可以假设数组中无重复元素。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [1,3,5,6], 5
<strong>输出:</strong> 2
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [1,3,5,6], 2
<strong>输出:</strong> 1
</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入:</strong> [1,3,5,6], 7
<strong>输出:</strong> 4
</pre>

<p><strong>示例 4:</strong></p>

<pre><strong>输入:</strong> [1,3,5,6], 0
<strong>输出:</strong> 0
</pre>


 #### 题解
 1、一次遍历
 遍历列表中所有元素，因为列表是增序的，所有只要比target大或相等，那么就返回该位置的索引。
 ```go
 func searchInsert(nums []int, target int) int {
 	for i := 0; i < len(nums); i++ {
 		if target <= nums[i] {
 			return i
 		}
 	}
 	return len(nums)
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-05/003501.png)
 时间复杂度O(n),空间复杂度O(1)
 
 2、二分法
 ```go
 func searchInsert(nums []int, target int) int {
 	left,right := 0, len(nums)-1
 	for left <= right {
 		mid := (left + right)/2
 		if target == nums[mid] {
 			return mid
 		} else if target < nums[mid] {
 			right = mid - 1
 		} else {
 			left = mid + 1
 		}
 	}
 	return left
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-05/003502.png)
 时间复杂度O(logn),空间复杂度O(1)