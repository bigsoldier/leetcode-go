package code

func solveNQueens(n int) [][]string {
	col, diagonal1, diagonal2, row, result := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1), []int{}, [][]string{}
	putQueen(n, 0, &col, &diagonal1, &diagonal2, &row, &result)
	return result
}

func putQueen(n, index int, col, dia1, dia2 *[]bool, row *[]int, result *[][]string) {
	if index == n {
		*result = append(*result, generateBoard(n, row))
		return
	}
	for i := 0; i < n; i++ {
		if !(*col)[i] && !(*dia1)[index+i] && !(*dia2)[index-i+n-1] {
			*row = append(*row, i)
			(*col)[i] = true
			(*dia1)[index+i] = true
			(*dia2)[index-i+n-1] = true
			putQueen(n, index+1, col, dia1, dia2, row, result)
			(*col)[i] = false
			(*dia1)[index+i] = false
			(*dia2)[index-i+n-1] = false
			*row = (*row)[:len(*row)-1]
		}
	}
	return
}

func generateBoard(n int, row *[]int) []string {
	var board []string
	var result string
	for i := 0; i < n; i++ {
		result += "."
	}
	for i := 0; i < n; i++ {
		board = append(board, result)
	}
	for i := 0; i < n; i++ {
		tmp := []byte(board[i])
		tmp[(*row)[i]] = 'Q'
		board[i] = string(tmp)
	}
	return board
}
