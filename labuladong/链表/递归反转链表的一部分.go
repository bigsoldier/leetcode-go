package main

type listNode struct {
	val  int
	next *listNode
}

// 反转链表
func reverse(head *listNode) *listNode {
	if head == nil {
		return head
	}
	last := reverse(head.next)
	head.next.next = head
	head.next = nil
	return last
}

var successor *listNode // 后驱节点

// 反转链表前N个节点

// 反转以head为起点的n个节点，返回新的头节点
func reverseN(head *listNode, n int) *listNode {
	if n == 1 {
		// 记录第n+1个节点
		successor = head.next
		return head
	}
	// 以head.next为起点，需要反转前n-1个节点
	last := reverseN(head.next, n-1)
	head.next.next = head
	// 让反转之后的head节点和后面的节点连起来
	head.next = successor
	return last
}

// 反转链表的一部分
func reverseBetween(head *listNode, m, n int) *listNode {
	if m == 1 {
		return reverseN(head, n)
	}
	// 前进到反转的起点触发 reverseN
	head.next = reverseBetween(head, m-1, n-1)
	return head
}
