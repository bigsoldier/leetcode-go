package code

import "container/heap"

type MedianFinder struct {
	left  *MaxHeap
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
		heap.Push(this.left, num)
	} else {
		heap.Push(this.right, num)
	}
	if this.right.Len() > 0 && this.left.intHeap[0] > this.right.intHeap[0] {
		this.left.intHeap[0], this.right.intHeap[0] = this.right.intHeap[0], this.left.intHeap[0]
		heap.Fix(this.left, 0)
		heap.Fix(this.right, 0)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.intHeap.Len() == this.right.intHeap.Len() {
		return float64(this.left.intHeap[0]+this.right.intHeap[0]) * 0.5
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

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h intHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{} {
	p := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return p
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
