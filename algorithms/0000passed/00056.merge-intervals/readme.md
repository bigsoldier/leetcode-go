#### 题目
<p>给出一个区间的集合，请合并所有重叠的区间。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [[1,3],[2,6],[8,10],[15,18]]
<strong>输出:</strong> [[1,6],[8,10],[15,18]]
<strong>解释:</strong> 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [[1,4],[4,5]]
<strong>输出:</strong> [[1,5]]
<strong>解释:</strong> 区间 [1,4] 和 [4,5] 可被视为重叠区间。</pre>


 #### 题解
 这道题其实我们小学就会做，合并区间，画一个一维向量图就能解出来。
 那么这里，先对所有的区间首位进行排序，比较下一个区间首位和当前区间的末尾数字大小。
 ```go
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int,0,len(intervals))

	temp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= temp[1] {
			if temp[1] < intervals[i][1] {
				temp[1] =  intervals[i][1]
			}
		} else {
			result = append(result, temp)
			temp = intervals[i]
		}
	}
	result = append(result, temp)
	return result
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-13/005601.png)
时间复杂度O(nlogn),空间复杂度O(1)