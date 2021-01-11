package code

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	fa := make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if fa[i] != i {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}
	merge := func(x, y int) {
		fa[find(x)] = find(y)
	}

	for _, p := range pairs {
		merge(p[0], p[1])
	}

	groups := make(map[int][]int, n)
	for i, f := range fa {
		f = find(f)
		// 相连的索引值，放在一起
		// 例: [[0,2,4],[1,3,5]]
		groups[f] = append(groups[f], i)
	}

	bytes := []byte(s)
	res := make([]byte, n)

	for _, g := range groups {
		size := len(g)
		a := make([]int, size)
		copy(a, g)
		// a中的索引值，按照其在bytes中排序
		sort.Slice(a, func(i, j int) bool {
			return bytes[a[i]] < bytes[a[j]]
		})
		// g中索引值排序
		sort.Ints(g)
		for i := 0; i < size; i++ {
			res[g[i]] = bytes[a[i]]
		}
	}
	return string(res)
}
