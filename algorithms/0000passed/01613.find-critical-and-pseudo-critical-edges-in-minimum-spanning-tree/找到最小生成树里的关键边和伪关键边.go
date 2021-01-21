package code

import (
	"math"
	"sort"
)

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	for i, e := range edges {
		edges[i] = append(e, i)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2] // 比较权重
	})

	caclMST := func(uf *unionFind, ignoreID int) (mstVal int) {
		for i, e := range edges {
			if i != ignoreID && uf.merge(e[0], e[1]) {
				mstVal += e[2]
			}
		}
		if uf.count > 1 {
			return math.MaxInt64
		}
		return
	}
	// 先求最小生成树的权重
	mstVal := caclMST(newUnionFind(n), -1)

	var keyEdges, pseudoKeyEdges []int
	for i, e := range edges {
		// 去掉一条边，比较权重值，如果大于最小权重，那么就是关键边
		if caclMST(newUnionFind(n), i) > mstVal {
			keyEdges = append(keyEdges, e[3])
			continue
		}
		// 是否是伪关建边
		uf := newUnionFind(n)
		uf.merge(e[0], e[1])
		if e[2]+caclMST(uf, i) == mstVal {
			pseudoKeyEdges = append(pseudoKeyEdges, e[3])
		}
	}
	return [][]int{keyEdges, pseudoKeyEdges}
}

type unionFind struct {
	fa, rank []int
	count    int // 当前连通分量
}

func newUnionFind(n int) *unionFind {
	fa := make([]int, n)
	rank := make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	return &unionFind{fa: fa, rank: rank, count: n}
}

func (u *unionFind) find(x int) int {
	if x != u.fa[x] {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *unionFind) merge(x, y int) bool {
	fx, fy := u.find(x), u.find(y)
	if fx == fy {
		return false
	}
	if u.rank[fx] < u.rank[fy] {
		fx, fy = fy, fx
	}
	u.rank[fx] += u.rank[fy]
	u.fa[fy] = fx
	u.count--
	return true
}
