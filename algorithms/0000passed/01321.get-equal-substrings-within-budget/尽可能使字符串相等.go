package code

func equalSubstring(s string, t string, maxCost int) (ans int) {
	n := len(s)
	diff := make([]int, n)
	for i, ch := range s {
		diff[i] = abs(int(ch) - int(t[i]))
	}
	var sum, left int
	for i, d := range diff {
		sum += d
		for sum > maxCost {
			sum -= diff[left]
			left++
		}
		ans = max(ans, i-left+1)
	}
	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
