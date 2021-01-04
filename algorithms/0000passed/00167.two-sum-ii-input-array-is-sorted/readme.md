#### 题目
<p>给定一个已按照<strong><em>升序排列</em>&nbsp;</strong>的有序数组，找到两个数使得它们相加之和等于目标数。</p>

<p>函数应该返回这两个下标值<em> </em>index1 和 index2，其中 index1&nbsp;必须小于&nbsp;index2<em>。</em></p>

<p><strong>说明:</strong></p>

<ul>
	<li>返回的下标值（index1 和 index2）不是从零开始的。</li>
	<li>你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。</li>
</ul>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> numbers = [2, 7, 11, 15], target = 9
<strong>输出:</strong> [1,2]
<strong>解释:</strong> 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。</pre>


 #### 题解
 前面有一道同名题，用的相同的方法
 ```go
func twoSum(numbers []int, target int) []int {
	var m = map[int]int{}
	for i := 0; i < len(numbers); i++ {
		if v,ok := m[numbers[i]]; ok {
			return []int{v,i+1}
		}
		m[target - numbers[i]] = i+1
	}
	return nil
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 因为数组是升序数组，我们可以使用
 - 双指针
 ```go
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}
```
 时间复杂度O(n),空间复杂度O(1)