#### 题目
<p>删除链表中等于给定值&nbsp;<strong><em>val&nbsp;</em></strong>的所有节点。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 1-&gt;2-&gt;6-&gt;3-&gt;4-&gt;5-&gt;6, <em><strong>val</strong></em> = 6
<strong>输出:</strong> 1-&gt;2-&gt;3-&gt;4-&gt;5
</pre>


 #### 题解
 使用哨兵节点作为伪头
 ```go
func removeElements(head *ListNode, val int) *ListNode {
	node := new(ListNode)
	node.Next = head
	prev,curr := node,head
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return node.Next
}
```
 时间复杂度O(n),空间复杂度O(1)