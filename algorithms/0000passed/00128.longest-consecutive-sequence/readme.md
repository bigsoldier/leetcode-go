#### 题目
<p>给定一个未排序的整数数组，找出最长连续序列的长度。</p>

<p>要求算法的时间复杂度为&nbsp;<em>O(n)</em>。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong>&nbsp;[100, 4, 200, 1, 3, 2]
<strong>输出:</strong> 4
<strong>解释:</strong> 最长连续序列是 <code>[1, 2, 3, 4]。它的长度为 4。</code></pre>


 #### 题解
 哈希表
 
 用哈希表去重，然后x-1去试探下线，然后通过x+1去获取上线
 ```go
func longestConsecutive(nums []int) int {
	var m = map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}
	var longestStreak int
	for _, num := range nums {
		if !m[num-1] {
			cur := num
			curStreak := 1
			for m[cur+1] {
				cur++
				curStreak++
			}
			if longestStreak < curStreak {
				longestStreak = curStreak
			}
		}
	}
	return longestStreak
}
```
 时间复杂度O(n),空间复杂度O(n)