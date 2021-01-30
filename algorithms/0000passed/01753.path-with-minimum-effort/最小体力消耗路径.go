package code

type edge struct{ x, y, diff int }

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	edges := []edge{}
	for i, row := range heights {
		for j, col := range row {
			id := i*n + j
			if i > 0 {
				edges = append(edges, edge{id - n, id, abs(col - heights[i-1][j])})
			}
			if j > 0 {
				edges = append(edges, edge{id - 1, id, abs(col - heights[i][j-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].diff < edges[j].diff
	})
	uf := NewUnionFind(m * n)
	for _, e := range edges {
		uf.union(e.x, e.y)
		if uf.isSame(0, m*n-1) {
			return e.diff
		}
	}
	return 0
}

type UnionFind struct {
	fa, rank []int
}

func NewUnionFind(n int) *UnionFind {
	fa, rank := make([]int, n), make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	return &UnionFind{fa, rank}
}

func (u *UnionFind) find(x int) int {
	if x != u.fa[x] {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *UnionFind) union(x, y int) {
	fx, fy := u.find(x), u.find(y)
	if fx == fy {
		return
	}
	if u.rank[fx] < u.rank[fy] {
		fx, fy = fy, fx
	}
	u.rank[fx] += u.rank[fy]
	u.fa[fy] = fx
}

func (u *UnionFind) isSame(x, y int) bool {
	return u.find(x) == u.find(y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
