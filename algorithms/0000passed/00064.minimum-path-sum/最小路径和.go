package code

func minPathSum(grid [][]int) int {
	height := len(grid)
	width := len(grid[0])

	for i := height - 1; i >= 0; i-- {
		for j := width - 1; j >= 0; j-- {
			if i == height-1 && j != width-1 {
				grid[i][j] = grid[i][j] + grid[i][j+1]
			} else if j == width-1 && i != height-1 {
				grid[i][j] = grid[i][j] + grid[i+1][j]
			} else if i != height-1 && j != width-1 {
				var min int
				if grid[i+1][j] < grid[i][j+1] {
					min = grid[i+1][j]
				} else {
					min = grid[i][j+1]
				}
				grid[i][j] = grid[i][j] + min
			} else {
				grid[i][j] = grid[i][j]
			}
		}
	}
	return grid[0][0]
}
