#### 题目
<p>给定一个大小为 <em>n </em>的数组，找到其中的多数元素。多数元素是指在数组中出现次数<strong>大于</strong>&nbsp;<code>&lfloor; n/2 &rfloor;</code>&nbsp;的元素。</p>

<p>你可以假设数组是非空的，并且给定的数组总是存在多数元素。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> [3,2,3]
<strong>输出:</strong> 3</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [2,2,1,1,1,2,2]
<strong>输出:</strong> 2
</pre>


 #### 题解
 1、哈希表
 ```go
func majorityElement(nums []int) int {
	n := len(nums)
	var m = map[int]int{}
	for _, num := range nums {
		m[num]++
		if m[num] > n/2 {
			return num
		}
	}
	return 0
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 2、排序
 ```go
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
```
 时间复杂度O(nlogn),空间复杂度O(logn)
 
 3、投票
 
 因为相同的数大于一半，所以将众数与普通的数消减，那么剩下的一定是众数
 ```go
func majorityElement(nums []int) int {
	count, result := 0,-1
	for _, num := range nums {
		if count == 0 {
			result = num
		}
		if num == result {
			count++
		} else {
			count--
		}
	}
	return result
}
```
 时间复杂度O(n),空间复杂度O(1)
 