package main

type UnionFind struct {
	fa    []int
	size  []int
	count int
}

func NewUnionFind(n int) *UnionFind {
	fa := make([]int, n)
	size := make([]int, n)
	for i := range fa {
		fa[i] = i
		size[i] = 1
	}
	return &UnionFind{fa: fa, size: size, count: n}
}

func (u *UnionFind) Union(p, q int) {
	fp, fq := u.Find(p), u.Find(q)
	if fp == fq {
		return
	}
	if fp > fq {
		u.fa[fq] = fp
		u.size[fp] += u.size[fq]
	} else {
		u.fa[fp] = fq
		u.size[fq] = u.size[fp]
	}
	u.count--
}

func (u *UnionFind) Find(x int) int {
	if x != u.fa[x] {
		u.fa[x] = u.Find(u.fa[x])
	}
	return x
}

func (u *UnionFind) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *UnionFind) Count() int {
	return u.count
}

// 被围绕的区域
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	uf := NewUnionFind(m*n + 1)
	dummy := m * n
	for i := 0; i < m; i++ {
		if board[i][0] == '0' {
			uf.Union(dummy, i*n)
		}
		if board[i][n-1] == '0' {
			uf.Union(dummy, i*(n+1)-1)
		}
	}
	for j := 0; j < n; j++ {
		if board[0][j] == '0' {
			uf.Union(dummy, j)
		}
		if board[m-1][j] == '0' {
			uf.Union(dummy, n*(m-1)+j)
		}
	}
	direct := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			// 上下左右相连
			if board[i][j] == '0' {
				for k := 0; k < 4; k++ {
					x, y := i+direct[k][0], j+direct[k][1]
					if board[i][j] == '0' {
						uf.Union(x*n+y, i*n+j)
					}
				}
			}
		}
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if !uf.Connected(dummy, i*n+j) {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	board := [][]byte{{'X'}}
	solve(board)
}
