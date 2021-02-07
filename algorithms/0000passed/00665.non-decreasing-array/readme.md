#### 题目
<p>给你一个长度为&nbsp;<code>n</code>&nbsp;的整数数组，请你判断在 <strong>最多 </strong>改变&nbsp;<code>1</code> 个元素的情况下，该数组能否变成一个非递减数列。</p>

<p>我们是这样定义一个非递减数列的：&nbsp;对于数组中所有的&nbsp;<code>i</code> <code>(0 &lt;= i &lt;= n-2)</code>，总满足 <code>nums[i] &lt;= nums[i + 1]</code>。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> nums = [4,2,3]
<strong>输出:</strong> true
<strong>解释:</strong> 你可以通过把第一个4变成1来使得它成为一个非递减数列。
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> nums = [4,2,1]
<strong>输出:</strong> false
<strong>解释:</strong> 你不能在只改变一个元素的情况下将其变为非递减数列。
</pre>

<p>&nbsp;</p>

<p><strong>说明：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 10 ^ 4</code></li>
	<li><code>- 10 ^ 5&nbsp;&lt;= nums[i] &lt;= 10 ^ 5</code></li>
</ul>


 #### 题解
 非递减序列，判断前一位与后一位的大小，如果有两次不同，那么返回false。
 
 但是这样会有特例，例如`[3,4,1,2]`，这是无法改变成非递减序列的。
 
 所以还需要在此基础上判断后一位`nums[i+1]`与前一位`[nums[i-1]`的大小，如果小于，将后一位修正为`nums[i]`，然后继续比较。
 
 ```go
func checkPossibility(nums []int) bool {
	var cnt int
	for i := 0; i < len(nums)-1; i++ {
		x,y := nums[i],nums[i+1]
		if x > y {
			cnt++
			if cnt > 1 {
				return false
			}
			if i > 0 && y < nums[i-1] {
				nums[i+1] = x
			}
		}
	}
	return true
}
```
 时间复杂度O(n),空间复杂度O(1)