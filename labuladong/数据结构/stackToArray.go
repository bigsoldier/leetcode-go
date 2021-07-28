package main

// 栈实现队列
type MyQueue struct {
	s1, s2 []int
}

/** Initialize your data structure here. */
func Constructor5() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	for len(this.s1) > 0 {
		this.s2 = append(this.s2, this.Pop())
	}
	this.s2 = append(this.s2, x)
	for len(this.s2) > 0 {
		this.s1 = append(this.s1, this.s2[len(this.s2)-1])
		this.s2 = this.s2[:len(this.s2)-1]
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	k := this.s1[len(this.s1)-1]
	this.s1 = this.s1[:len(this.s1)-1]
	return k
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.s1[len(this.s1)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0
}

// 用队列实现栈
type MyStack struct {
	q []int
}

/** Initialize your data structure here. */
func Constructor6() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	n := len(this.q)
	this.q = append(this.q, x)
	for ; n > 0; n-- {
		this.q = append(this.q, this.q[0])
		this.q = this.q[1:]
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	x := this.q[0]
	this.q = this.q[1:]
	return x
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}
