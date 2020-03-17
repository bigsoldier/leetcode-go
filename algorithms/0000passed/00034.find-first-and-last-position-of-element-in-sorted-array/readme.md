#### 题目
<p>给定一个按照升序排列的整数数组 <code>nums</code>，和一个目标值 <code>target</code>。找出给定目标值在数组中的开始位置和结束位置。</p>

<p>你的算法时间复杂度必须是&nbsp;<em>O</em>(log <em>n</em>) 级别。</p>

<p>如果数组中不存在目标值，返回&nbsp;<code>[-1, -1]</code>。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> nums = [<code>5,7,7,8,8,10]</code>, target = 8
<strong>输出:</strong> [3,4]</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> nums = [<code>5,7,7,8,8,10]</code>, target = 6
<strong>输出:</strong> [-1,-1]</pre>


 #### 题解
 时间复杂度是O(logn),所以还是要用二分法。
 ```go
 func searchRange(nums []int, target int) []int {
 	var index = []int{-1,-1}
 	var leftIndex = insertIndex(nums,target,true)
 	if leftIndex == len(nums) || nums[leftIndex] != target {
 		return index
 	}
 	index[0] = leftIndex
 	index[1] = insertIndex(nums,target,false)-1
 	return index
 }
 
 // 遍历列表，haw为true时，找左边界，haw为false时，找右边界
 func insertIndex(nums []int, target int, haw bool) int {
 	left,right := 0,len(nums)
 	for left < right {
 		mid := (left + right) / 2
 		if target < nums[mid] || (haw && target == nums[mid]) {
 			right = mid
 		} else {
 			left = mid + 1
 		}
 	}
 	return left
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-05/003401.png)
 时间复杂度O(logn)，空间复杂度O(1)