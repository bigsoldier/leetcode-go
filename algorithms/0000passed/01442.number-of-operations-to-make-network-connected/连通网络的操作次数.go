package code

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	var fa = make([]int, n)
	var cnt int
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if x != fa[x] {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(x, y int) {
		fx, fy := find(x), find(y)
		if fx != fy {
			fa[fx] = fy
		}
	}
	for _, conn := range connections {
		merge(conn[0], conn[1])
	}
	for i := range fa {
		if i == fa[i] {
			cnt++
		}
	}

	return cnt - 1
}
