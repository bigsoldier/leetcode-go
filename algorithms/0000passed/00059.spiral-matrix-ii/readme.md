#### 题目
<p>给定一个正整数&nbsp;<em>n</em>，生成一个包含 1 到&nbsp;<em>n</em><sup>2</sup>&nbsp;所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 3
<strong>输出:</strong>
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]</pre>


 #### 题解
 与54题类似,将打印改为生成螺旋矩阵。
 ```go
func generateMatrix(n int) [][]int {
	var ret = make([][]int,n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int,n)
	}
	var num = 1
	var up,down,left,right = 0,n-1,0,n-1
	for i := 0; i < n/2+1; i++ {
		for j := left; j <= right; j++ {
			ret[up][j] = num
			num++
		}
		for j := up + 1; j <= down; j++ {
			ret[j][right] = num
			num++
		}
		for j := right - 1; j >= left; j-- {
			ret[down][j] = num
			num++
		}
		for j := down - 1; j > up; j-- {
			ret[j][left] = num
			num++
		}
		up++
		left++
		down--
		right--
	}
	return ret
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-16/005901.png)
时间复杂度O(n^2^),空间复杂度O(1)