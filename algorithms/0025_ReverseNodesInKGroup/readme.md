#### 题目
给定一个链表，每k个节点一组进行反转。
k是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
**示例：**
```
给定这个链表，1->2->3->4->5
当k=2时，应当返回：2->1->4->3->5
当k=3时，应当返回：3->2->1->4->5
```
**说明：**
- 你的算法只能使用常熟的额外空间
- **你不能只是单纯的改变节点内部的值**，而是需要实际的进行节点交换。

#### 题解
与24类似
```go
func reverseKGroup(head *ListNode, k int) *ListNode {
	var node = head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head,node)
	head.Next = reverseKGroup(node,k)
	return newHead
}

func reverse(start,end *ListNode) *ListNode {
	pre := end
	for start != end {
		tmp := start.Next
		start.Next = pre
		pre = start
		start = tmp	// 移动节点
	}
	return pre
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-18/002501.png)