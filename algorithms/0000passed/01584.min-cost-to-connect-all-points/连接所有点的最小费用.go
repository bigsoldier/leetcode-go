package code

import "sort"

func minCostConnectPoints(points [][]int) (ans int) {
	n := len(points)
	type edge struct{ v, w, dist int }
	edges := []edge{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{i, j, dist(points[i], points[j])})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	fa, rank := make([]int, n), make([]int, n)
	for i := 0; i < len(fa); i++ {
		fa[i] = i
		rank[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if x != fa[x] {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(x, y int) bool {
		fx, fy := find(x), find(y)
		if fx == fy {
			return false
		}
		if rank[fx] < rank[fy] {
			fx, fy = fy, fx
		}
		rank[fx] += rank[fy]
		fa[fy] = fx
		return true
	}
	left := n - 1
	for _, e := range edges {
		if merge(e.v, e.w) {
			ans += e.dist
			left--
			if left == 0 {
				break
			}
		}
	}
	return
}

func dist(p, q []int) int {
	return abs(p[0]-q[0]) + abs(p[1]-q[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
