#### 题目
<p>给定一个二维平面，平面上有&nbsp;<em>n&nbsp;</em>个点，求最多有多少个点在同一条直线上。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [[1,1],[2,2],[3,3]]
<strong>输出:</strong> 3
<strong>解释:</strong>
^
|
| &nbsp; &nbsp; &nbsp; &nbsp;o
| &nbsp; &nbsp; o
| &nbsp;o &nbsp;
+-------------&gt;
0 &nbsp;1 &nbsp;2 &nbsp;3  4
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
<strong>输出:</strong> 4
<strong>解释:</strong>
^
|
|  o
| &nbsp;&nbsp;&nbsp;&nbsp;o&nbsp;&nbsp;      o
| &nbsp;&nbsp;&nbsp;&nbsp;   o
| &nbsp;o &nbsp;      o
+-------------------&gt;
0 &nbsp;1 &nbsp;2 &nbsp;3 &nbsp;4 &nbsp;5 &nbsp;6</pre>


 #### 题解
 两点确定一条直线
 ```go
func maxPoints(points [][]int) int {
	n := len(points)
	if n < 3 {
		return n
	}

	// 都在同一点
	var i int
	for ; i < n-1; i++ {
		if points[i][0] != points[i+1][0] || points[i][1] != points[i+1][1] {
			break
		}
	}
	if i == n-1 {
		return n
	}

	var maxAns int
	for i = 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				continue
			}
			var count int
			for k := 0; k < n; k++ {
				if k !=i &&k !=j {
					if isSameLine(points[i], points[j], points[k]) {
						count++
					}
				}
			}
			if count > maxAns {
				maxAns = count
			}
		}
	}
	return maxAns + 2
}

func isSameLine(a, b, c []int) bool {
	return (b[1]-a[1])*(c[0]-b[0]) == (c[1]-b[1])*(b[0]-a[0])
}
```
 时间复杂度O(n^3^),空间复杂度O(1)
 
 优化 isSameLine 函数，