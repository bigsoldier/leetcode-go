#### 题目
<p>给定一个包含&nbsp;<em>m</em> x <em>n</em>&nbsp;个元素的矩阵（<em>m</em> 行, <em>n</em> 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong>
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
<strong>输出:</strong> [1,2,3,6,9,8,7,4,5]
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong>
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
<strong>输出:</strong> [1,2,3,4,8,12,11,10,9,5,6,7]
</pre>


 #### 题解
 按层模拟
 将一个矩阵按顺时针螺旋顺序输出，可分为4个重复的步骤，分别对应的上右下左四个方向的数据。
 1 1 1 1 1
 4       2  <== 最终取1,2,3,4对应的数据
 4       2
 4 3 3 3 2
 ```go
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var result []int
	m,n := len(matrix),len(matrix[0])
	var up,down,left,right = 0,m-1,0,n-1
	for up <= down && left <= right {
		for i := left; i <= right; i++ {
			result = append(result, matrix[up][i])
		}
		for j := up+1; j <= down; j++ {
			result = append(result, matrix[j][right])
		}
		if up < down && left < right {
			for i := right - 1; i > left; i-- {
				result = append(result, matrix[down][i])
			}
			for j := down; j > up; j-- {
				result = append(result, matrix[j][left])
			}
		}
		up++
		down--
		left++
		right--
	}
	return result
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-08/005401.png)
时间复杂度O(n^2^),空间复杂度O(1)