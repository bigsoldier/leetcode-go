### 题目描述

<https://github.com/betterfor/leetcode-go/tree/master/algorithms/0001_Two_Sum>

给定一个整数数组 *nums* 和一个目标值 *target* ，请你在该数组中找出和未目标值的那 **两个** 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一种答案。但是，你不能重复利用这个数组中同样的元素。

**示例：**

```
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回[0，1]
```

### 题解

- 第一种方法

```go
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return nil
}
```

两重循环遍历元素，比较获取 *target* ，这就是俗称的暴力法

![暴力破解法](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2019-12-28/0001-01.png)


$$
时间复杂度为O(n^2)，空间复杂度O(1)
$$
那么有没有什么方法简化呢？



- 第二种方法

```
func twoSum(nums []int, target int) []int {
    var m = make(map[int]int)

    for i:=0; i<len(nums);i++ {
    	res := target - nums[i]
		if _,ok:=m[res];ok{
			return []int{m[res],i}
		}
		m[nums[i]] = i
	}
    return nil
}
```

遍历列表，将元素放入map中，根据 *target* 与当前数的相减可以在map中找出另一个数。

![hash算法](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2019-12-28/0001-02.png)
$$
时间复杂度为O(n)，空间复杂度O(n)
$$
问：**会有人说数组遇到重复数字怎么办？例如[2,2,3,4,5]**

答：题目中已经声明不能使用重复元素。