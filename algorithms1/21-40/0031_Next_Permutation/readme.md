#### 题目
实现获取下一个排列的函数，算法需要将给定数字序列重新排列称字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须*原地*修改，只允许使用额外的常数空间。
以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 -> 1,3,2
3,2,1 -> 1,2,3
1,1,5 -> 1,5,1

#### 题解
1、暴力法
列出所有的排列可能性，然后排序，找到当前的位置，找下一个。
时间复杂度O(n!)，空间复杂度O(n)

2、一次扫描
首先先了解怎么找到下一个排列。例如给定数*23451*，会知道下一个数为*23514*。
这是因为先从右向左，找到一个升序的排列*45*，然后将*4*与右边的数进行比较，如果小则交换[184==>418]，然后将交换后的位置进行升序排列。
```go
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1

	// find: A[i]<A[j]
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 { // 不是最后一个排列
		// find: A[i]<A[k]
		for nums[i] >= nums[k] {
			k--
		}
		// swap A[i], A[k]
		nums[i], nums[k] = nums[k], nums[i]
	}

	// reverse A[j:end]
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
```
时间复杂度O(n)，空间复杂度O(1)
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-27/003101.png)