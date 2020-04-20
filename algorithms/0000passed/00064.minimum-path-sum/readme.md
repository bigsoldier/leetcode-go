#### 题目
<p>给定一个包含非负整数的 <em>m</em>&nbsp;x&nbsp;<em>n</em>&nbsp;网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。</p>

<p><strong>说明：</strong>每次只能向下或者向右移动一步。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong>
[
&nbsp; [1,3,1],
  [1,5,1],
  [4,2,1]
]
<strong>输出:</strong> 7
<strong>解释:</strong> 因为路径 1&rarr;3&rarr;1&rarr;1&rarr;1 的总和最小。
</pre>


 #### 题解
 1、暴力法
 每条路都计算一下，时间复杂度O(2^m+n^),空间复杂度O(m+n)
 
 2、二维动态规划
 dp[i][j] = grid[i][j]+min(dp[i+1][j],dp[i][j+1])
 建立一个额外的dp数组，和原矩阵大小相同，dp[i][j]表示到右小角的最小路径权值
 这个是从右下角向左上角推的
 ```go
func minPathSum(grid [][]int) int {
	height := len(grid)
	width := len(grid[0])

	var dp = make([][]int,height)
	for i := 0; i < height; i++ {
		dp[i] = make([]int,width)
	}

	for i := height - 1; i >= 0; i-- {
		for j := width - 1; j >= 0; j-- {
			if i == height-1 && j != width-1 {
				dp[i][j] = grid[i][j] + dp[i][j+1]
			} else if j == width-1 && i != height-1 {
				dp[i][j] = grid[i][j] + dp[i+1][j]
			} else if i != height-1 && j != width-1 {
				var min int
				if dp[i+1][j] < dp[i][j+1] {
					min = dp[i+1][j]
				} else {
					min = dp[i][j+1]
				}
				dp[i][j] = grid[i][j] + min
			} else {
				dp[i][j] = grid[i][j]
			}
		}
	}
	return dp[0][0]
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-20/006401.png)
时间复杂度O(m*n),空间复杂度O(m*n)

3、一维动态规划
因为对于某个固定节点来说，只需要考虑它的下侧和右侧节点
dp[j] = grid[i][j] + min(dp[j],dp[j+1])
```go
func minPathSum(grid [][]int) int {
	height := len(grid)
	width := len(grid[0])

	var dp = make([]int,width)

	for i := height - 1; i >= 0; i-- {
		for j := width - 1; j >= 0; j-- {
			if i == height-1 && j != width-1 {
				dp[j] = grid[i][j] + dp[j+1]
			} else if j == width-1 && i != height-1 {
				dp[j] = grid[i][j] + dp[j]
			} else if i != height-1 && j != width-1 {
				var min int
				if dp[j] < dp[j+1] {
					min = dp[j]
				} else {
					min = dp[j+1]
				}
				dp[j] = grid[i][j] + min
			} else {
				dp[j] = grid[i][j]
			}
		}
	}
	return dp[0]
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-20/006402.png)
时间复杂度O(m*n),空间复杂度O(n)

4、动态规划
我们观察到，每次循环其实都只用到当前grid的值和他的右和下侧的值，可以使用原数组直接存储，不用额外的空间
grid[i][j] = grid[i][j] + min(grid[i+1][j],grid[i][j+1])
```go
func minPathSum(grid [][]int) int {
	height := len(grid)
	width := len(grid[0])
	
	for i := height - 1; i >= 0; i-- {
		for j := width - 1; j >= 0; j-- {
			if i == height-1 && j != width-1 {
				grid[i][j] = grid[i][j] + grid[i][j+1]
			} else if j == width-1 && i != height-1 {
				grid[i][j] = grid[i][j] + grid[i+1][j]
			} else if i != height-1 && j != width-1 {
				var min int
				if grid[i+1][j] < grid[i][j+1] {
					min = grid[i+1][j]
				} else {
					min = grid[i][j+1]
				}
				grid[i][j] = grid[i][j] + min
			} else {
				grid[i][j] = grid[i][j]
			}
		}
	}
	return grid[0][0]
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-20/006403.png)
时间复杂度O(m*n),空间复杂度O(1)

- 其实使用左上角到右下角的方法也可以，将循环 for i := 0;i < height;i++ 递增计算