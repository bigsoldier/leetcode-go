package _036_Valid_Sudoku

//Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be val
//idated according to the following rules:
//
//
// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without r
//epetition.
//
//
//
//A partially filled sudoku which is valid.
//
// The Sudoku board could be partially filled, where empty cells are filled with
// the character '.'.
//
// Example 1:
//
//
//Input:
//[
//  ["5","3",".",".","7",".",".",".","."],
//  ["6",".",".","1","9","5",".",".","."],
//  [".","9","8",".",".",".",".","6","."],
//  ["8",".",".",".","6",".",".",".","3"],
//  ["4",".",".","8",".","3",".",".","1"],
//  ["7",".",".",".","2",".",".",".","6"],
//  [".","6",".",".",".",".","2","8","."],
//  [".",".",".","4","1","9",".",".","5"],
//  [".",".",".",".","8",".",".","7","9"]
//]
//Output: true
//
//
// Example 2:
//
//
//Input:
//[
//  ["8","3",".",".","7",".",".",".","."],
//  ["6",".",".","1","9","5",".",".","."],
//  [".","9","8",".",".",".",".","6","."],
//  ["8",".",".",".","6",".",".",".","3"],
//  ["4",".",".","8",".","3",".",".","1"],
//  ["7",".",".",".","2",".",".",".","6"],
//  [".","6",".",".",".",".","2","8","."],
//  [".",".",".","4","1","9",".",".","5"],
//  [".",".",".",".","8",".",".","7","9"]
//]
//Output: false
//Explanation: Same as Example 1, except with the 5 in the top left corner being
//
//    modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is
//invalid.
//
//
// Note:
//
//
// A Sudoku board (partially filled) could be valid but is not necessarily solva
//ble.
// Only the filled cells need to be validated according to the mentioned rules.
//
// The given board contain only digits 1-9 and the character '.'.
// The given board size is always 9x9.
//
// Related Topics Hash Table

//leetcode submit region begin(Prohibit modification and deletion)
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !isValidSudokuRow(board, i) {
			return false
		}
		if !isValidSudokuCol(board, i) {
			return false
		}
		if !isValidSudokuBox(board, i) {
			return false
		}
	}
	return true
}

func isValidSudokuRow(board [][]byte, row int) bool {
	var rows [10]bool
	for i := 0; i < 9; i++ {
		num := convertNum(board[row][i])
		if num < 0 {
			continue
		}
		if rows[num] {
			return false
		}
		rows[num] = true
	}
	return true
}

func isValidSudokuCol(board [][]byte, col int) bool {
	var rows [10]bool
	for i := 0; i < 9; i++ {
		num := convertNum(board[i][col])
		if num < 0 {
			continue
		}
		if rows[num] {
			return false
		}
		rows[num] = true
	}
	return true
}

func isValidSudokuBox(board [][]byte, box int) bool {
	var rows [10]bool

	row := (box / 3) * 3
	col := (box % 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := convertNum(board[i+row][j+col])
			if num < 0 {
				continue
			}
			if rows[num] {
				return false
			}
			rows[num] = true
		}
	}
	return true
}

func convertNum(bt byte) int {
	if bt == '.' {
		return -1
	}
	return int(bt - '0')
}

//leetcode submit region end(Prohibit modification and deletion)
