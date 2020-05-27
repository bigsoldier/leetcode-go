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

func deleteDuplicates(head *ListNode) *ListNode {

	var newHead = head
	for newHead != nil && newHead.Next != nil {
		if newHead.Val == newHead.Next.Val {
			newHead.Next = newHead.Next.Next
		} else {
			newHead = newHead.Next
		}
	}
	return head
}
