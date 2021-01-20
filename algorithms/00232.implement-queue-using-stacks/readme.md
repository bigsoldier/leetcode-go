#### 题目
<p>使用栈实现队列的下列操作：</p>

<ul>
	<li>push(x) -- 将一个元素放入队列的尾部。</li>
	<li>pop() -- 从队列首部移除元素。</li>
	<li>peek() -- 返回队列首部的元素。</li>
	<li>empty() -- 返回队列是否为空。</li>
</ul>

<p><strong>示例:</strong></p>

<pre>MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);  
queue.peek();  // 返回 1
queue.pop();   // 返回 1
queue.empty(); // 返回 false</pre>

<p><strong>说明:</strong></p>

<ul>
	<li>你只能使用标准的栈操作 -- 也就是只有&nbsp;<code>push to top</code>,&nbsp;<code>peek/pop from top</code>,&nbsp;<code>size</code>, 和&nbsp;<code>is empty</code>&nbsp;操作是合法的。</li>
	<li>你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。</li>
	<li>假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。</li>
</ul>


 #### 题解
 队列具有先进先出，栈具有后进先出
 1、两个栈
 ```go
type MyQueue struct {
	stack1,stack2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	for len(this.stack1) > 0 {
		this.stack2 = append(this.stack2, this.Pop())
	}
	this.stack2 = append(this.stack2, x)
	for len(this.stack2) > 0 {
		top := this.stack2[len(this.stack2)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
		this.stack1 = append(this.stack1, top)
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	top := this.stack1[len(this.stack1)-1]
	this.stack1 = this.stack1[:len(this.stack1)-1]
	return top
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stack1[len(this.stack1)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0
}
```
 时间复杂度O(n),空间复杂度O(n)