#### 题目
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
你的时间复杂度必须是 O(logn)级别。
如果数组中不存在目标值，返回[-1,-1]
**示例1：**
```
输入：nums = [5,7,7,8,8,10],target = 8
输出：[3,4]
```
**示例2：**
```
输入：nums = [5,7,7,8,8,10],target = 6
输出：[-1,-1]
```

#### 题解
时间复杂度是O(logn),所以还是要用二分法。
```
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