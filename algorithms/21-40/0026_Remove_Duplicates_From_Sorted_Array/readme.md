#### 题目
给定一个排序数组，需要![原地](https://baike.baidu.com/item/%E5%8E%9F%E5%9C%B0%E7%AE%97%E6%B3%95)删除重复的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，必须修改输入数组并在使用O(1)额外空间的条件下完成。
**示例1：**
```
给定数组 nums = [1,1,2]
函数应该返回新的长度 2，并且原数组 nums 的前两个元素修改为 1,2。
不需要考虑数组中超出新元素后面的元素
```
**示例2：**
```
给定数组 nums=[0,0,1,1,1,2,2,3,3,4]
函数应返回新的长度 5，并且原数组 nums 的前五个元素被修改为0,1,2,3,4,。
```

#### 题解
双指针，放置两个指针i和j。只要nums[i]=nums[j]，就增加i，跳过重复项。
```go
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var i int
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i+1
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-18/002601.png)
时间复杂度O(n)，空间复杂度O(1)