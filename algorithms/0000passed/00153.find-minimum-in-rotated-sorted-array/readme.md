#### 题目
<p>假设按照升序排序的数组在预先未知的某个点上进行了旋转。</p>

<p>( 例如，数组&nbsp;<code>[0,1,2,4,5,6,7]</code> <strong> </strong>可能变为&nbsp;<code>[4,5,6,7,0,1,2]</code>&nbsp;)。</p>

<p>请找出其中最小的元素。</p>

<p>你可以假设数组中不存在重复元素。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [3,4,5,1,2]
<strong>输出:</strong> 1</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> [4,5,6,7,0,1,2]
<strong>输出:</strong> 0</pre>


 #### 题解
 升序的数组中间
 ```go
func findMin(nums []int) int {
	min := nums[0]
	for i:=1;i<len(nums);i++ {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}
	return min
}
```
 时间复杂度O(n),空间复杂度(1)
 
 ```go
func findMin(nums []int) int {
	left,right := 0,len(nums)-1
	for left < right {
		mid := (left+right)/2
		if nums[mid] > nums[right] {
			left = mid+1
		} else {
			right = mid
		}
	}
	return nums[left]
}
```
 时间复杂度O(logn),空间复杂度(1)