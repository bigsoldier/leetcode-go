#### 题目
<p>给定一个数组，将数组中的元素向右移动&nbsp;<em>k&nbsp;</em>个位置，其中&nbsp;<em>k&nbsp;</em>是非负数。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> <code>[1,2,3,4,5,6,7]</code> 和 <em>k</em> = 3
<strong>输出:</strong> <code>[5,6,7,1,2,3,4]</code>
<strong>解释:</strong>
向右旋转 1 步: <code>[7,1,2,3,4,5,6]</code>
向右旋转 2 步: <code>[6,7,1,2,3,4,5]
</code>向右旋转 3 步: <code>[5,6,7,1,2,3,4]</code>
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> <code>[-1,-100,3,99]</code> 和 <em>k</em> = 2
<strong>输出:</strong> [3,99,-1,-100]
<strong>解释:</strong> 
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]</pre>

<p><strong>说明:</strong></p>

<ul>
	<li>尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。</li>
	<li>要求使用空间复杂度为&nbsp;O(1) 的&nbsp;<strong>原地&nbsp;</strong>算法。</li>
</ul>


 #### 题解
 1、暴力法
 每次数组旋转一个元素
 ```go
func rotate(nums []int, k int)  {
	n := len(nums)
	for ; k > 0; k-- {
		tmp := nums[n-1]
		for i := n-1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = tmp
	}
}
```
 时间复杂度O(n*k),空间复杂度O(1)
 
 2、使用额外的数组
 ```go
func rotate(nums []int, k int)  {
	n := len(nums)
	var a = make([]int,n)
	for i := 0; i < n; i++ {
		a[(i+k)%n] = nums[i]
	}
	for i := 0; i < n; i++ {
		nums[i] = a[i]
	}
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 3、原地置换
 ```go
func rotate(nums []int, k int)  {
	n := len(nums)
	k = k%n
	for start,count := 0,0; count < len(nums); start++ {
		current := start //
		prev := nums[start]
		for {
			next := (current+k)%n
			prev,nums[next] = nums[next],prev // 元素交换
			current = next
			count++
			if start == current {
				break
			}
		}
	}
}
```
 时间复杂度O(n),空间复杂度O(1)
 
 4、反转数组
 例子：
 原：1 2 3 4 5 6 7
 全：7 6 5 4 3 2 1
 前：5 6 7 4 3 2 1
 后：5 6 7 1 2 3 4 
 ```go
func rotate(nums []int, k int)  {
	n := len(nums)
	k = k%n
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i],nums[len(nums)-i-1] = nums[len(nums)-i-1],nums[i]
	}
}
```
 时间复杂度O(n),空间复杂度O(1)