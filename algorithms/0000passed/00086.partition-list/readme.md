#### 题目
<p>给定一个链表和一个特定值<em> x</em>，对链表进行分隔，使得所有小于 <em>x</em> 的节点都在大于或等于 <em>x</em> 的节点之前。</p>

<p>你应当保留两个分区中每个节点的初始相对位置。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> head = 1-&gt;4-&gt;3-&gt;2-&gt;5-&gt;2, <em>x</em> = 3
<strong>输出:</strong> 1-&gt;2-&gt;2-&gt;4-&gt;3-&gt;5
</pre>


 #### 题解
 分成两个链表，一个存放比x小的，一个存放比x大的，然后两个链表合并
 ```go
func partition(head *ListNode, x int) *ListNode {
	var beforeHead,afterHead = new(ListNode),new(ListNode)
	var before,after *ListNode
	before = beforeHead
	after = afterHead
	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterHead.Next
	return beforeHead.Next
}

```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-05-26/008601.png)
 时间复杂度O(n),空间复杂度O(1)