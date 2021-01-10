#### 题目
<p>现在你总共有 <em>n</em> 门课需要选，记为&nbsp;<code>0</code>&nbsp;到&nbsp;<code>n-1</code>。</p>

<p>在选修某些课程之前需要一些先修课程。&nbsp;例如，想要学习课程 0 ，你需要先完成课程&nbsp;1 ，我们用一个匹配来表示他们: <code>[0,1]</code></p>

<p>给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。</p>

<p>可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> 2, [[1,0]] 
<strong>输出: </strong><code>[0,1]</code>
<strong>解释:</strong>&nbsp;总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 <code>[0,1] 。</code></pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> 4, [[1,0],[2,0],[3,1],[3,2]]
<strong>输出: </strong><code>[0,1,2,3] or [0,2,1,3]</code>
<strong>解释:</strong>&nbsp;总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
&nbsp;    因此，一个正确的课程顺序是&nbsp;<code>[0,1,2,3]</code> 。另一个正确的排序是&nbsp;<code>[0,2,1,3]</code> 。
</pre>

<p><strong>说明:</strong></p>

<ol>
	<li>输入的先决条件是由<strong>边缘列表</strong>表示的图形，而不是邻接矩阵。详情请参见<a href="http://blog.csdn.net/woaidapaopao/article/details/51732947" target="_blank">图的表示法</a>。</li>
	<li>你可以假定输入的先决条件中没有重复的边。</li>
</ol>

<p><strong>提示:</strong></p>

<ol>
	<li>这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。</li>
	<li><a href="https://www.coursera.org/specializations/algorithms" target="_blank">通过 DFS 进行拓扑排序</a> - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。</li>
	<li>
	<p>拓扑排序也可以通过&nbsp;<a href="https://baike.baidu.com/item/%E5%AE%BD%E5%BA%A6%E4%BC%98%E5%85%88%E6%90%9C%E7%B4%A2/5224802?fr=aladdin&amp;fromid=2148012&amp;fromtitle=%E5%B9%BF%E5%BA%A6%E4%BC%98%E5%85%88%E6%90%9C%E7%B4%A2" target="_blank">BFS</a>&nbsp;完成。</p>
	</li>
</ol>


 #### 题解
 同207题
 1、广度优先算法
 ```go
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges = make([][]int,numCourses)
		indeg = make([]int,numCourses) // 该们课程的前置数量
		result []int
	)
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
		indeg[p[0]]++
	}
	var queue []int
	for i, v := range indeg {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		result = append(result, q)
		for _, e := range edges[q] {
			indeg[e]--
			if indeg[e] == 0 {
				queue = append(queue, e)
			}
		}
	}
	if len(result) != numCourses {
		return nil
	}
	return result
}
```
 时间复杂度O(n+m),空间复杂度O(n+m),n是课程数,m是先修课程数
 
 2、深度优先算法
 ```go
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges = make([][]int,numCourses)
		visited = make([]int,numCourses) // 0表示未搜索，1表示进行中，2表示已完成
		valid = true // 是否无环
		dfs func(i int)
		result []int
	)
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
	}
	dfs = func(i int) {
		visited[i] = 1
		for _, v := range edges[i] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[i] = 2
		result = append(result, i)
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return nil
	}
	for i := 0; i < len(result)/2; i++ {
		result[i],result[len(result)-i-1] = result[len(result)-i-1],result[i]
	}
	return result
}
```