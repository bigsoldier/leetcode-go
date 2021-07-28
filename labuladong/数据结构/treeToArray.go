package main

// 二叉堆实现优先级队列
type pq struct {
	data  []int
	count int // 元素个数
}

func newPQ(cap int) pq {
	return pq{data: make([]int, 0, cap+1)}
}

func (p *pq) max() int {
	return p.data[1]
}

func (p *pq) insert(e int) {
	p.count++
	p.data[p.count] = e
	p.swim(p.count)
}

func (p *pq) delMax() int {
	k := p.data[1]
	p.swap(1, p.count)
	p.data[p.count] = 0
	p.count--
	p.sink(1)
	return k
}

func (p *pq) swim(k int) {
	for k > 1 && p.less(parent(k), k) {
		p.swap(parent(k), k)
		k = parent(k)
	}
}

func (p *pq) sink(k int) {
	for left(k) <= p.count {
		l, r := left(k), right(k)
		if r <= p.count && p.less(l, r) {
			l = r
		}
		if p.less(l, k) {
			break
		}
		p.swap(l, k)
		k = l
	}
}

func (p *pq) swap(i, j int) {
	p.data[i], p.data[j] = p.data[j], p.data[i]
}

func (p *pq) less(i, j int) bool {
	return p.data[i] < p.data[j]
}

func parent(root int) int {
	return root / 2
}

func left(root int) int {
	return root * 2
}

func right(root int) int {
	return root*2 + 1
}
