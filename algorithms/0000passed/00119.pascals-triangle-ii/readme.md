#### 题目
<p>给定一个非负索引&nbsp;<em>k</em>，其中 <em>k</em>&nbsp;&le;&nbsp;33，返回杨辉三角的第 <em>k </em>行。</p>

<p><img alt="" src="https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif"></p>

<p><small>在杨辉三角中，每个数是它左上方和右上方的数的和。</small></p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 3
<strong>输出:</strong> [1,3,3,1]
</pre>

<p><strong>进阶：</strong></p>

<p>你可以优化你的算法到 <em>O</em>(<em>k</em>) 空间复杂度吗？</p>


 #### 题解
 正向加起来会改变当前的值，那么久反向累加
 ```go
func getRow(rowIndex int) []int {
	var result = make([]int,rowIndex+1)
	result[0] = 1
	for i := 1; i < rowIndex; i++ {
		for j := i; j > 0; j-- {
			result[j] += result[j-1]
		}
	}
	return result
}
```
 时间复杂度O(n^2^),空间复杂度O(n)