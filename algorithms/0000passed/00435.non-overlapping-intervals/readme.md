#### 题目
<p>给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。</p>

<p><strong>注意:</strong></p>

<ol>
	<li>可以认为区间的终点总是大于它的起点。</li>
	<li>区间 [1,2] 和 [2,3] 的边界相互&ldquo;接触&rdquo;，但没有相互重叠。</li>
</ol>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> [ [1,2], [2,3], [3,4], [1,3] ]

<strong>输出:</strong> 1

<strong>解释:</strong> 移除 [1,3] 后，剩下的区间没有重叠。
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> [ [1,2], [1,2], [1,2] ]

<strong>输出:</strong> 2

<strong>解释:</strong> 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
</pre>

<p><strong>示例 3:</strong></p>

<pre>
<strong>输入:</strong> [ [1,2], [2,3] ]

<strong>输出:</strong> 0

<strong>解释:</strong> 你不需要移除任何区间，因为它们已经是无重叠的了。
</pre>


 #### 题解
 可以将题目转化为 求最多不重叠的区间
 
 贪心算法
 
 假设某种最优解法中，[lk,rk]是首个区间，左侧没有其他区间，右侧是若干个不重叠的区间。
 假设此时存在一个区间[lj,rj]，使得rj<rk，即区间j的右端点在k区间的左侧，
 我们将区间k替换成区间j，其与剩余右侧的区间仍然是不重叠的。
 而当区间k替换成区间j，成为另一种最优解法。
 
 我们不断寻找右端点在首个区间右端点左侧的新区间，并将首个区间替换该区间。
 那么当我们无法被替换时，**首个区间就是所有可以选择的区间右端点最小的那个**。
 如果有多个区间的右端点相同怎么办？由于我们选择首个区间，因此在左端点在何处不重要。
 
 ```go
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans,right := 1,intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= right {
			ans++
			right = intervals[i][1]
		}
	}
	return len(intervals) - ans
}
```
 时间复杂度O(nlogn),空间复杂度O(logn)