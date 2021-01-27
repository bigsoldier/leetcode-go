#### 题目
<p>Alice 和 Bob 共有一个无向图，其中包含 n 个节点和 3&nbsp; 种类型的边：</p>

<ul>
	<li>类型 1：只能由 Alice 遍历。</li>
	<li>类型 2：只能由 Bob 遍历。</li>
	<li>类型 3：Alice 和 Bob 都可以遍历。</li>
</ul>

<p>给你一个数组 <code>edges</code> ，其中 <code>edges[i] = [type<sub>i</sub>, u<sub>i</sub>, v<sub>i</sub>]</code>&nbsp;表示节点 <code>u<sub>i</sub></code> 和 <code>v<sub>i</sub></code> 之间存在类型为 <code>type<sub>i</sub></code> 的双向边。请你在保证图仍能够被 Alice和 Bob 完全遍历的前提下，找出可以删除的最大边数。如果从任何节点开始，Alice 和 Bob 都可以到达所有其他节点，则认为图是可以完全遍历的。</p>

<p>返回可以删除的最大边数，如果 Alice 和 Bob 无法完全遍历图，则返回 -1 。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/09/06/5510ex1.png" style="height: 191px; width: 179px;"></strong></p>

<pre><strong>输入：</strong>n = 4, edges = [[3,1,2],[3,2,3],[1,1,3],[1,2,4],[1,1,2],[2,3,4]]
<strong>输出：</strong>2
<strong>解释：</strong>如果删除<strong> </strong>[1,1,2] 和 [1,1,3] 这两条边，Alice 和 Bob 仍然可以完全遍历这个图。再删除任何其他的边都无法保证图可以完全遍历。所以可以删除的最大边数是 2 。
</pre>

<p><strong>示例 2：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/09/06/5510ex2.png" style="height: 190px; width: 178px;"></strong></p>

<pre><strong>输入：</strong>n = 4, edges = [[3,1,2],[3,2,3],[1,1,4],[2,1,4]]
<strong>输出：</strong>0
<strong>解释：</strong>注意，删除任何一条边都会使 Alice 和 Bob 无法完全遍历这个图。
</pre>

<p><strong>示例 3：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/09/06/5510ex3.png" style="height: 190px; width: 178px;"></strong></p>

<pre><strong>输入：</strong>n = 4, edges = [[3,2,3],[1,1,2],[2,3,4]]
<strong>输出：</strong>-1
<strong>解释：</strong>在当前图中，Alice 无法从其他节点到达节点 4 。类似地，Bob 也不能达到节点 1 。因此，图无法完全遍历。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 10^5</code></li>
	<li><code>1 &lt;= edges.length &lt;= min(10^5, 3 * n * (n-1) / 2)</code></li>
	<li><code>edges[i].length == 3</code></li>
	<li><code>1 &lt;= edges[i][0] &lt;= 3</code></li>
	<li><code>1 &lt;= edges[i][1] &lt; edges[i][2] &lt;= n</code></li>
	<li>所有元组 <code>(type<sub>i</sub>, u<sub>i</sub>, v<sub>i</sub>)</code> 互不相同</li>
</ul>


 #### 题解
 并查集
 ```go
func maxNumEdgesToRemove(n int, edges [][]int) int {
	ans := len(edges)
	alice,bob := NewUnionFind(n),NewUnionFind(n)
	for _, e := range edges {
		x,y := e[1]-1,e[2]-1
		// 保留公共边
		if e[0] == 3 && (!alice.isSame(x, y) || !bob.isSame(x, y)) {
			alice.union(x,y)
			bob.union(x,y)
			ans--
		}
	}
	uf := [2]*UnionFind{alice,bob}
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
	fa,rank []int
	count int
}

func NewUnionFind(n int) *UnionFind {
	fa := make([]int,n)
	rank := make([]int,n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	return &UnionFind{fa,rank,n}
}

func (u *UnionFind) find(x int) int {
	if x != u.fa[x] {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *UnionFind) union(x, y int) bool {
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

func (u *UnionFind) isSame(x, y int) bool {
	return u.find(x) == u.find(y)
}
```
 时间复杂度O(m*a(n)),a是阿克曼函数反函数,空间复杂度O(n)