#### 题目
<p>给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。</p>

<p>如果数组元素个数小于 2，则返回 0。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> [3,6,9,1]
<strong>输出:</strong> 3
<strong>解释:</strong> 排序后的数组是 [1,3,6,9]<strong><em>, </em></strong>其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [10]
<strong>输出:</strong> 0
<strong>解释:</strong> 数组元素个数小于 2，因此返回 0。</pre>

<p><strong>说明:</strong></p>

<ul>
	<li>你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。</li>
	<li>请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。</li>
</ul>


 #### 题解
 排序方法是线性的只有基数排序和桶排序
 
 1、基数排序
 ```go
func maximumGap(nums []int) (ans int) {
	var max int
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	nums = sort(nums,max)
	for i := 1; i < len(nums); i++ {
		n := nums[i] - nums[i-1]
		if n > ans {
			ans = n
		}
	}
	return
}

func sort(nums []int,max int) []int {
	n := len(nums)
	buf := make([]int,n)
	for e := 1; e <= max; e = e*10 {
		cnt := [10]int{} // 0-9
		for _, v := range nums {
			digit := v / e%10 // 取个位数
			cnt[digit]++
		}
		for i := 1; i < 10; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n-1; i >= 0; i-- {
			digit := nums[i] / e%10
			buf[cnt[digit]-1] = nums[i]
			cnt[digit]--
		}
		copy(nums,buf)
	}
	return buf
}
```
 时间复杂度O(n),空间复杂度O(n)