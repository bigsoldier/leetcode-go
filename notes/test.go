package main

import "fmt"

func main() {
	board := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(board))
}

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
