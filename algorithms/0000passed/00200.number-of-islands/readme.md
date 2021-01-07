#### 题目
<p>给定一个由&nbsp;<code>&#39;1&#39;</code>（陆地）和 <code>&#39;0&#39;</code>（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong>
11110
11010
11000
00000

<strong>输出:</strong>&nbsp;1
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong>
11000
11000
00100
00011

<strong>输出: </strong>3
</pre>


 #### 题解
 1、深度搜索
 ```go
func numIslands(grid [][]byte) (ans int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(grid,i,j)
			}
		}
	}
	return 
}

func dfs(grid [][]byte, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	dfs(grid,x-1,y)
	dfs(grid,x+1,y)
	dfs(grid,x,y-1)
	dfs(grid,x,y+1)
}
```
 时间复杂度O(MN),空间复杂度O(MN)
 
 2、广度搜索
 ```go
func numIslands(grid [][]byte) (ans int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m,n := len(grid),len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				grid[i][j] = '0'
				queue := [][]int{{i,j}}
				for len(queue) > 0 {
					q := queue[0]
					queue = queue[1:]
					row,col := q[0],q[1]
					if row-1 >= 0 && grid[row-1][col] == '1' {
						grid[row-1][col] = '0'
						queue = append(queue, []int{row-1,col})
					}
					if row+1 < m && grid[row+1][col] == '1' {
						grid[row+1][col] = '0'
						queue = append(queue, []int{row+1,col})
					}
					if col-1 >= 0 && grid[row][col-1] == '1' {
						grid[row][col-1] = '0'
						queue = append(queue, []int{row,col-1})
					}
					if col + 1 < n && grid[row][col+1] == '1' {
						grid[row][col+1] = '0'
						queue = append(queue, []int{row,col+1})
					}
				}
			}
		}
	}
	return
}
```
 时间复杂度O(MN),空间复杂度O(min(M,N))
 
 3、并查集
 ```go
func numIslands(grid [][]byte) (ans int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	fa := map[int]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				fa[i*n+j] = i*n + j
				ans++
			}
		}
	}
	var find func(x int) int
	find = func(x int) int {
		if x != fa[x] {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(x, y int) {
		xx, yy := find(x), find(y)
		if xx != yy {
			fa[xx] = yy
			ans--
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				num := i*m + j
				if i-1 >= 0 && grid[i-1][j] == '1' {
					merge(num, (i-1)*n+j)
				}
				if i+1 < m && grid[i+1][j] == '1' {
					merge(num, (i+1)*n+j)
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					merge(num, i*n+j-1)
				}
				if j+1 < n && grid[i][j+1] == '1' {
					merge(num, i*n+j+1)
				}
			}
		}
	}
	return
}
```
 时间复杂度O(MN*a(MN)),空间复杂度O(min(M,N))