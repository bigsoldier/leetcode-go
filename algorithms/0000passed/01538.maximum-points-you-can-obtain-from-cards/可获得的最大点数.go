package code

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	windows := n - k
	var sum, total int
	for i, point := range cardPoints {
		if i < windows {
			sum += point
		}
		total += point
	}
	ans := sum
	for i := windows; i < n; i++ {
		sum += cardPoints[i] - cardPoints[i-windows]
		ans = min(ans, sum)
	}
	return total - ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
