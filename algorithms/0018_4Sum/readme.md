#### 题目
给定一个包含 *n* 个整数的数组 *nums*，判断 *nums* 中是否存在四个元素 a，b，c和d，使得 a+b+c+d=target?找出所有满足条件且**不重复**的四元组。
**示例**
```$xslt
给定数组 nums = [-1,0,-1,0,-2,2]，target=0
满足要求的三元组集合为：
[
    [-1,0,0,1],
    [-2.-1,1,2],
    [-2,0,0,2]
]
```

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