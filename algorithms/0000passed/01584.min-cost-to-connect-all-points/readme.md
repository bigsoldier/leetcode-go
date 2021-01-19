#### 题目
 给你一个`points`数组，表示2D平面上的一些点，其中`points[i]=[xi,yi]。
 
 连接点[xi,yi]和点[xj,yj]的费用为他们之间的 **曼哈顿距离**: `|xi-xj|+|yi-yj|`，其中`|val|`表示`val`的绝对值。
 
 请你返回将所有点连接的最小总费用。只有任意两点之间  **有且仅有** 一条简单路径时，才会认为所有点都已连接。
 
 ![](https://assets.leetcode.com/uploads/2020/08/26/d.png)
 **输入:** points=[[0,0],[2,2],[3,10],[5,2],[7,0]]
 **输出:** 20
 **解释:** ![](https://assets.leetcode.com/uploads/2020/08/26/c.png)
 
 示例 2：
 
 输入：points = [[3,12],[-2,5],[-4,1]]
 输出：18
 示例 3：
 
 输入：points = [[0,0],[1,1],[1,0],[-1,1]]
 输出：4
 示例 4：
 
 输入：points = [[-1000000,-1000000],[1000000,1000000]]
 输出：4000000
 示例 5：
 
 输入：points = [[0,0]]
 输出：0

 #### 题解
 1、Kruskal算法
 
 ```go
func minCostConnectPoints(points [][]int) (ans int) {
	n := len(points)
	type edge struct{ v, w, dist int }
	edges := []edge{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{i, j, dist(points[i], points[j])})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	fa, rank := make([]int, n), make([]int, n)
	for i := 0; i < len(fa); i++ {
		fa[i] = i
		rank[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if x != fa[x] {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(x, y int) bool {
		fx, fy := find(x), find(y)
		if fx == fy {
			return false
		}
		if rank[fx] < rank[fy] {
			fx, fy = fy, fx
		}
		rank[fx] += rank[fy]
		fa[fy] = fx
		return true
	}
	left := n - 1
	for _, e := range edges {
		if merge(e.v, e.w) {
			ans += e.dist
			left--
			if left == 0 {
				break
			}
		}
	}
	return
}

func dist(p, q []int) int {
	return abs(p[0]-q[0]) + abs(p[1]-q[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
 时间复杂度O(n^2^logn),空间复杂度O(n^2^)
 
 2、最小生成树
 ```go
//  最小生成树
func minCostConnectPoints(points [][]int) int {
	n := len(points)
	visited := make([]bool, n)

	minDist := make([]int, n)
	for i := 1; i < n; i++ {
		minDist[i] = math.MaxInt32
	}

	res := 0

	for _ = range visited {
		nextPoint := -1
		d := math.MaxInt32
		for j := range visited {
			if !visited[j] && minDist[j] < d {
				d = minDist[j]
				nextPoint = j
			}
		}

		res += d

		visited[nextPoint] = true

		for j := range visited {
			if !visited[j] {
				minDist[j] = min(minDist[j], getDistance(points[j], points[nextPoint]))
			}
		}
	}

	return res
}

func getDistance(p1, p2 []int) int {
	return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```