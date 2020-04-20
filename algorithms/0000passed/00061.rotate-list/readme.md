#### 题目
<p>给定一个链表，旋转链表，将链表每个节点向右移动&nbsp;<em>k&nbsp;</em>个位置，其中&nbsp;<em>k&nbsp;</em>是非负数。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> 1-&gt;2-&gt;3-&gt;4-&gt;5-&gt;NULL, k = 2
<strong>输出:</strong> 4-&gt;5-&gt;1-&gt;2-&gt;3-&gt;NULL
<strong>解释:</strong>
向右旋转 1 步: 5-&gt;1-&gt;2-&gt;3-&gt;4-&gt;NULL
向右旋转 2 步: 4-&gt;5-&gt;1-&gt;2-&gt;3-&gt;NULL
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> 0-&gt;1-&gt;2-&gt;NULL, k = 4
<strong>输出:</strong> <code>2-&gt;0-&gt;1-&gt;NULL</code>
<strong>解释:</strong>
向右旋转 1 步: 2-&gt;0-&gt;1-&gt;NULL
向右旋转 2 步: 1-&gt;2-&gt;0-&gt;NULL
向右旋转 3 步:&nbsp;<code>0-&gt;1-&gt;2-&gt;NULL</code>
向右旋转 4 步:&nbsp;<code>2-&gt;0-&gt;1-&gt;NULL</code></pre>


 #### 题解
 闭环
 先让链表形成一个闭环，然后移动索引，然后再断开
 ```go
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	// 形成闭环
	oldHead := head
	var length int
	for length = 1; oldHead.Next != nil; length++ {
		oldHead = oldHead.Next
	}
	oldHead.Next = head

	newTail := head
	for i := 0; i < length-k%length-1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-17/006101.png)
时间复杂度O(n),空间复杂度O(1)