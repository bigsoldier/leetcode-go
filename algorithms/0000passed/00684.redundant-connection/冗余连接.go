package code

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	var fa = make([]int, n+1)
	for i := 0; i < len(fa); i++ {
		fa[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if x != fa[x] {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(x, y int) bool {
		xRoot, yRoot := find(x), find(y)
		if xRoot == yRoot {
			return false
		}
		fa[xRoot] = yRoot
		return true
	}

	for _, edge := range edges {
		if !merge(edge[0], edge[1]) {
			return edge
		}
	}
	return nil
}
