package code

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

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
