package code

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	lx := max(A, E)
	rx := min(C, G)
	ly := max(B, F)
	ry := min(D, H)
	area1, area2 := (C-A)*(D-B), (G-E)*(H-F)
	if lx > rx || ly > ry {
		return area1 + area2
	}
	return area1 + area2 - (rx-lx)*(ry-ly)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
