#### 题目
<p>使用队列实现栈的下列操作：</p>

<ul>
	<li>push(x) -- 元素 x 入栈</li>
	<li>pop() -- 移除栈顶元素</li>
	<li>top() -- 获取栈顶元素</li>
	<li>empty() -- 返回栈是否为空</li>
</ul>

<p><strong>注意:</strong></p>

<ul>
	<li>你只能使用队列的基本操作-- 也就是&nbsp;<code>push to back</code>, <code>peek/pop from front</code>, <code>size</code>, 和&nbsp;<code>is empty</code>&nbsp;这些操作是合法的。</li>
	<li>你所使用的语言也许不支持队列。&nbsp;你可以使用 list 或者 deque（双端队列）来模拟一个队列&nbsp;, 只要是标准的队列操作即可。</li>
	<li>你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。</li>
</ul>


 #### 题解
 1、两个队列实现栈
 queue1是存储栈的元素，queue2用来辅助入栈元素。
 
 出栈时直接取queue1的队列0号元素
 
 入栈时，将元素放入queue2中，然后将queue1中的元素依次加入queue2中
 
 ```go
type MyStack struct {
	queue1,queue2 []int // queue1是存储栈元素，queue2入栈辅助队列
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue2 = append(this.queue2, x)
	for len(this.queue1) > 0 {
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}
	this.queue1,this.queue2 = this.queue2,this.queue1
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	v := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return v
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue1[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}
```
 入栈时间复杂度O(n),其他操作时间复杂度O(1),空间复杂度O(1)

 2、一个队列实现栈
 ```go
type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	n := len(this.queue)
	this.queue = append(this.queue, x)
	for ; n > 0; n-- {
		this.queue = append(this.queue, this.queue[0])
		this.queue = this.queue[1:]
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	v := this.queue[0]
	this.queue = this.queue[1:]
	return v
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
```
 入栈时间复杂度O(n),其他操作时间复杂度O(1),空间复杂度O(1)