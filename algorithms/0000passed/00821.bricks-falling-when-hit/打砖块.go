package code

func hitBricks(grid [][]int, hits [][]int) []int {
	rows, cols := len(grid), len(grid[0])
	total := rows * cols

	var fa = make([]int, total+1)
	size := make([]int, total+1)
	for i := 0; i < len(fa); i++ {
		fa[i] = i
		size[i] = 1
	}
	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(x, y int) {
		xRoot, yRoot := find(x), find(y)
		if xRoot != yRoot {
			size[yRoot] += size[xRoot]
			fa[xRoot] = yRoot
		}
	}
	index := func(x, y int) int {
		return x*cols + y
	}
	isArea := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < rows && y < cols
	}
	// 1、将grid砖头全部击碎
	var graph = make([][]int, rows)
	for i := 0; i < rows; i++ {
		graph[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			graph[i][j] = grid[i][j]
		}
	}
	for _, hit := range hits {
		graph[hit[0]][hit[1]] = 0
	}
	// 2、建图
	// 将下标为0的这一行与屋顶相连
	for j := 0; j < cols; j++ {
		if graph[0][j] == 1 {
			merge(index(0, j), total)
		}
	}
	// 其余网格，如果是砖块，查看上、左
	for i := 1; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if graph[i][j] == 1 {
				// 如果上方是砖块
				if graph[i-1][j] == 1 {
					merge(index(i-1, j), index(i, j))
				}
				// 如果左方是砖块
				if j > 0 && graph[i][j-1] == 1 {
					merge(index(i, j-1), index(i, j))
				}
			}
		}
	}
	// 按照hits的逆序，补回砖块
	hitsLen := len(hits)
	var res = make([]int, hitsLen) // 补回砖块与屋顶相连的增量
	for i := hitsLen - 1; i >= 0; i-- {
		x, y := hits[i][0], hits[i][1]
		if grid[x][y] == 0 { // 如果原来是0，那么就不存在砖块
			continue
		}
		// 补回之前的砖块数
		origin := size[find(total)]
		// 如果是再第一行，告诉并查集它与屋顶香兰
		if x == 0 {
			merge(index(x, y), total)
		}
		// 查询4个方向
		for _, direct := range direction {
			newX, newY := x+direct[0], y+direct[1]
			if isArea(newX, newY) && graph[newX][newY] == 1 {
				merge(index(x, y), index(newX, newY))
			}
		}
		// 补回后的砖块数
		curr := size[find(total)]
		if curr-origin-1 > 0 {
			res[i] = curr - origin - 1
		}
		graph[x][y] = 1
	}
	return res
}

var direction = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
