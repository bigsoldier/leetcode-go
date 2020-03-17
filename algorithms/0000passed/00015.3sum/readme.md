#### 题目
<p>给你一个包含 <em>n</em> 个整数的数组&nbsp;<code>nums</code>，判断&nbsp;<code>nums</code>&nbsp;中是否存在三个元素 <em>a，b，c ，</em>使得&nbsp;<em>a + b + c = </em>0 ？请你找出所有满足条件且不重复的三元组。</p>

<p><strong>注意：</strong>答案中不可以包含重复的三元组。</p>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>

<pre>给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
</pre>


 #### 题解
 1、暴力法
 遍历三遍数组，很容易获取到对应的三元组，但有个问题是如何取到不重复的三元组（可以通过对三元组排序然后存入map中去重），时间复杂度为O(n^3^)，所以放弃这个方法
 
 2、双指针法
 根据定点元素的左右指针移动计算
 ```go
 func threeSum(nums []int) [][]int {
 	if len(nums) < 3 { // 如果数组小于3，返回[]
 		return nil
 	}
 
 	// 对数组排序(快排，时间复杂度O(nlogn))
 	sort.Ints(nums)
 
 	var result [][]int
 	for i := 0; i < len(nums)-2; i++ {
 		if nums[i] > 0 { // 如果当前元素大于0，后续肯定不会存在符合的三元组
 			break
 		}
 		if i > 0 && nums[i] == nums[i-1] { // 对于出现的重复解，跳过
 			continue
 		}
 
 		left,right := i+1,len(nums)-1
 		for left < right {
 			if nums[i]+nums[left]+nums[right] == 0 {
 				result = append(result, []int{nums[i],nums[left],nums[right]})
 				for left < right && nums[left] == nums[left+1] { // 去除左指针的重复元素
 					left ++
 				}
 				for left < right && nums[right] == nums[right-1] { // 去除右指针的重复元素
 					right--
 				}
 				left++
 				right--
 			} else if nums[i]+nums[left]+nums[right] > 0 { // 三数之和大于0，移动右指针
 				right--
 			} else { // 三数之和大于0，移动左指针
 				left++
 			}
 		}
 	}
 	return result
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-12/001501.png)
 时间复杂度：O(n^2^)[排序O(nlogn)，双指针遍历数组O(n^2^)],空间复杂度O(1)