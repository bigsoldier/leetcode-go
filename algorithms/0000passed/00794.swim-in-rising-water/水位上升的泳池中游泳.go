package code

type pair struct{ v, w int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func swimInWater(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var graph = make([]pair, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			graph[grid[i][j]] = pair{v: i, w: j}
		}
	}
	uf := NewUnionFind(m * n)
	for i, p := range graph {
		for _, d := range dirs {
			x, y := p.v+d.v, p.w+d.w
			if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] < i {
				uf.union(x*m+y, p.v*m+p.w)
			}
		}
		if uf.isSame(0, m*n-1) {
			return i
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
	return &UnionFind{fa: fa, rank: rank}
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
