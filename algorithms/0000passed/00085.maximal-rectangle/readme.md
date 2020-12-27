#### 题目
<p>给定一个仅包含&nbsp;0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong>
[
  [&quot;1&quot;,&quot;0&quot;,&quot;1&quot;,&quot;0&quot;,&quot;0&quot;],
  [&quot;1&quot;,&quot;0&quot;,&quot;<strong>1</strong>&quot;,&quot;<strong>1</strong>&quot;,&quot;<strong>1</strong>&quot;],
  [&quot;1&quot;,&quot;1&quot;,&quot;<strong>1</strong>&quot;,&quot;<strong>1</strong>&quot;,&quot;<strong>1</strong>&quot;],
  [&quot;1&quot;,&quot;0&quot;,&quot;0&quot;,&quot;1&quot;,&quot;0&quot;]
]
<strong>输出:</strong> 6</pre>


 #### 题解
 1、动态规划
 记录每行的连续数字，然后再遍历每列的数字，计算出最大面积
 ```go
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	var maxArea int
	var dp = make(map[int]map[int]int)
	for i := 0; i < len(matrix); i++ {
		dp[i] = make(map[int]int)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				if j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i][j-1] + 1// 记录一行连续的数量
				}
				width := dp[i][j]
				for k := i; k >= 0; k-- {
					if width > dp[k][j] { // 取边界较小的
						width = dp[k][j]
					}
					if width*(i-k+1) > maxArea {
						maxArea = width*(i-k+1)
					}
				}
			}
		}
	}
	return maxArea
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-05-23/008501.png)
 时间复杂度O(n^2^m),空间复杂度O(nm)
 
 2、动态规划
 根据每一行，遍历每列，获取每列元素的高，以及左右边界，比较面积
 ```go
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	m := len(matrix)
	n := len(matrix[0])

	var left, right, height = make([]int, n), make([]int, n), make([]int, n)
	var maxArea int
	for i := 0; i < n; i++ {
		right[i] = n
	}
	for i := 0; i < m; i++ {
		curLeft, curRight := 0, n
		// update height
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		// update left
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if left[j] < curLeft {
					left[j] = curLeft
				}
			} else {
				left[j] = 0
				curLeft = j + 1
			}
		}
		// update right
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				if right[j] > curRight {
					right[j] = curRight
				}
			} else {
				right[j] = n
				curRight = j
			}
		}
		// update area
		for j := 0; j < n; j++ {
			if maxArea < (right[j]-left[j])*height[j] {
				maxArea = (right[j] - left[j]) * height[j]
			}
		}
	}
	return maxArea
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-05-26/008502.png)
 时间复杂度O(nm),空间复杂度O(n)
 
 3、方法1的变种
 ```go
func maximalRectangle(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return 0
	}
	m,n := len(matrix),len(matrix[0])
	left := make([][]int,m)
    // 动态规划，获取每行的最大连续
	for i, row := range matrix {
		left[i] = make([]int,n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1]+1
			}
		}
	}
    // 计算面积
    // 假设某列是1,2,3,4 那么面积[1*4,2*3,3*2,4*1]，通过这样来求出最大面积
	for i, row := range matrix {
		for j, v := range row {
			if v == '0' {
				continue
			}
			width := left[i][j]
			area := width
			for k := i-1; k >= 0; k-- {
				width = min(width,left[k][j])
				area = max(area,(i-k+1)*width)
			}
			ans = max(ans,area)
		}
	}
	return
}
```
 时间复杂度O(n^2^m),空间复杂度O(nm)