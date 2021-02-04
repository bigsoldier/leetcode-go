package code

type unionFind struct {
	fa, rank []int
	count    int
}

func newUnionFind(n int) *unionFind {
	fa, rank := make([]int, n), make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	return &unionFind{fa, rank, n}
}

func (u *unionFind) find(x int) int {
	if x != u.fa[x] {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *unionFind) union(x, y int) {
	fx, fy := u.find(x), u.find(y)
	if fx == fy {
		return
	}
	if u.rank[fx] < u.rank[fy] {
		fx, fy = fy, fx
	}
	u.rank[fx] += u.rank[fy]
	u.fa[fy] = fx
	u.count--
}

func (u *unionFind) isSame(x, y int) bool {
	return u.find(x) == u.find(y)
}

func numSimilarGroups(strs []string) int {
	n := len(strs)
	uf := newUnionFind(n)
	for i, s := range strs {
		for j := i + 1; j < n; j++ {
			if !uf.isSame(i, j) && isSimilar(s, strs[j]) {
				uf.union(i, j)
			}
		}
	}
	return uf.count
}

// 判断两个字符串是否相似
func isSimilar(x, y string) bool {
	var diff int
	for i := range x {
		if x[i] != y[i] {
			diff++
			if diff > 2 {
				return false
			}
		}
	}
	return true
}
