package code

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	m := len(matrix)
	n := len(matrix[0])

	var left, right, height = make([]int, n), make([]int, n), make([]int, n)
	var maxArea int
	for i := 0; i < n; i++ {
		right[i] = n
	}
	for i := 0; i < m; i++ {
		curLeft, curRight := 0, n
		// update height
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		// update left
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if left[j] < curLeft {
					left[j] = curLeft
				}
			} else {
				left[j] = 0
				curLeft = j + 1
			}
		}
		// update right
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				if right[j] > curRight {
					right[j] = curRight
				}
			} else {
				right[j] = n
				curRight = j
			}
		}
		// update area
		for j := 0; j < n; j++ {
			if maxArea < (right[j]-left[j])*height[j] {
				maxArea = (right[j] - left[j]) * height[j]
			}
		}
	}
	return maxArea
}
