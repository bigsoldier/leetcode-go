package code

func solve(board [][]byte) {
	if len(board) <= 1 {
		return
	}
	m, n := len(board), len(board[0])
	// 查询列边界
	for i := 0; i < m; i++ {
		dfs(board, i, 0)
		dfs(board, i, n-1)
	}
	// 查询行边界
	for i := 1; i < n-1; i++ {
		dfs(board, 0, i)
		dfs(board, m-1, i)
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
	dfs(board, x+1, y)
	dfs(board, x-1, y)
	dfs(board, x, y+1)
	dfs(board, x, y-1)
}
