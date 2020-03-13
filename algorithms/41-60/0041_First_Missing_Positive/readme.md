#### 题目
给定一个未排序的整数数组，找出其中没有出现的最小的正整数
**示例1：**
```
输入：[1,2,0]
输出：3
```
**示例2：**
```
输入：[3,4-1,1]
输出：2
```
**示例3：**
```
输入：[7,8,9,11,12]
输出：1
```
**说明**
你的算法的时间复杂度为O(n),并且只能使用常数级别的空间

#### 题解
1、hash表(时间复杂度O(n),空间复杂度O(n))
```go
func firstMissingPositive(nums []int) int {
	var newMap = make(map[int]int)
	for _, num := range nums {
		newMap[num] = num
	}

	for i := 1; i <= len(nums); i++ {
		if _, ok := newMap[i]; !ok {
			return i
		}
	}
	return len(nums) + 1
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-13/004102.png)

2、排序后统计(时间复杂度O(nlogn),空间复杂度O(1))
```go
func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	var result = 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == result {
			result++
		}
	}
	return result
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-13/004103.png)

3、数组视为hash表
```go
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		// 这里使用for循环，是因为交换完成后该位置的新值也需要进行处理
		for nums[i]-1 >= 0 && nums[i]-1 < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i],nums[nums[i]-1] = nums[nums[i]-1],nums[i]
		}
		// 这里的时间复杂度是O(n),如果元素在一次遍历里已经被交换到他们该在的位置上，那么在执行的时候会跳过。
		// 最坏的情况是第一次里循环的时候把所有元素都交换了一遍，那么之后的循环则不会再进入里循环
	}
	for j := range nums {
		if j+1 != nums[j] {
			return j+1
		}
	}
	return len(nums)+1
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-03-13/004101.png)
时间复杂度O(n),空间复杂度O(1)