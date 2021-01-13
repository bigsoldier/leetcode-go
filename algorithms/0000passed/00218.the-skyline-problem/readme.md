#### 题目
<p>城市的天际线是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。现在，假设您获得了城市风光照片（图A）上<strong>显示的所有建筑物的位置和高度</strong>，请编写一个程序以输出由这些建筑物<strong>形成的天际线</strong>（图B）。</p>

<p><a href="/static/images/problemset/skyline1.jpg" target="_blank"><img alt="Buildings" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/skyline1.png" style="width: 300px;"> </a> <a href="/static/images/problemset/skyline2.jpg" target="_blank"> <img alt="Skyline Contour" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/skyline2.png" style="width: 300px;"> </a></p>

<p>每个建筑物的几何信息用三元组&nbsp;<code>[Li，Ri，Hi]</code> 表示，其中 <code>Li</code> 和 <code>Ri</code> 分别是第 i 座建筑物左右边缘的 x 坐标，<code>Hi</code> 是其高度。可以保证&nbsp;<code>0 &le; Li, Ri &le; INT_MAX</code>,&nbsp;<code>0 &lt; Hi &le; INT_MAX</code> 和 <code>Ri - Li &gt; 0</code>。您可以假设所有建筑物都是在绝对平坦且高度为 0 的表面上的完美矩形。</p>

<p>例如，图A中所有建筑物的尺寸记录为：<code>[ [2 9 10], [3 7 15], [5 12 12], [15 20 10], [19 24 8] ] </code>。</p>

<p>输出是以&nbsp;<code>[ [x1,y1], [x2, y2], [x3, y3], ... ]</code> 格式的&ldquo;<strong>关键点</strong>&rdquo;（图B中的红点）的列表，它们唯一地定义了天际线。<strong>关键点是水平线段的左端点</strong>。请注意，最右侧建筑物的最后一个关键点仅用于标记天际线的终点，并始终为零高度。此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。</p>

<p>例如，图B中的天际线应该表示为：<code>[ [2 10], [3 15], [7 12], [12 0], [15 10], [20 8], [24, 0] ]</code>。</p>

<p><strong>说明:</strong></p>

<ul>
	<li>任何输入列表中的建筑物数量保证在 <code>[0, 10000]</code>&nbsp;范围内。</li>
	<li>输入列表已经按左&nbsp;<code>x</code> 坐标&nbsp;<code>Li</code>&nbsp; 进行升序排列。</li>
	<li>输出列表必须按 x 位排序。</li>
	<li>输出天际线中不得有连续的相同高度的水平线。例如 <code>[...[2 3], [4 5], [7 5], [11 5], [12 7]...]</code> 是不正确的答案；三条高度为 5 的线应该在最终输出中合并为一个：<code>[...[2 3], [4 5], [12 7], ...]</code></li>
</ul>

输入：buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]

输出：[[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]

示例2:
输入：buildings = [[0,2,3],[2,5,3]]

输出：[[0,3],[5,0]]

 #### 题解
 1、归并
 考虑如果只有一个建筑[x,y,h]，那么很明显输出的解是[[x,h],[y,0]]
 
 那么如果有建筑 A B C D E ,我们知道 A B C 和 D E 的输出，怎么把两个输出合并就能获取最终解?
 
 例子:
 x = 0
 
 skyLine1 = {(1,11),(3,3),(8,4),(9,0)},h1=0
 
 skyLine2 = {(4,8),(9,10),(10,0)},h2=0
 
 先比较(1,11)和(4,8)，比较坐标 x= 1，然后高度 h1=11,最终我们取(1,11),
 
 然后比较(3,3)和(4,8),坐标x=3, 高度h1=3,我们取(3,3)
 
 然后比较(8,4)和(4,8),坐标x=4,高度h2=8,我们取(4,8)
 
 然后比较(8,4)和(9,10),坐标x=8,高度h1=4,我们取(8,8),由于最高和上一个一样，舍弃这条
 
 
 ```go
func getSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return nil
	}
	return merge(buildings,0,len(buildings)-1)
}

func merge(buildings [][]int, start, end int) (ans [][]int) {
	if start == end {
		// 只有一个建筑，将结果集加入
		ans = append(ans,
			[]int{buildings[start][0],buildings[start][2]},
			[]int{buildings[start][1],0})
		return ans
	}
	mid := start + (end-start)/2
	left := merge(buildings,start,mid)
	right := merge(buildings,mid+1,end)

	// 合并解
	var i,j,h1,h2 int
	for i < len(left) || j < len(right) {
		var x1,x2 = math.MaxInt64,math.MaxInt64
		if i < len(left) {
			x1 = left[i][0]
		}
		if j < len(right) {
			x2 = right[j][0]
		}
		// 比较两个坐标,取较小的那一边
		if x1 < x2 {
			h1 = left[i][1]
			i++
		} else if x1 > x2 {
			h2 = right[j][1]
			x1 = x2
			j++
		} else {
			h1 = left[i][1]
			h2 = right[j][1]
			i++
			j++
		}
		height := max(h1,h2)
		if len(ans) == 0 || height != ans[len(ans)-1][1] {
			ans = append(ans, []int{x1,height})
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
 时间复杂度O(nlogn),空间复杂度O(n)
 
 2、线性扫描
 ```go
func getSkyline(buildings [][]int) (ans [][]int) {
	if len(buildings) == 0 {
		return nil
	}

	highs := new(highHeap)
	heap.Init(highs)
	left := 0
	heap.Push(highs,left) // 最左边的高度

	edges := make([][3]int,0,2*len(buildings))
	for _, b := range buildings {
		edges = append(edges, [3]int{b[0],b[2],-1}) // 进入边
		edges = append(edges, [3]int{b[1],b[2],1}) // 退出边
	}
	sort.Slice(edges, func(i, j int) bool {
		if edges[i][0] != edges[j][0] {
			return edges[i][0]<edges[j][0]
		}
		// 当 es[i][0] == es[j][0] 的时候
		// i,j 分别为 进入边 与 退出边
		//     则，进入边应在前
		// i,j 同为 进入边
		//     则，e[1] 高的在前
		// i,j 同为 退出边
		//     则，e[1] 低的在前
		return edges[i][1]*edges[i][2] < edges[j][1]*edges[j][2]
	})
	
	// 从左向右依次扫描
	for _, e := range edges {
		if e[2] < 0 {// 进入边
			heap.Push(highs,e[1])
		} else {//退出边
			i := 0
			for i < len(*highs) {
				if (*highs)[i] == e[1] {
					break
				}
				i++
			}
			heap.Remove(highs,i)
		}
		// curr是当前的最高
		curr := (*highs)[0]
		if curr != left {
			ans = append(ans, []int{e[0],curr})
			left = curr
		}
	}
	return
}

type highHeap []int
func (h highHeap) Len() int {return len(h)}
func (h highHeap) Swap(i, j int) {h[i],h[j] = h[j],h[i]}
func (h highHeap) Less(i, j int) bool {return h[i]>h[j]}
func (h *highHeap) Push(x interface{}) {*h = append(*h, x.(int))}
func (h *highHeap) Pop() interface{} {
	ret := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ret
}
```