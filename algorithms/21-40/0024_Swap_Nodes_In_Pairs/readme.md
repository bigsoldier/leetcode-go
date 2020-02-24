#### 题目
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
**你不能只是单纯的改变节点内部的值**，而是需要实际的进行节点交换。
**示例：**
```
给定 1->2->3->4，应返回 2->1->4->3
```

#### 题解
1、递归
每次选取奇数位的链表元素，交换当前元素和后一位元素。
```go
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-18/002401.png)
时间复杂度O(n)，空间复杂度O(1)

2、非递归
```go
func swapPairs(head *ListNode) *ListNode {
	var pre = new(ListNode)
	pre.Next = head
	temp := pre
	for temp.Next != nil && temp.Next.Next != nil {
		start := temp.Next
		end := temp.Next.Next
		temp.Next = end
		start.Next = end.Next
		end.Next = start
		temp = start	// 移动
	}
	return pre.Next
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-18/002402.png)
时间复杂度O(n)，空间复杂度O(1)