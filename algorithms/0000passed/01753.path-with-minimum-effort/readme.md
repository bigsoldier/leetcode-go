#### 题目
<p>你准备参加一场远足活动。给你一个二维 <code>rows x columns</code> 的地图 <code>heights</code> ，其中 <code>heights[row][col]</code> 表示格子 <code>(row, col)</code> 的高度。一开始你在最左上角的格子 <code>(0, 0)</code> ，且你希望去最右下角的格子 <code>(rows-1, columns-1)</code> （注意下标从 <strong>0</strong> 开始编号）。你每次可以往 <strong>上</strong>，<strong>下</strong>，<strong>左</strong>，<strong>右</strong> 四个方向之一移动，你想要找到耗费 <strong>体力</strong> 最小的一条路径。</p>

<p>一条路径耗费的 <strong>体力值</strong> 是路径上相邻格子之间 <strong>高度差绝对值</strong> 的 <strong>最大值</strong> 决定的。</p>

<p>请你返回从左上角走到右下角的最小<strong> 体力消耗值</strong> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/10/25/ex1.png" style="width: 300px; height: 300px;" /></p>

<pre>
<b>输入：</b>heights = [[1,2,2],[3,8,2],[5,3,5]]
<b>输出：</b>2
<b>解释：</b>路径 [1,3,5,3,5] 连续格子的差值绝对值最大为 2 。
这条路径比路径 [1,2,2,2,5] 更优，因为另一条路径差值最大值为 3 。
</pre>

<p><strong>示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/10/25/ex2.png" style="width: 300px; height: 300px;" /></p>

<pre>
<b>输入：</b>heights = [[1,2,3],[3,8,4],[5,3,5]]
<b>输出：</b>1
<b>解释：</b>路径 [1,2,3,4,5] 的相邻格子差值绝对值最大为 1 ，比路径 [1,3,5,3,5] 更优。
</pre>

<p><strong>示例 3：</strong></p>
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/10/25/ex3.png" style="width: 300px; height: 300px;" />
<pre>
<b>输入：</b>heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
<b>输出：</b>0
<b>解释：</b>上图所示路径不需要消耗任何体力。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>rows == heights.length</code></li>
	<li><code>columns == heights[i].length</code></li>
	<li><code>1 <= rows, columns <= 100</code></li>
	<li><code>1 <= heights[i][j] <= 10<sup>6</sup></code></li>
</ul>


 #### 题解
 1、二分查找
 ```go
type pair struct{ x,y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumEffortPath(heights [][]int) int {
	m,n := len(heights),len(heights[0])
	return sort.Search(1e6, func(threshold int) bool {
		visited := make([][]bool,m)
		for i := range visited {
			visited[i] = make([]bool,n)
		}
		visited[0][0] = true
		queue := []pair{{}}
		for len(queue) > 0 {
			q := queue[0]
			queue = queue[1:]
			if q.x == m-1 && q.y == n-1 {
				return true
			}
			for _, d := range dirs {
				x,y := q.x+d.x,q.y+d.y
				if x >= 0 && x < m && y >= 0 && y < n &&
					!visited[x][y] && abs(heights[x][y]-heights[q.x][q.y]) <= threshold {
					visited[x][y] = true
					queue = append(queue, pair{x,y})
				}
			}
		}
		return false
	})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
 时间复杂度O(nmlogC),空间复杂度O(nm)
 
 2、并查集
 ```go
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
```
 时间复杂度O(nmlogmn),空间复杂度O(nm)
 
 3、Dijkstra算法
 ```go
type edge struct{ x, y, diff int }

type pair struct {
	x, y int
}

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	edges := make([][]int, m)
	for i := range edges {
		edges[i] = make([]int, n)
		for j := range edges[i] {
			edges[i][j] = math.MaxInt64
		}
	}

	edges[0][0] = 0
	h := &hp{{}}
	for {
		p := heap.Pop(h).(edge)
		if p.x == m-1 && p.y == n-1 {
			return p.diff
		}
		if edges[p.x][p.y] < p.diff {
			continue
		}
		for _, d := range dirs {
			x, y := p.x+d.x, p.y+d.y
			if x >= 0 && x < m && y >= 0 && y < n {
				if diff := max(p.diff, abs(heights[x][y]-heights[p.x][p.y])); diff < edges[x][y] {
					edges[x][y] = diff
					heap.Push(h, edge{x, y, diff})
				}
			}
		}
	}
}

type hp []edge

func (h hp) Len() int {
	return len(h)
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hp) Less(i, j int) bool {
	return h[i].diff < h[j].diff
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(edge))
}

func (h *hp) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```
 时间复杂度O(nmlogmn),空间复杂度O(nm)