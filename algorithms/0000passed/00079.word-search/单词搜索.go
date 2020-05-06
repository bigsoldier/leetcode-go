package code

var (
	marked         map[int]map[int]bool
	direction      = [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	boardM, boardN int
)

func exist(board [][]byte, word string) bool {
	boardM = len(board)
	if boardM == 0 {
		return false
	}
	boardN = len(board[0])
	marked = make(map[int]map[int]bool, boardM)
	for i := 0; i < boardM; i++ {
		marked[i] = make(map[int]bool, boardN)
	}
	for i := 0; i < boardM; i++ {
		for j := 0; j < boardN; j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

// 查找(i,j)
func dfs(board [][]byte, word string, i, j, start int) bool {
	if start == len(word)-1 {
		return board[i][j] == word[start]
	}

	if board[i][j] == word[start] {
		marked[i][j] = true
		for k := 0; k < 4; k++ {
			newX := i + direction[k][0]
			newY := j + direction[k][1]
			if inArea(newX, newY) && !marked[newX][newY] {
				if dfs(board, word, newX, newY, start+1) {
					return true
				}
			}
		}
		marked[i][j] = false
	}
	return false
}

func inArea(x, y int) bool {
	return x >= 0 && x < boardM && y >= 0 && y < boardN
}
