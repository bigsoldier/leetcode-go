#### 题目
<p>在二维平面上，我们将石头放置在一些整数坐标点上。每个坐标点上最多只能有一块石头。<br>
<br>
现在，<em>move</em> 操作将会移除与网格上的某一块石头共享一列或一行的一块石头。<br>
<br>
我们最多能执行多少次 <em>move</em> 操作？</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
<strong>输出：</strong>5
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>stones = [[0,0],[0,2],[1,1],[2,0],[2,2]]
<strong>输出：</strong>3
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>stones = [[0,0]]
<strong>输出：</strong>0
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ol>
	<li><code>1 &lt;= stones.length &lt;= 1000</code></li>
	<li><code>0 &lt;= stones[i][j] &lt; 10000</code></li>
</ol>


 #### 题解
 1、深度优先
 ```go
func removeStones(stones [][]int) int {
	n := len(stones)
	graph := make([][]int,n)
	for i, p := range stones {
		for j, q := range stones {
			if p[0] == q[0] || p[1] == q[1] {
				graph[i] = append(graph[i], j)
			}
		}
	}
	visited := make([]bool,n)
	var dfs func(v int)
	dfs = func(v int) {
		visited[v] = true
		for _, w := range graph[v] {
			if !visited[w] {
				dfs(w)
			}
		}
	}

	var cnt int
	for i, v := range visited {
		if !v {
			cnt++
			dfs(i)
		}
	}
	return n-cnt
}
```
 时间复杂度O(n^2^),空间复杂度O(n^2^)
 
 2、优化建图+深度搜索
 ```go
func removeStones(stones [][]int) int {
	n := len(stones)
	rowCol := map[int][]int{}
	for i, p := range stones {
		x,y := p[0],p[1]+10001
		rowCol[x] = append(rowCol[x], i)
		rowCol[y] = append(rowCol[y], i)
	}
	graph := make([][]int,n)
	for _, id := range rowCol {
		for i := 1; i < len(id); i++ {
			v,w := id[i-1],id[i]
			graph[v] = append(graph[v], w)
			graph[w] = append(graph[w], v)
		}
	}
	visited := make([]bool,n)
	var dfs func(x int)
	dfs = func(x int) {
		visited[x] = true
		for _, v := range graph[x] {
			if !visited[v] {
				dfs(v)
			}
		}
	}
	var cnt int
	for i, v := range visited {
		if !v {
			cnt++
			dfs(i)
		}
	}
	return n-cnt
}
```
 时间复杂度O(n),空间复杂度O(n),每个石子最多只有4条边与之相连
 
 3、并查集
 在方法一二中，我们维护的是石子，也就是行和列。
 
 实际上，我们将每个石子的行列合并
 ```go
func removeStones(stones [][]int) int {
	fa,rank := map[int]int{},map[int]int{}
	var find func(int) int
	find = func(x int) int {
		if _, has := fa[x]; !has {
			fa[x],rank[x] = x,1
		}
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(x,y int) {
		fx,fy := find(x),find(y)
		if fx == fy {
			return
		}
		if rank[x] < rank[y] {
			fx,fy = fy,fx
		}
		rank[fx] += rank[fy]
		fa[fx] = fy
	}
	for _, p := range stones {
		merge(p[0],p[1]+10001)
	}
	ans := len(stones)
	for x, fx := range fa {
		if x == fx {
			ans--
		}
	}
	return ans
}
```
 时间复杂度O(n*alpha(n)),空间复杂度O(n)