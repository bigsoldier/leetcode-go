#### 题目
<p>给出方程式&nbsp;<code>A / B = k</code>, 其中&nbsp;<code>A</code> 和&nbsp;<code>B</code> 均为代表字符串的变量，&nbsp;<code>k</code> 是一个浮点型数字。根据已知方程式求解问题，并返回计算结果。如果结果不存在，则返回&nbsp;<code>-1.0</code>。</p>

<p><strong>示例 :</strong><br />
给定&nbsp;<code>a / b = 2.0, b / c = 3.0</code><br />
问题: <code> a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?&nbsp;</code><br />
返回&nbsp;<code>[6.0, 0.5, -1.0, 1.0, -1.0 ]</code></p>

<p>输入为: <code> vector&lt;pair&lt;string, string&gt;&gt; equations, vector&lt;double&gt;&amp; values, vector&lt;pair&lt;string, string&gt;&gt; queries</code>(方程式，方程式结果，问题方程式)，&nbsp;其中&nbsp;<code>equations.size() == values.size()</code>，即方程式的长度与方程式结果长度相等（程式与结果一一对应），并且结果值均为正数。以上为方程式的描述。&nbsp;返回<code>vector&lt;double&gt;</code>类型。</p>

<p>基于上述例子，输入如下：</p>

<pre>
equations(方程式) = [ [&quot;a&quot;, &quot;b&quot;], [&quot;b&quot;, &quot;c&quot;] ],
values(方程式结果) = [2.0, 3.0],
queries(问题方程式) = [ [&quot;a&quot;, &quot;c&quot;], [&quot;b&quot;, &quot;a&quot;], [&quot;a&quot;, &quot;e&quot;], [&quot;a&quot;, &quot;a&quot;], [&quot;x&quot;, &quot;x&quot;] ]. 
</pre>

<p>输入总是有效的。你可以假设除法运算中不会出现除数为0的情况，且不存在任何矛盾的结果。</p>


 #### 题解
 1、广度优先搜索
 将问题建模成一张图:给定图中的一些点(变量),以及某些边的权值(两个变量的比值),求任意两边的路径长。
 
 因此，首先遍历equations数组，找出其中所有不同的字符串，通过哈希表映射成整数。
 
 在建完图后，对于任意查询，从起点出发，通过广度搜索的方式，不断更新起点和当前点的路径长度，直至搜索到终点。
 
 ```go
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给变量编号
	id := map[string]int{}
	for _, eq := range equations {
		a,b := eq[0],eq[1]
		if _, ok := id[a]; !ok {
			id[a] = len(id)
		}
		if _, ok := id[b]; !ok {
			id[b] = len(id)
		}
	}

	// 建图
	graph := make([][]edge,len(id))
	for i, eq := range equations {
		u,v := id[eq[0]],id[eq[1]]
		graph[u] = append(graph[u], edge{v,values[i]})
		graph[v] = append(graph[v], edge{u,1/values[i]})
	}

	ans := make([]float64,len(queries))
	for i, query := range queries {
		start,ok1 := id[query[0]]
		end,ok2 := id[query[1]]
		if !ok1 || !ok2 {
			ans[i] = -1
		} else {
			ans[i] = bfs(graph,start,end)
		}
	}
	return ans
}

type edge struct {
	to int
	weight float64
}

func bfs(graph [][]edge, start, end int) float64 {
	ratios := make([]float64,len(graph))
	ratios[start] = 1
	queue := []int{start}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == end {
			return ratios[end]
		}
		for _, e := range graph[v] {
			if w := e.to; ratios[w] == 0 {
				ratios[w] = ratios[v] * e.weight
				queue = append(queue, w)
			}
		}
	}
	return -1
}
```
 时间复杂度O(ML+Q(L+M)),M为边的个数，Q为查询的数量，L为字符串的平均长度
 空间复杂度O(NL+M),N为点的数量
 
 2、floyd算法(弗洛伊德)
 [floyd算法](https://www.cnblogs.com/wangyuliang/p/9216365.html)
 ```go
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给变量编号
	id := map[string]int{}
	for _, eq := range equations {
		a,b := eq[0],eq[1]
		if _, ok := id[a]; !ok {
			id[a] = len(id)
		}
		if _, ok := id[b]; !ok {
			id[b] = len(id)
		}
	}

	// 建图
	graph := make([][]float64,len(id))
	for i := range graph {
		graph[i] = make([]float64,len(id))
	}
	for i, eq := range equations {
		u,v := id[eq[0]],id[eq[1]]
		graph[u][v] = values[i]
		graph[v][u] = 1/values[i]
	}

	// floyd算法
	for k := range graph {
		for i := range graph {
			for j := range graph {
				if graph[i][k] > 0 && graph[k][j] > 0 {
					graph[i][j] = graph[i][k]*graph[k][j]
				}
			}
		}
	}

	ans := make([]float64,len(queries))
	for i, query := range queries {
		start,ok1 := id[query[0]]
		end,ok2 := id[query[1]]
		if !ok1 || !ok2 {
			ans[i] = -1
		} else {
			ans[i] = graph[start][end]
		}
	}
	return ans
}
```
 时间复杂度O(ML+N^3^+QL),N^3^ 为floyd需要的时间
 空间复杂度O(NL+N^2^)
 
 3、加权并查集
 ```go
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给变量编号
	id := map[string]int{}
	for _, eq := range equations {
		a,b := eq[0],eq[1]
		if _, ok := id[a]; !ok {
			id[a] = len(id)
		}
		if _, ok := id[b]; !ok {
			id[b] = len(id)
		}
	}

	fa := make([]int,len(id))
	w := make([]float64,len(id))
	for i := range fa {
		fa[i] = i
		w[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			f := find(fa[x])
			w[x] *= w[fa[x]]
			fa[x] = f
		}
		return fa[x]
	}
	merge := func(from,to int,val float64) {
		fFrom,fTo := find(from),find(to)
		w[fFrom] = val*w[to]/w[from]
		fa[fFrom] = fTo
	}

	for i, eq := range equations {
		merge(id[eq[0]],id[eq[1]],values[i])
	}

	ans := make([]float64,len(queries))
	for i, q := range queries {
		start,ok1 := id[q[0]]
		end,ok2 := id[q[1]]
		if ok1 && ok2 && find(start) == find(end) {
			ans[i] = w[start]/w[end]
		} else {
			ans[i] = -1
		}
	}
	return ans
}
```
 时间复杂度：O(ML+N+MlogN+Q⋅(L+logN))。构建图需要 O(ML)O(ML) 的时间；初始化并查集需要 O(N)O(N) 的初始化时间；构建并查集的单次操作复杂度为 O(\log N)O(logN)，共需 O(M\log N)O(MlogN) 的时间；每个查询需要 O(L)O(L) 的字符串比较以及 O(\log N)O(logN) 的查询。
 空间复杂度：O(NL)O(NL)。哈希表需要 O(NL)O(NL) 的空间，并查集需要 O(N)O(N) 的空间。
 