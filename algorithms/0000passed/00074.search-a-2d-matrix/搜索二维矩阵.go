package code

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) / 2
		if target == matrix[mid/n][mid%n] {
			return true
		}
		if target < matrix[mid/n][mid%n] {
			right = mid - 1
		}
		if target > matrix[mid/n][mid%n] {
			left = mid + 1
		}
	}
	return false
}
