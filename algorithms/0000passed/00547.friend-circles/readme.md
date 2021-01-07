#### 题目
<p>班上有&nbsp;<strong>N&nbsp;</strong>名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B&nbsp;的朋友，B 是 C&nbsp;的朋友，那么我们可以认为 A 也是 C&nbsp;的朋友。所谓的朋友圈，是指所有朋友的集合。</p>

<p>给定一个&nbsp;<strong>N * N&nbsp;</strong>的矩阵&nbsp;<strong>M</strong>，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生<strong>互为</strong>朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> 
[[1,1,0],
 [1,1,0],
 [0,0,1]]
<strong>输出:</strong> 2 
<strong>说明：</strong>已知学生0和学生1互为朋友，他们在一个朋友圈。
第2个学生自己在一个朋友圈。所以返回2。
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> 
[[1,1,0],
 [1,1,1],
 [0,1,1]]
<strong>输出:</strong> 1
<strong>说明：</strong>已知学生0和学生1互为朋友，学生1和学生2互为朋友，所以学生0和学生2也是朋友，所以他们三个在一个朋友圈，返回1。
</pre>

<p><strong>注意：</strong></p>

<ol>
	<li>N 在[1,200]的范围内。</li>
	<li>对于所有学生，有M[i][i] = 1。</li>
	<li>如果有M[i][j] = 1，则有M[j][i] = 1。</li>
</ol>


 #### 题解
 1、深度优先搜索
 ```go
func findCircleNum(isConnected [][]int) (ans int) {
	isVisited := make([]bool,len(isConnected))
	var dfs func(i int)
	dfs = func(i int) {
		isVisited[i] = true
		for index,connect := range isConnected[i] {
			if connect == 1 && !isVisited[index] {
				dfs(index)
			}
		}
	}
	for i, visited := range isVisited {
		if !visited {
			ans++
			dfs(i)
		}
	}
	return
}
```
 时间复杂度O(n^2^),空间复杂度O(n)
 
 2、广度优先搜索
 ```go
func findCircleNum(isConnected [][]int) (ans int) {
	isVisited := make([]bool,len(isConnected))
	for i, v := range isVisited {
		if !v {
			ans++
			queue := []int{i}
			for len(queue) > 0 {
				from := queue[0]
				queue = queue[1:]
				isVisited[from] = true
				for index, conn := range isConnected[from] {
					if conn==1&&!isVisited[index] {
						queue = append(queue, index)
					}
				}
			}
		}
	}
	return
}
```
 时间复杂度O(n^2^),空间复杂度O(n)
 
 3、并查集
 ```go
func findCircleNum(isConnected [][]int) (ans int) {
	n := len(isConnected)
	parent := make([]int,n)
	for i := range parent {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	merge := func(from,to int) {
		parent[find(from)] = find(to)
	}
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			if isConnected[i][j] == 1 {
				merge(i,j)
			}
		}
	}
	for i, p := range parent {
		if i == p {
			ans++
		}
	}
	return
}
```
 时间复杂度O(n^2^logn),空间复杂度O(n)