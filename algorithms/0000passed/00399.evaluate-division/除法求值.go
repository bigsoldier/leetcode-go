package code

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给变量编号
	id := map[string]int{}
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, ok := id[a]; !ok {
			id[a] = len(id)
		}
		if _, ok := id[b]; !ok {
			id[b] = len(id)
		}
	}

	fa := make([]int, len(id))    // 父亲
	w := make([]float64, len(id)) // 权重
	for i := range fa {
		fa[i] = i
		w[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			f := find(fa[x])
			w[x] *= w[fa[x]]
			fa[x] = f
		}
		return fa[x]
	}
	merge := func(from, to int, val float64) {
		fFrom, fTo := find(from), find(to)
		w[fFrom] = val * w[to] / w[from]
		fa[fFrom] = fTo
	}

	for i, eq := range equations {
		merge(id[eq[0]], id[eq[1]], values[i])
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, ok1 := id[q[0]]
		end, ok2 := id[q[1]]
		if ok1 && ok2 && find(start) == find(end) {
			ans[i] = w[start] / w[end]
		} else {
			ans[i] = -1
		}
	}
	return ans
}
