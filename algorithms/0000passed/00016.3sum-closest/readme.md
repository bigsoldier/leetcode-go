#### 题目
<p>给定一个包括&nbsp;<em>n</em> 个整数的数组&nbsp;<code>nums</code><em>&nbsp;</em>和 一个目标值&nbsp;<code>target</code>。找出&nbsp;<code>nums</code><em>&nbsp;</em>中的三个整数，使得它们的和与&nbsp;<code>target</code>&nbsp;最接近。返回这三个数的和。假定每组输入只存在唯一答案。</p>

<pre>例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
</pre>


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