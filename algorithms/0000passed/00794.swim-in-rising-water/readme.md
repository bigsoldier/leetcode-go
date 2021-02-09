#### 题目
<p>在一个 N x N 的坐标方格&nbsp;<code>grid</code> 中，每一个方格的值 <code>grid[i][j]</code> 表示在位置 <code>(i,j)</code> 的平台高度。</p>

<p>现在开始下雨了。当时间为&nbsp;<code>t</code>&nbsp;时，此时雨水导致水池中任意位置的水位为&nbsp;<code>t</code>&nbsp;。你可以从一个平台游向四周相邻的任意一个平台，但是前提是此时水位必须同时淹没这两个平台。假定你可以瞬间移动无限距离，也就是默认在方格内部游动是不耗时的。当然，在你游泳的时候你必须待在坐标方格里面。</p>

<p>你从坐标方格的左上平台 (0，0) 出发。最少耗时多久你才能到达坐标方格的右下平台&nbsp;<code>(N-1, N-1)</code>？</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [[0,2],[1,3]]
<strong>输出:</strong> 3
<strong>解释:</strong>
时间为0时，你位于坐标方格的位置为 <code>(0, 0)。</code>
此时你不能游向任意方向，因为四个相邻方向平台的高度都大于当前时间为 0 时的水位。

等时间到达 3 时，你才可以游向平台 (1, 1). 因为此时的水位是 3，坐标方格中的平台没有比水位 3 更高的，所以你可以游向坐标方格中的任意位置
</pre>

<p><strong>示例2:</strong></p>

<pre><strong>输入:</strong> [[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]
<strong>输出:</strong> 16
<strong>解释:</strong>
<strong> 0  1  2  3  4</strong>
24 23 22 21  <strong>5</strong>
<strong>12 13 14 15 16</strong>
<strong>11</strong> 17 18 19 20
<strong>10  9  8  7  6</strong>

最终的路线用加粗进行了标记。
我们必须等到时间为 16，此时才能保证平台 (0, 0) 和 (4, 4) 是连通的
</pre>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ol>
	<li><code>2 &lt;= N &lt;= 50</code>.</li>
	<li>grid[i][j] 位于区间 [0, ..., N*N - 1] 内。</li>
</ol>


 #### 题解
 1、并查集
 这里不需要排序的原因是按照高度来
 ```go
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
```
 时间复杂度O(n^2^logn),空间复杂度O(n^2^)
 
 2、堆+广度优先
 ```go
type pair struct{ x,y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func swimInWater(grid [][]int) (ans int) {
	n := len(grid)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true
	h := &hp{{0,0,grid[0][0]}}
	for  {
		e := heap.Pop(h).(entry)
		ans = max(ans,e.val)
		if e.i == n-1 && e.j == n-1 {
			return 
		}
		for _, dir := range dirs {
			x,y := e.i+dir.x,e.j+dir.y
			if x >= 0 && x < n && y >= 0 && y < n && !visited[x][y] {
				visited[x][y] = true
				heap.Push(h,entry{x,y,grid[x][y]})
			}
		}
	}
}

type entry struct{ i, j, val int }
type hp []entry

func (h hp) Len() int            { return len(h) }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h hp) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(entry)) }
func (h *hp) Pop() interface{} {
	p := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return p
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
 时间复杂度O(n^2^logn),空间复杂度O(n^2^)
 
 3、二分查找
 ```go
type pair struct{ x,y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func swimInWater(grid [][]int) (ans int) {
	n := len(grid)
	return sort.Search(n*n-1, func(threshold int) bool {
		if threshold < grid[0][0] {
			return false
		}
		visited := make([][]bool,n)
		for i := range visited {
			visited[i] = make([]bool,n)
		}
		visited[0][0] = true
		queue := []pair{{}}
		for len(queue) > 0 {
			q := queue[0]
			queue = queue[1:]
			for _, dir := range dirs {
				x,y := q.x+dir.x,q.y+dir.y
				if x >= 0 && x < n && y >= 0 && y < n && !visited[x][y] && grid[x][y] <= threshold {
					visited[x][y] = true
					queue = append(queue, pair{x,y})
				}
			}
		}
		return visited[n-1][n-1]
	})
}
```
 时间复杂度O(n^2^logn),空间复杂度O(n^2^)