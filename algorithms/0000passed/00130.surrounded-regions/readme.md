 #### 题目
<p>给定一个二维的矩阵，包含&nbsp;<code>&#39;X&#39;</code>&nbsp;和&nbsp;<code>&#39;O&#39;</code>（<strong>字母 O</strong>）。</p>

<p>找到所有被 <code>&#39;X&#39;</code> 围绕的区域，并将这些区域里所有的&nbsp;<code>&#39;O&#39;</code> 用 <code>&#39;X&#39;</code> 填充。</p>

<p><strong>示例:</strong></p>

<pre>X X X X
X O O X
X X O X
X O X X
</pre>

<p>运行你的函数后，矩阵变为：</p>

<pre>X X X X
X X X X
X X X X
X O X X
</pre>

<p><strong>解释:</strong></p>

<p>被围绕的区间不会存在于边界上，换句话说，任何边界上的&nbsp;<code>&#39;O&#39;</code>&nbsp;都不会被填充为&nbsp;<code>&#39;X&#39;</code>。 任何不在边界上，或不与边界上的&nbsp;<code>&#39;O&#39;</code>&nbsp;相连的&nbsp;<code>&#39;O&#39;</code>&nbsp;最终都会被填充为&nbsp;<code>&#39;X&#39;</code>。如果两个元素在水平或垂直方向相邻，则称它们是&ldquo;相连&rdquo;的。</p>


 #### 题解
 我们可以观察到，如果边界出现O，那么与边界O相连的O不会被填充X。那么问题就转换为寻找边界O，我们给他换一个字符A，最后将O变成X，将A变成O
 
 深度优先
 ```go
func solve(board [][]byte) {
	if len(board) <= 1 {
		return
	}
	m,n := len(board),len(board[0])
	// 查询列边界
	for i := 0; i < m; i++ {
		dfs(board,i,0)
		dfs(board,i,n-1)
	}
	// 查询行边界
	for i := 1; i < n-1; i++ {
		dfs(board,0,i)
		dfs(board,m-1,i)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x > len(board)-1 || y < 0 || y > len(board[0])-1 || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	dfs(board,x+1,y)
	dfs(board,x-1,y)
	dfs(board,x,y+1)
	dfs(board,x,y-1)
}
```
 时间复杂度O(m*n),空间复杂度O(m*n)
 
 广度优先
 ```go
func solve(board [][]byte) {
	if len(board) <= 1 {
		return
	}
	m,n := len(board),len(board[0])
	var q [][]int
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			q = append(q, []int{i,0})
		}
		if board[i][n-1] == 'O' {
			q = append(q, []int{i,n-1})
		}
	}
	for j := 0; j < n-1; j++ {
		if board[0][j] == 'O' {
			q = append(q, []int{0,j})
		}
		if board[m-1][j] == 'O' {
			q = append(q, []int{m-1,j})
		}
	}
	for len(q) > 0 {
		cell := q[0]
		q = q[1:]
		x,y := cell[0],cell[1]
		board[x][y] = 'A'
		for i := 0; i < 4; i++ {
			mx,my := x+dx[i],y+dy[i]
			if mx < 0 || my < 0 || mx >= m || my >= n || board[mx][my] != 'O' {
				continue
			}
			q = append(q, []int{mx,my})
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
```
 时间复杂度O(m*n),空间复杂度O(m*n)