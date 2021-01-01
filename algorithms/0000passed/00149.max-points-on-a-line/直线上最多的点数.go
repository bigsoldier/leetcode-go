package code

func maxPoints(points [][]int) int {
	n := len(points)
	if n < 3 {
		return n
	}

	// 都在同一点
	var i int
	for ; i < n-1; i++ {
		if points[i][0] != points[i+1][0] || points[i][1] != points[i+1][1] {
			break
		}
	}
	if i == n-1 {
		return n
	}

	var maxAns int
	for i = 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				continue
			}
			var count int
			for k := 0; k < n; k++ {
				if k != i && k != j {
					if isSameLine(points[i], points[j], points[k]) {
						count++
					}
				}
			}
			if count > maxAns {
				maxAns = count
			}
		}
	}
	return maxAns + 2
}

func isSameLine(a, b, c []int) bool {
	return (b[1]-a[1])*(c[0]-b[0]) == (c[1]-b[1])*(b[0]-a[0])
}
