package code

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	var head bool

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 { // 如果首列有数据，则首列置0
			head = true
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if head {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
