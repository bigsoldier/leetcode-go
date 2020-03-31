package code

func totalNQueens(n int) int {
	col, diagonal1, diagonal2, row, result := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1), []int{}, 0
	putQueen(n, 0, &col, &diagonal1, &diagonal2, &row, &result)
	return result
}

func putQueen(n, index int, col, dia1, dia2 *[]bool, row *[]int, result *int) {
	if index == n {
		*result++
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
