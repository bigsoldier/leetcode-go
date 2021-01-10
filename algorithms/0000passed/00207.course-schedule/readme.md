#### 题目
<p>你这个学期必须选修 <code>numCourse</code> 门课程，记为&nbsp;<code>0</code>&nbsp;到&nbsp;<code>numCourse-1</code> 。</p>

<p>在选修某些课程之前需要一些先修课程。&nbsp;例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：<code>[0,1]</code></p>

<p>给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> 2, [[1,0]] 
<strong>输出: </strong>true
<strong>解释:</strong>&nbsp;总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> 2, [[1,0],[0,1]]
<strong>输出: </strong>false
<strong>解释:</strong>&nbsp;总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ol>
	<li>输入的先决条件是由 <strong>边缘列表</strong> 表示的图形，而不是 邻接矩阵 。详情请参见<a href="http://blog.csdn.net/woaidapaopao/article/details/51732947" target="_blank">图的表示法</a>。</li>
	<li>你可以假定输入的先决条件中没有重复的边。</li>
	<li><code>1 &lt;=&nbsp;numCourses &lt;= 10^5</code></li>
</ol>


 #### 题解
 1、广度优先
 
 [[2.0],[2,1],[3,2]]
 意思是想要学习课程2，需要先完成课程0
 
 那么我们可以得出课程的先修课程情况
 - 0:
 - 1:
 - 2:0,1
 - 3:2
 
 ```go
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges = make([][]int,numCourses) // 记录某课程完成后可以进行的课程
		indeg = make([]int,numCourses)// 记录课程的先修量
		result []int
	)

	for _, info := range prerequisites {
		// info第一位表示后修课程，第二位表示先修课程
		edges[info[1]] = append(edges[info[1]], info[0]) // 这里添加先修课程完之后对应的后修课程
		indeg[info[0]]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		course := q[0]
		q = q[1:]
		result = append(result, course)
		for _, v := range edges[course] {
			indeg[v]--
			if indeg[v]==0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}
```
 时间复杂度O(m+n),空间复杂度O(m+n),m是先修课程数，n是总课程数
 
 2、深度优先
 ```go
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges = make([][]int,numCourses)
		// 标记每个节点的状态:0未搜索，1搜索中，2已完成
		visited = make([]int,numCourses)
		valid = true // 是否无环
		result []int
		dfs func(i int)
	)

	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
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
	return valid
}
```
 时间复杂度O(m+n),空间复杂度O(m+n),m是先修课程数，n是总课程数