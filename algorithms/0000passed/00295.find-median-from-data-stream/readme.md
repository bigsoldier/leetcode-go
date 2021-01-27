#### 题目
<p>中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。</p>

<p>例如，</p>

<p>[2,3,4]&nbsp;的中位数是 3</p>

<p>[2,3] 的中位数是 (2 + 3) / 2 = 2.5</p>

<p>设计一个支持以下两种操作的数据结构：</p>

<ul>
	<li>void addNum(int num) - 从数据流中添加一个整数到数据结构中。</li>
	<li>double findMedian() - 返回目前所有元素的中位数。</li>
</ul>

<p><strong>示例：</strong></p>

<pre>addNum(1)
addNum(2)
findMedian() -&gt; 1.5
addNum(3) 
findMedian() -&gt; 2</pre>

<p><strong>进阶:</strong></p>

<ol>
	<li>如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？</li>
	<li>如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？</li>
</ol>


 #### 题解
 1、简单排序
 ```go
type MedianFinder struct {
	nums []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	this.nums = append(this.nums, num)
}

func (this *MedianFinder) FindMedian() float64 {
	sort.Ints(this.nums)
	n := len(this.nums)
	if n%2 == 1 {
		return float64(this.nums[n/2])
	} else {
		return float64(this.nums[n/2-1] + this.nums[n/2])*0.5
	}
}
```
 时间复杂度O(nlogn)+O(1),空间复杂度O(n)
 
 2、插入排序
 ```go
type MedianFinder struct {
	nums []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	this.nums = append(this.nums, num)
	n := len(this.nums)
	for i := n-2; i >= 0 && this.nums[i+1] < this.nums[i]; i-- {
		this.nums[i+1],this.nums[i] = this.nums[i],this.nums[i+1]
	}
}

func (this *MedianFinder) FindMedian() float64 {
	n := len(this.nums)
	if n%2 == 1 {
		return float64(this.nums[n/2])
	} else {
		return float64(this.nums[n/2-1] + this.nums[n/2])*0.5
	}
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 3、两个堆
 需要实现一个大根堆和一个小根堆。大根堆存储较小一半的数，小根堆存储较大一半的数。
 - 添加一个数，添加到大根堆，然后从大根堆中pop到小根堆中
 - 如果小根堆的大小大于大根堆，从小根堆中pop到大根堆中
 
 - 大根堆多返回大根堆的值，小返回两个堆的平均值
 ```go
type MedianFinder struct {
	left *MaxHeap
	right *MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	var left = new(MaxHeap)
	var right = new(MinHeap)
	heap.Init(left)
	heap.Init(right)
	return MedianFinder{left, right}
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() == this.right.Len() {
		heap.Push(this.left,num)
	} else {
		heap.Push(this.right,num)
	}
	if this.right.Len() > 0 && this.left.intHeap[0] > this.right.intHeap[0] {
		this.left.intHeap[0],this.right.intHeap[0] = this.right.intHeap[0],this.left.intHeap[0]
		heap.Fix(this.left,0)
		heap.Fix(this.right,0)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.intHeap.Len() == this.right.intHeap.Len() {
		return float64(this.left.intHeap[0] + this.right.intHeap[0])*0.5
	}
	return float64(this.left.intHeap[0])
}

type MaxHeap struct {
	intHeap
}

func (m MaxHeap) Less(i, j int) bool {
	return m.intHeap[i] > m.intHeap[j]
}

type MinHeap struct {
	intHeap
}

type intHeap []int
func (h intHeap) Len() int {return len(h)}
func (h intHeap) Swap(i, j int) {h[i],h[j] = h[j],h[i]}
func (h intHeap) Less(i, j int) bool {return h[i]<h[j]}
func (h *intHeap) Push(x interface{}) {*h = append(*h, x.(int))}
func (h *intHeap) Pop() interface{} {
	p := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return p
}
```
 时间复杂度O(logn),空间复杂度O(n)