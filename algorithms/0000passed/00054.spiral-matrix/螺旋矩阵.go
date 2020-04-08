package code

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var result []int
	m, n := len(matrix), len(matrix[0])
	var up, down, left, right = 0, m - 1, 0, n - 1
	for up <= down && left <= right {
		for i := left; i <= right; i++ {
			result = append(result, matrix[up][i])
		}
		for j := up + 1; j <= down; j++ {
			result = append(result, matrix[j][right])
		}
		if up < down && left < right {
			for i := right - 1; i > left; i-- {
				result = append(result, matrix[down][i])
			}
			for j := down; j > up; j-- {
				result = append(result, matrix[j][left])
			}
		}
		up++
		down--
		left++
		right--
	}
	return result
}
