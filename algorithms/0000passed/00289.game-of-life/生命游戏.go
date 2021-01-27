package code

func gameOfLife(board [][]int) {
	rows, cols := len(board), len(board[0])
	neighbors := []int{0, 1, -1}
	// 遍历面板
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			var liveNeighbors int
			// 计数每个细胞相邻的8个位置的活细胞数量
			for w := 0; w < 3; w++ {
				for v := 0; v < 3; v++ {
					if !(neighbors[w] == 0 && neighbors[v] == 0) { // 除去本格
						r, c := i+neighbors[w], j+neighbors[v]
						if (r < rows && r >= 0) && (c < cols && c >= 0) && abs(board[r][c]) == 1 {
							liveNeighbors++
						}
					}
				}
			}
			// 活细胞少于两个或大于三个，活细胞死亡
			if (board[i][j] == 1) && (liveNeighbors < 2) || liveNeighbors > 3 {
				board[i][j] = -1
			}
			// 死细胞周围有三个活细胞，死细胞复活
			if board[i][j] == 0 && liveNeighbors == 3 {
				board[i][j] = 2
			}
		}
	}
	// 复原状态
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] > 0 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
