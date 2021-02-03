package code

import (
	"container/heap"
	"sort"
)

var small, large *hp
var delayed map[int]int

func medianSlidingWindow(nums []int, k int) (ans []float64) {
	delayed = map[int]int{}     // 哈希表，记录延迟删除的元素
	small, large = &hp{}, &hp{} // 大根堆，小根堆
	makeBalance := func() {
		// 调整small和large中的元素个数
		if small.size > large.size+1 {
			large.push(-small.pop())
			small.prune()
		} else if small.size < large.size {
			small.push(-large.pop())
			large.prune()
		}
	}

	insert := func(x int) {
		if small.Len() == 0 || x <= -small.IntSlice[0] {
			small.push(-x)
		} else {
			large.push(x)
		}
		makeBalance()
	}
	erase := func(x int) {
		delayed[x]++
		if x <= -small.IntSlice[0] {
			small.size--
			if x == -small.IntSlice[0] {
				small.prune()
			}
		} else {
			large.size--
			if x == large.IntSlice[0] {
				large.prune()
			}
		}
		makeBalance()
	}
	getMedian := func() float64 {
		if k&1 > 0 { // 奇数
			return float64(-small.IntSlice[0])
		}
		return float64(-small.IntSlice[0]+large.IntSlice[0]) / 2
	}
	for _, num := range nums[:k] {
		insert(num)
	}
	n := len(nums)
	ans = make([]float64, 1, n-k+1)
	ans[0] = getMedian()
	for i := k; i < n; i++ {
		insert(nums[i])
		erase(nums[i-k])
		ans = append(ans, getMedian())
	}
	return
}

type hp struct {
	sort.IntSlice
	size int
}

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	p := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return p
}
func (h *hp) push(x int) {
	h.size++
	heap.Push(h, x)
}
func (h *hp) pop() int {
	h.size--
	return heap.Pop(h).(int)
}
func (h *hp) prune() {
	for h.Len() > 0 {
		top := h.IntSlice[0]
		if h == small {
			top = -top
		}
		if d, has := delayed[top]; has {
			if d > 1 {
				delayed[top]--
			} else {
				delete(delayed, top)
			}
			heap.Pop(h)
		} else {
			break
		}
	}
}
