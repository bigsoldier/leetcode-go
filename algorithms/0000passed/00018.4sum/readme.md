#### 题目
<p>给定一个包含&nbsp;<em>n</em> 个整数的数组&nbsp;<code>nums</code>&nbsp;和一个目标值&nbsp;<code>target</code>，判断&nbsp;<code>nums</code>&nbsp;中是否存在四个元素 <em>a，</em><em>b，c</em>&nbsp;和 <em>d</em>&nbsp;，使得&nbsp;<em>a</em> + <em>b</em> + <em>c</em> + <em>d</em>&nbsp;的值与&nbsp;<code>target</code>&nbsp;相等？找出所有满足条件且不重复的四元组。</p>

<p><strong>注意：</strong></p>

<p>答案中不可以包含重复的四元组。</p>

<p><strong>示例：</strong></p>

<pre>给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
</pre>


 #### 题解
 题目和三数相加[15]类似，只是4数相加需要处理重复四元组的情况
 ```go
 func fourSum(nums []int, target int) [][]int {
 	var result [][]int
 	sort.Ints(nums)
 	for i := 0; i < len(nums)-3; i++ {
 		// 去重
 		if i > 0 && nums[i] == nums[i-1] {
 			continue
 		} else if nums[i] + nums[i+1] + nums[i+2] + nums[i+3] > target {
 			break
 		} else if nums[i]+nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3] < target {
 			continue
 		}
 
 		for j := i + 1; j < len(nums)-2; j++ {
 			// 去重
 			if j-i > 1 && nums[j] == nums[j-1] {
 				continue
 			} else if nums[i] + nums[j] + nums[j+1] + nums[j+2] > target {
 				break
 			} else if nums[i]+nums[j]+nums[len(nums)-1]+nums[len(nums)-2] < target {
 				continue
 			}
 
 			left,right := j+1,len(nums)-1
 			for left < right {
 				sum := nums[i]+nums[j]+nums[left]+nums[right]
 				if sum == target {
 					result = append(result, []int{nums[i],nums[j],nums[left],nums[right]})
 					for left < right && nums[left] == nums[left+1] {
 						left++
 					}
 					for left < right && nums[right] == nums[right-1] {
 						right--
 					}
 					left++
 					right--
 				} else if sum > target {
 					right--
 				} else {
 					left++
 				}
 			}
 		}
 	}
 	return result
 }
 
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-13/001801.png)