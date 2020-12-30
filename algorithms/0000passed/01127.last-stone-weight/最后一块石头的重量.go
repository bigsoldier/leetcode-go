package code

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(x interface{}) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *hp) Pop() interface{} {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}
func (h *hp) push(x int) { heap.Push(h, x) }
func (h *hp) pop() int   { return heap.Pop(h).(int) }

func lastStoneWeight(stones []int) int {
	var q = &hp{stones}
	heap.Init(q)
	// 从堆中去除两个数据
	for q.Len() > 1 {
		x, y := q.pop(), q.pop()
		if x > y {
			q.push(x - y)
		}
	}
	if q.Len() > 0 {
		return q.pop()
	}
	return 0
}
