#### 题目
假设按照升序排列的数组在预先位置的某个点上进行了旋转。
（例如，数组[0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2]）
搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1.
你可以假设数组中不存在重复的元素。
你的时间复杂度必须是O(logn)级别。
**示例1：**
```
输入：nums = [4,5,6,7,0,1,2],target = 0
输出：4
```
**示例2：**
```
输入：nums = [4,5,6,7,0,1,2],target = 3
输出：-1
```

#### 题解
想要达到时间复杂度为O(logn)，只有二分法可以使用
```go
func search(nums []int, target int) int {
	// 二分法
	left,right := 0,len(nums)-1
	for left <= right {
		mid := (left + right)/2
		if target == nums[mid] {
			return mid
		}

		// 前半部分有序:[left:mid]递增
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target <= nums[right] && target > nums[mid] { // 在后半部分有序中
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-05/003301.png)
时间复杂度O(logn)，空间复杂度O(1)