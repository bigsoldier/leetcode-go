package code

func removeStones(stones [][]int) int {
	fa, rank := map[int]int{}, map[int]int{}
	var find func(int) int
	find = func(x int) int {
		if _, has := fa[x]; !has {
			fa[x], rank[x] = x, 1
		}
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(x, y int) {
		fx, fy := find(x), find(y)
		if fx == fy {
			return
		}
		if rank[x] < rank[y] {
			fx, fy = fy, fx
		}
		rank[fx] += rank[fy]
		fa[fx] = fy
	}
	for _, p := range stones {
		merge(p[0], p[1]+10001)
	}
	ans := len(stones)
	for x, fx := range fa {
		if x == fx {
			ans--
		}
	}
	return ans
}
