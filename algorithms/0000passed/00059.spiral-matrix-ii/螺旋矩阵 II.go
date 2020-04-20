package code

func generateMatrix(n int) [][]int {
	var ret = make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	var num = 1
	var up, down, left, right = 0, n - 1, 0, n - 1
	for i := 0; i < n/2+1; i++ {
		for j := left; j <= right; j++ {
			ret[up][j] = num
			num++
		}
		for j := up + 1; j <= down; j++ {
			ret[j][right] = num
			num++
		}
		for j := right - 1; j >= left; j-- {
			ret[down][j] = num
			num++
		}
		for j := down - 1; j > up; j-- {
			ret[j][left] = num
			num++
		}
		up++
		left++
		down--
		right--
	}
	return ret
}
