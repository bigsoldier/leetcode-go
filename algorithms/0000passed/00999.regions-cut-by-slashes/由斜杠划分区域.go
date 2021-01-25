package code

func regionsBySlashes(grid []string) (ans int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	// grid转化为矩阵
	m := len(grid)

	cnt := 4 * m * m
	fa := make([]int, cnt)
	for i := range fa {
		fa[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if x != fa[x] {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(x, y int) {
		fx, fy := find(x), find(y)
		if fx != fy {
			fa[fx] = fy
			cnt--
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			idx := i*m + j
			// 单元格间的合并
			if i < m-1 { // 向下
				buttom := idx + m
				merge(idx*4+2, buttom*4) // 2和下的0合并
			}
			if j < m-1 { // 向右
				right := idx + 1
				merge(idx*4+1, right*4+3) // 1和右的3合并
			}
			// 单元格内的合并
			if grid[i][j] == '/' {
				merge(idx*4, idx*4+3)
				merge(idx*4+1, idx*4+2)
			} else if grid[i][j] == '\\' {
				merge(idx*4, idx*4+1)
				merge(idx*4+2, idx*4+3)
			} else {
				merge(idx*4, idx*4+1)
				merge(idx*4, idx*4+2)
				merge(idx*4, idx*4+3)
			}
		}
	}
	return cnt
}
