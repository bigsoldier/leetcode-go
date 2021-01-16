#### 题目
<p>我们有一组包含1和0的网格；其中1表示砖块。&nbsp;当且仅当一块砖直接连接到网格的顶部，或者它至少有一块相邻（4&nbsp;个方向之一）砖块不会掉落时，它才不会落下。</p>

<p>我们会依次消除一些砖块。每当我们消除&nbsp;(i, j) 位置时， 对应位置的砖块（若存在）会消失，然后其他的砖块可能因为这个消除而落下。</p>

<p>返回一个数组表示每次消除操作对应落下的砖块数目。</p>

<pre><strong>示例 1：</strong>
<strong>输入：</strong>
grid = [[1,0,0,0],[1,1,1,0]]
hits = [[1,0]]
<strong>输出:</strong> [2]
<strong>解释: </strong>
如果我们消除(1, 0)位置的砖块, 在(1, 1) 和(1, 2) 的砖块会落下。所以我们应该返回2。</pre>

<pre><strong>示例 2：</strong>
<strong>输入：</strong>
grid = [[1,0,0,0],[1,1,0,0]]
hits = [[1,1],[1,0]]
<strong>输出：</strong>[0,0]
<strong>解释：</strong>
当我们消除(1, 0)的砖块时，(1, 1)的砖块已经由于上一步消除而消失了。所以每次消除操作不会造成砖块落下。注意(1, 0)砖块不会记作落下的砖块。</pre>

<p><strong>注意:</strong></p>

<ul>
	<li>网格的行数和列数的范围是[1, 200]。</li>
	<li>消除的数字不会超过网格的区域。</li>
	<li>可以保证每次的消除都不相同，并且位于网格的内部。</li>
	<li>一个消除的位置可能没有砖块，如果这样的话，就不会有砖块落下。</li>
</ul>


 #### 题解
 并查集
 
 打砖块是将原来的一个集合分成多个集合，那么逆序的话就是将多个集合合并成一个集合，采用并查集的方式。
 ```go
func hitBricks(grid [][]int, hits [][]int) []int {
	rows,cols := len(grid),len(grid[0])
	total := rows*cols

	var fa = make([]int, total+1)
	size := make([]int, total+1)
	for i := 0; i < len(fa); i++ {
		fa[i] = i
		size[i] = 1
	}
	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(x,y int) {
		xRoot,yRoot := find(x),find(y)
		if xRoot != yRoot {
			size[yRoot] += size[xRoot]
			fa[xRoot] = yRoot
		}
	}
	index := func(x,y int) int {
		return x*cols+y
	}
	isArea := func(x,y int) bool {
		return x>=0&&y>=0&&x<rows&&y<cols
	}
	// 1、将grid砖头全部击碎
	var graph = make([][]int,rows)
	for i := 0; i < rows; i++ {
		graph[i] = make([]int,cols)
		for j := 0; j < cols; j++ {
			graph[i][j] = grid[i][j]
		}
	}
	for _, hit := range hits {
		graph[hit[0]][hit[1]] = 0
	}
	// 2、建图
	// 将下标为0的这一行与屋顶相连
	for j := 0; j < cols; j++ {
		if graph[0][j] == 1 {
			merge(index(0,j), total)
		}
	}
	// 其余网格，如果是砖块，查看上、左
	for i := 1; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if graph[i][j] == 1 {
				// 如果上方是砖块
				if graph[i-1][j] == 1 {
					merge(index(i-1,j),index(i,j))
				}
				// 如果左方是砖块
				if j >0 && graph[i][j-1] == 1 {
					merge(index(i,j-1),index(i,j))
				}
			}
		}
	}
	// 按照hits的逆序，补回砖块
	hitsLen := len(hits)
	var res = make([]int,hitsLen) // 补回砖块与屋顶相连的增量
	for i := hitsLen-1; i >= 0; i-- {
		x,y := hits[i][0],hits[i][1]
		if grid[x][y] == 0 { // 如果原来是0，那么就不存在砖块
			continue
		}
		// 补回之前的砖块数
		origin := size[find(total)]
		// 如果是再第一行，告诉并查集它与屋顶香兰
		if x == 0 {
			merge(index(x,y), total)
		}
		// 查询4个方向
		for _, direct := range direction {
			newX,newY := x+direct[0],y+direct[1]
			if isArea(newX,newY) && graph[newX][newY]==1 {
				merge(index(x,y),index(newX,newY))
			}
		}
		// 补回后的砖块数
		curr := size[find(total)]
		if curr - origin -1 > 0  {
			res[i] = curr -origin-1
		}
		graph[x][y] = 1
	}
	return res
}

var direction = [][]int{{0,1},{1,0},{-1,0},{0,-1}}
```
 时间复杂度O(n*hw*log(hw)),空间复杂度O(hw)