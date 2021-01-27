package code

func maxNumEdgesToRemove(n int, edges [][]int) int {
	ans := len(edges)
	alice, bob := NewUnionFind(n), NewUnionFind(n)
	for _, e := range edges {
		x, y := e[1]-1, e[2]-1
		// 保留公共边
		if e[0] == 3 && (!alice.isSame(x, y) || !bob.isSame(x, y)) {
			alice.union(x, y)
			bob.union(x, y)
			ans--
		}
	}
	uf := [2]*UnionFind{alice, bob}
	for _, e := range edges {
		if tp := e[0]; tp < 3 && uf[tp-1].union(e[1]-1, e[2]-1) {
			ans--
		}
	}
	if alice.count > 1 || bob.count > 1 { // 无法遍历所有图
		return -1
	}
	return ans
}

type UnionFind struct {
	fa, rank []int
	count    int
}

func NewUnionFind(n int) *UnionFind {
	fa := make([]int, n)
	rank := make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	return &UnionFind{fa, rank, n}
}

func (u *UnionFind) find(x int) int {
	if x != u.fa[x] {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *UnionFind) union(x, y int) bool {
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

func (u *UnionFind) isSame(x, y int) bool {
	return u.find(x) == u.find(y)
}
