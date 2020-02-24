#### 题目
给定一个数组 nums 和一个值 val，需要 **原地** 移除所有值等于val的元素，返回移除数组后的新长度。
不需要使用额外的数组空间，你必须原地修改输入数组并使用O(1)额外空间下完成。
元素顺序可以改变，不用考虑数组中超出新长度后面的元素。
**示例1：**
```
给定 nums=[3,2,2,3]，val=3
函数应返回新的长度2，并且nums中的前两个元素均为2
```

**示例2：**
```
给定nums=[0,1,2,2,3,0,4,2]，val=2
函数应返回新的长度5，并且nums中的前五个元素为0,1,3,0,4。
```

#### 题解
```go
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	var i int
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-18/002701.png)
时间复杂度O(n)，空间复杂度O(1)