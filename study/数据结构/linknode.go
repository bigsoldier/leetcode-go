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

var successor *listNode

// 反转链表前N个节点
func reverseN(head *listNode, n int) *listNode {
	if n == 1 {
		successor = head.next
		return head
	}
	last := reverseN(head.next, n-1)
	head.next.next = head
	head.next = successor
	return last
}

// 反转链表的一部分
func reverseBetween(head *listNode, m, n int) *listNode {
	if m == 1 {
		return reverseN(head, n)
	}
	head.next = reverseBetween(head, m-1, n-1)
	return head
}

func main() {

}
