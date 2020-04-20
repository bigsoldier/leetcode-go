#### 题目
<p>给出一个<em>无重叠的 ，</em>按照区间起始端点排序的区间列表。</p>

<p>在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> intervals = [[1,3],[6,9]], newInterval = [2,5]
<strong>输出:</strong> [[1,5],[6,9]]
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> intervals = <code>[[1,2],[3,5],[6,7],[8,10],[12,16]]</code>, newInterval = <code>[4,8]</code>
<strong>输出:</strong> [[1,2],[3,10],[12,16]]
<strong>解释:</strong> 这是因为新的区间 <code>[4,8]</code> 与 <code>[3,5],[6,7],[8,10]</code>&nbsp;重叠。
</pre>


 #### 题解
 比较区间和新区间的边界。
 分三种情况：
 1、新区间的左边界大于区间的右边界
 2、新区间的右边界小于区间的左边界
 3、新区间的边界和区间的边界混合，这种需要分类讨论
 ```go
func insert(intervals [][]int, newInterval []int) [][]int {
	var ret [][]int

	var curIndex int
	for curIndex < len(intervals) && intervals[curIndex][1] < newInterval[0] {
		ret = append(ret, intervals[curIndex])
		curIndex ++
	}
	for curIndex < len(intervals) && intervals[curIndex][0] <= newInterval[1] {
		if intervals[curIndex][0] < newInterval[0] {
			newInterval[0] = intervals[curIndex][0]
		}
		if intervals[curIndex][1] > newInterval[1] {
			newInterval[1] = intervals[curIndex][1]
		}
		curIndex ++
	}
	ret = append(ret, newInterval)
	for curIndex < len(intervals) {
		ret = append(ret, intervals[curIndex])
		curIndex++
	}

	return ret
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-16/005701.png)
时间复杂度O(1)(一次遍历),空间复杂度O(1)