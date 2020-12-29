#### 题目
<p>给定一个已排序的正整数数组 <em>nums，</em>和一个正整数&nbsp;<em>n 。</em>从&nbsp;<code>[1, n]</code>&nbsp;区间内选取任意个数字补充到&nbsp;<em>nums&nbsp;</em>中，使得&nbsp;<code>[1, n]</code>&nbsp;区间内的任何数字都可以用&nbsp;<em>nums&nbsp;</em>中某几个数字的和来表示。请输出满足上述要求的最少需要补充的数字个数。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入: </strong><em>nums</em> = <code>[1,3]</code>, <em>n</em> = <code>6</code>
<strong>输出: </strong>1 
<strong>解释:</strong>
根据<em> nums&nbsp;</em>里现有的组合&nbsp;<code>[1], [3], [1,3]</code>，可以得出&nbsp;<code>1, 3, 4</code>。
现在如果我们将&nbsp;<code>2</code>&nbsp;添加到&nbsp;<em>nums 中，</em>&nbsp;组合变为: <code>[1], [2], [3], [1,3], [2,3], [1,2,3]</code>。
其和可以表示数字&nbsp;<code>1, 2, 3, 4, 5, 6</code>，能够覆盖&nbsp;<code>[1, 6]</code>&nbsp;区间里所有的数。
所以我们最少需要添加一个数字。</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入: </strong><em>nums</em> = <code>[1,5,10]</code>, <em>n</em> = <code>20</code>
<strong>输出:</strong> 2
<strong>解释: </strong>我们需要添加&nbsp;<code>[2, 4]</code>。
</pre>

<p><strong>示例&nbsp;3:</strong></p>

<pre><strong>输入: </strong><em>nums</em> = <code>[1,2,2]</code>, <em>n</em> = <code>5</code>
<strong>输出:</strong> 0
</pre>


 #### 题解
 对于正整数x，如果区间[1,x-1]内所有的数字已经被覆盖,
 且x在数组中，那么区间[1,2x-1]内所有的数字也会被覆盖。
 
 根据这个规则:
 [1,2,4,8]可以覆盖15内的所有数字;
 [1,2,4,8,16]可以覆盖21内所有的数字。
 
 那么如果给定的数组如 [1,2,5],16这样的条件，我们能得到数组[1,2,4,5,9]这样的数组
 ```go
func minPatches(nums []int, n int) (patches int) {
	for i, x := 0, 1; x <= n ; {
		if i < len(nums) && nums[i] <= x {
			x += nums[i]
			i++
		} else {
			x*=2
			patches++
		}
	}
	return
}
```
 时间复杂度O(n),空间复杂度O(1)