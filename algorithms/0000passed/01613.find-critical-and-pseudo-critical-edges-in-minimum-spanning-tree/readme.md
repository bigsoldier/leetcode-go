#### 题目
<p>给你一个 <code>n</code>&nbsp;个点的带权无向连通图，节点编号为 <code>0</code>&nbsp;到 <code>n-1</code>&nbsp;，同时还有一个数组 <code>edges</code>&nbsp;，其中 <code>edges[i] = [from</code><code><sub>i</sub>, to<sub>i</sub>, weight<sub>i</sub>]</code>&nbsp;表示在&nbsp;<code>from<sub>i</sub></code>&nbsp;和&nbsp;<code>to<sub>i</sub></code>&nbsp;节点之间有一条带权无向边。最小生成树&nbsp;(MST) 是给定图中边的一个子集，它连接了所有节点且没有环，而且这些边的权值和最小。</p>

<p>请你找到给定图中最小生成树的所有关键边和伪关键边。如果从图中删去某条边，会导致最小生成树的权值和增加，那么我们就说它是一条关键边。伪关键边则是可能会出现在某些最小生成树中但不会出现在所有最小生成树中的边。</p>

<p>请注意，你可以分别以任意顺序返回关键边的下标和伪关键边的下标。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/06/21/ex1.png" style="height: 262px; width: 259px;"></p>

<pre><strong>输入：</strong>n = 5, edges = [[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],[1,4,6]]
<strong>输出：</strong>[[0,1],[2,3,4,5]]
<strong>解释：</strong>上图描述了给定图。
下图是所有的最小生成树。
<img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/06/21/msts.png" style="height: 553px; width: 540px;">
注意到第 0 条边和第 1 条边出现在了所有最小生成树中，所以它们是关键边，我们将这两个下标作为输出的第一个列表。
边 2，3，4 和 5 是所有 MST 的剩余边，所以它们是伪关键边。我们将它们作为输出的第二个列表。
</pre>

<p><strong>示例 2 ：</strong></p>

<p><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/06/21/ex2.png" style="height: 253px; width: 247px;"></p>

<pre><strong>输入：</strong>n = 4, edges = [[0,1,1],[1,2,1],[2,3,1],[0,3,1]]
<strong>输出：</strong>[[],[0,1,2,3]]
<strong>解释：</strong>可以观察到 4 条边都有相同的权值，任选它们中的 3 条可以形成一棵 MST 。所以 4 条边都是伪关键边。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n &lt;= 100</code></li>
	<li><code>1 &lt;= edges.length &lt;= min(200, n * (n - 1) / 2)</code></li>
	<li><code>edges[i].length == 3</code></li>
	<li><code>0 &lt;= from<sub>i</sub> &lt; to<sub>i</sub> &lt; n</code></li>
	<li><code>1 &lt;= weight<sub>i</sub>&nbsp;&lt;= 1000</code></li>
	<li>所有 <code>(from<sub>i</sub>, to<sub>i</sub>)</code>&nbsp;数对都是互不相同的。</li>
</ul>


 #### 题解
 最小生成树对应的解法 Kruskal算法
 1、枚举+最小生成树判定
 通过不断生成最小生成树来找到关建边和伪关建边。
 ```go
func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	for i, e := range edges {
		edges[i] = append(e,i)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2] // 比较权重
	})

	caclMST := func(uf *unionFind,ignoreID int) (mstVal int) {
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
	mstVal := caclMST(newUnionFind(n),-1)

	var keyEdges,pseudoKeyEdges []int
	for i, e := range edges {
		// 去掉一条边，比较权重值，如果大于最小权重，那么就是关键边
		if caclMST(newUnionFind(n), i) > mstVal {
			keyEdges = append(keyEdges, e[3])
			continue
		}
		// 是否是伪关建边
		uf := newUnionFind(n)
		uf.merge(e[0],e[1])
		if e[2]+caclMST(uf, i) == mstVal {
			pseudoKeyEdges = append(pseudoKeyEdges, e[3])
		}
	}
	return [][]int{keyEdges,pseudoKeyEdges}
}

type unionFind struct {
	fa,rank []int
	count int // 当前连通分量
}

func newUnionFind(n int) *unionFind {
	fa := make([]int,n)
	rank := make([]int,n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	return &unionFind{fa: fa,rank: rank,count: n}
}

func (u *unionFind) find(x int) int {
	if x != u.fa[x] {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *unionFind) merge(x, y int) bool {
	fx,fy := u.find(x),u.find(y)
	if fx == fy {
		return false
	}
	if u.rank[fx] < u.rank[fy] {
		fx,fy = fy,fx
	}
	u.rank[fx] += u.rank[fy]
	u.fa[fy] = fx
	u.count--
	return true
}
```
 时间复杂度O(m^2^*alpha(n)),n和m是图中节点和边数，空间复杂度O(m+n)