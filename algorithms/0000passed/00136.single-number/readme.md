#### 题目
<p>给定一个<strong>非空</strong>整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。</p>

<p><strong>说明：</strong></p>

<p>你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [2,2,1]
<strong>输出:</strong> 1
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [4,1,2,1,2]
<strong>输出:</strong> 4</pre>


 #### 题解
 1、暴力法
 两次循环，时间复杂度O(n^2^),空间复杂度O(1)
 
 2、快排
 ```go
func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i+=2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}
```
 时间复杂度O(nlogn),空间复杂度O(1)
 
 3、哈希表
 ```go
func singleNumber(nums []int) int {
	var m = make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for i, count := range m {
		if count == 1 {
			return i
		}
	}
	return 0
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 4、位运算
 ```go
func singleNumber(nums []int) int {
	var ret int
	for _, num := range nums {
		ret ^= num
	}
	return ret
}
```
 时间复杂度O(n),空间复杂度O(1)