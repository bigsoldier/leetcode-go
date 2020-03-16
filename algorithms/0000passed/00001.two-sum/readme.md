#### 题目
<p>给定一个整数数组 <code>nums</code>&nbsp;和一个目标值 <code>target</code>，请你在该数组中找出和为目标值的那&nbsp;<strong>两个</strong>&nbsp;整数，并返回他们的数组下标。</p>

<p>你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。</p>

<p><strong>示例:</strong></p>

<pre>给定 nums = [2, 7, 11, 15], target = 9

因为 nums[<strong>0</strong>] + nums[<strong>1</strong>] = 2 + 7 = 9
所以返回 [<strong>0, 1</strong>]
</pre>


#### 题解
- 第一种方法：暴力破解法

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
两重循环遍历元素，比较获取 *target* 。
![暴力破解法](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2019-12-28/0001-01.png)
时间复杂度为O(n^2^)，空间复杂度O(1)

- 第二种方法:一次遍历法
遍历列表，将元素放入map中，根据 *target* 与当前数的相减可以在map中找出另一个数。
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

![hash算法](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2019-12-28/0001-02.png)
时间复杂度为O(n)，空间复杂度O(n)