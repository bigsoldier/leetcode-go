package main

type NestedIterator struct {
	nestedList []*NestedInteger
}

func Constructor1(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{nestedList: nestedList}
}

func (this *NestedIterator) Next() int {
	queue := this.nestedList[0]
	val := queue.GetInteger()
	this.nestedList = this.nestedList[1:]
	return val
}

func (this *NestedIterator) HasNext() bool {
	for len(this.nestedList) > 0 && !this.nestedList[0].IsInteger() {
		queue := this.nestedList[0]
		this.nestedList = this.nestedList[1:]
		first := queue.GetList()
		// 将打开的元素加入头部
		for i := len(first) - 1; i >= 0; i-- {
			this.nestedList = append([]*NestedInteger{first[i]}, this.nestedList...)
		}
	}
	return len(this.nestedList) != 0
}
