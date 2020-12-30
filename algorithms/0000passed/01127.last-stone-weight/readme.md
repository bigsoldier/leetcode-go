#### 题目
<p>有一堆石头，每块石头的重量都是正整数。</p>

<p>每一回合，从中选出两块<strong>最重的</strong>石头，然后将它们一起粉碎。假设石头的重量分别为&nbsp;<code>x</code> 和&nbsp;<code>y</code>，且&nbsp;<code>x &lt;= y</code>。那么粉碎的可能结果如下：</p>

<ul>
	<li>如果&nbsp;<code>x == y</code>，那么两块石头都会被完全粉碎；</li>
	<li>如果&nbsp;<code>x != y</code>，那么重量为&nbsp;<code>x</code>&nbsp;的石头将会完全粉碎，而重量为&nbsp;<code>y</code>&nbsp;的石头新重量为&nbsp;<code>y-x</code>。</li>
</ul>

<p>最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 <code>0</code>。</p>

<p>&nbsp;</p>

示例：

输入：[2,7,4,1,8,1]
输出：1
解释：
先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，
再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，
接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，
最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。

<p><strong>提示：</strong></p>

<ol>
	<li><code>1 &lt;= stones.length &lt;= 30</code></li>
	<li><code>1 &lt;= stones[i] &lt;= 1000</code></li>
</ol>


 #### 题解
 需要取出最大值，同时又要放入值，所以我们使用堆来处理这种问题
 ```go
type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {return h.IntSlice[i] > h.IntSlice[j]}
func (h *hp) Push(x interface{}) {h.IntSlice = append(h.IntSlice, x.(int))}
func (h *hp) Pop() interface{} {v := h.IntSlice[len(h.IntSlice)-1];h.IntSlice = h.IntSlice[:len(h.IntSlice)-1];return v}
func (h *hp) push(x int) {heap.Push(h,x)}
func (h *hp) pop() int {return heap.Pop(h).(int)}

func lastStoneWeight(stones []int) int {
	var q = &hp{stones}
	heap.Init(q)
	// 从堆中去除两个数据
	for q.Len() > 1 {
		x,y := q.pop(),q.pop()
		if x > y {
			q.push(x-y)
		}
	}
	if q.Len() > 0 {
		return q.pop()
	}
	return 0
}
```
 时间复杂度O(nlogn),空间复杂度O(n)