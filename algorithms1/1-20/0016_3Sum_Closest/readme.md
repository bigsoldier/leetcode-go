#### 题目
给定一个包括 n 个整数的数组 nums 和一个目标值 target。找出nums中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
```$xslt
例如，给定数组 nums =[-1,2,1,-4] 和 target = 1
与 target 最接近的三个数之和为2.（-1+1+2=2）
```

#### 题解
这题与15题的题型解法都类似，就多一个取绝对值的过程。
```go
func threeSumClosest(nums []int, target int) int {
	// 排序
	sort.Ints(nums)

	var gap = math.MaxInt32
	var result int
	for i := 0; i < len(nums)-2 ; i++ {
		left := i+1
		right := len(nums) -1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > target {
				if sum - target < gap {
					gap = sum - target
					result = sum
				}
				right--
			} else if sum == target {
				return target
			} else {
				if target - sum < gap {
					gap = target - sum
					result = sum
				}
				left++
			}
		}
	}
	return result
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-13/001601.png)
时间复杂度O(n^2^)，空间复杂度O(1)