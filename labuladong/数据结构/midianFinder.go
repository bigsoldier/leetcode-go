package main

import "container/heap"

// 数据流的中位数

// 使用两个堆，一个大根堆存小数，一个小根堆存大数
type MedianFinder struct {
	left  *MaxHeap
	right *MinHeap
}

/** initialize your data structure here. */
func Constructor3() MedianFinder {
	left := new(MaxHeap)
	right := new(MinHeap)
	heap.Init(left)
	heap.Init(right)
	return MedianFinder{left: left, right: right}
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() >= this.right.Len() {
		heap.Push(this.left, num)
		a := heap.Pop(this.left)
		heap.Push(this.right, a)
	} else {
		heap.Push(this.right, num)
		a := heap.Pop(this.right)
		heap.Push(this.left, a)
	}

	heap.Fix()
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() < this.right.Len() {
		return float64(heap.Pop(this.right).(int))
	} else if this.left.Len() > this.right.Len() {
		return float64(heap.Pop(this.left).(int))
	} else {
		return float64(heap.Pop(this.left).(int)+heap.Pop(this.right).(int)) / 2
	}
}

// 大根堆
type MaxHeap struct {
	intHeap
}

func (m MaxHeap) Less(i, j int) bool {
	return m.intHeap[i] > m.intHeap[j]
}

// 小根堆
type MinHeap struct {
	intHeap
}

func (m MinHeap) Less(i, j int) bool {
	return m.intHeap[i] < m.intHeap[j]
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	a := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return a
}
