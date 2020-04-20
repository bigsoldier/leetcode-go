package code

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	height := len(obstacleGrid)
	width := len(obstacleGrid[0])

	obstacleGrid[0][0] = 1

	// 找出第一列的障碍，如果有一个有障碍，那么后续全部置0
	for i := 1; i < height; i++ {
		if obstacleGrid[i][0] == 0 && obstacleGrid[i-1][0] == 1 {
			obstacleGrid[i][0] = 1
		} else {
			obstacleGrid[i][0] = 0
		}
	}

	// 同理，找出第一行的障碍，如果有一个有障碍，那么后续全部置0
	for i := 1; i < width; i++ {
		if obstacleGrid[0][i] == 0 && obstacleGrid[0][i-1] == 1 {
			obstacleGrid[0][i] = 1
		} else {
			obstacleGrid[0][i] = 0
		}
	}

	for i := 1; i < height; i++ {
		for j := 0; j < width; j++ {
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}

	return obstacleGrid[height-1][width-1]
}
