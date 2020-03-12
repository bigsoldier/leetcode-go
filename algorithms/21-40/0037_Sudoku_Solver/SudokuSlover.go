package _037_Sudoku_Solver

//Write a program to solve a Sudoku puzzle by filling the empty cells.
//
// A sudoku solution must satisfy all of the following rules:
//
//
// Each of the digits 1-9 must occur exactly once in each row.
// Each of the digits 1-9 must occur exactly once in each column.
// Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-b
//oxes of the grid.
//
//
// Empty cells are indicated by the character '.'.
//
//
//A sudoku puzzle...
//
//
//...and its solution numbers marked in red.
//
// Note:
//
//
// The given board contain only digits 1-9 and the character '.'.
// You may assume that the given Sudoku puzzle will have a single unique solutio
//n.
// The given board size is always 9x9.
//
// Related Topics Hash Table Backtracking

//leetcode submit region begin(Prohibit modification and deletion)
func solveSudoku(board [][]byte) {
	solveSudo(board, 0)
}

func solveSudo(board [][]byte, index int) bool {
	if index == 81 {
		return true
	}

	row, col := index/9, index%9
	if board[row][col] != '.' {
		return solveSudo(board, index+1)
	}

	boxI, boxJ := (row/3)*3, (col/3)*3

	isValid := func(b byte) bool { // 校验横行、纵行、9宫格的数字是否符合标准
		for i := 0; i < 9; i++ {
			if board[row][i] == b || board[i][col] == b || board[boxI+i/3][boxJ+i%3] == b {
				return false
			}
		}
		return true
	}

	for b := byte('1'); b <= '9'; b++ {
		// 这里尝试填数字
		if isValid(b) { // 如果有符合的数字，将该空赋值并校验下一个空
			board[row][col] = b
			if solveSudo(board, index+1) {
				return true
			}
		}
	}

	// 沒有找到，置'.'
	board[row][col] = '.'

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
