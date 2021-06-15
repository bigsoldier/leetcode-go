package main

/**
迭代方法实现
*/
// 逆转链表
func reverse1(head *listNode) *listNode {
	var prev, curr, next *listNode = nil, head, head
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 反转区间[a,b)的元素
func reverseBetween1(a, b *listNode) *listNode {
	var prev, curr, next *listNode = nil, a, a
	for curr != b { // 改一下对应的终止条件
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}

// k个一组反转链表
func reverseKGroup(head *listNode, k int) *listNode {
	if head == nil {
		return head
	}
	// 区间[a,b)包含k个待反转的元素
	var a, b *listNode = head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.next
	}
	newHead := reverseBetween1(a, b)
	a.next = reverseKGroup(b, k)
	return newHead
}
