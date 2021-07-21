package main

// 想要让put和get方法的时间复杂度为O(1)
type LRUCache struct {
	data     *DLinkedList
	cache    map[int]*DLinkedNode
	capacity int
}

// 双向链表
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

type DLinkedList struct {
	head, tail *DLinkedNode
	size       int
}

func initDLinkedNode(key, val int) *DLinkedNode {
	return &DLinkedNode{key: key, value: val}
}

func initDLinkedList() *DLinkedList {
	list := &DLinkedList{
		head: initDLinkedNode(0, 0),
		tail: initDLinkedNode(0, 0),
		size: 0,
	}
	list.head.next = list.tail
	list.tail.prev = list.head
	return list
}

// 先不要急着实现Get和Put方法，容易漏操作

// 尾部添加节点
func (this *DLinkedList) addLast(x *DLinkedNode) {
	x.prev = this.tail.next
	x.next = this.tail
	this.tail.prev.next = x
	this.tail.prev = x
	this.size++
}

// 删除链表中的x节点
func (this *DLinkedList) remove(x *DLinkedNode) {
	x.prev.next = x.next
	x.next.prev = x.prev
	this.size--
}

// 删除头节点
func (this *DLinkedList) removeFirst() *DLinkedNode {
	if this.head.next == this.tail {
		return nil
	}
	first := this.head.next
	this.remove(first)
	return first
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		data:     initDLinkedList(),
		cache:    make(map[int]*DLinkedNode),
		capacity: capacity,
	}
	return lru
}

// key提升为最近使用
func (this *LRUCache) makeRecently(key int) {
	x := this.cache[key]
	this.data.remove(x)
	this.data.addLast(x)
}

// 添加元素
func (this *LRUCache) addRecently(key, val int) {
	x := initDLinkedNode(key, val)
	this.data.addLast(x)
	this.cache[key] = x
}

// 删除key
func (this *LRUCache) deleteKey(key int) {
	x := this.cache[key]
	this.data.remove(x)
	delete(this.cache, key)
}

// 删除最久未使用的元素
func (this *LRUCache) removeLeastRecently() {
	x := this.data.removeFirst()
	delete(this.cache, x.key)
}

func (this *LRUCache) Get(key int) int {
	if x, ok := this.cache[key]; !ok {
		return -1
	} else {
		this.makeRecently(key)
		return x.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		this.deleteKey(key)
	} else {
		if this.capacity == this.data.size {
			this.removeLeastRecently()
		}
	}
	this.addRecently(key, value)
}
