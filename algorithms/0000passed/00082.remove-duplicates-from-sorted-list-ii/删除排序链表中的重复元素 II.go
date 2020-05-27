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
	var newHead = new(ListNode)
	newHead.Next = head
	p := newHead
	for p.Next != nil && p.Next.Next != nil {
		var tag bool

		// 移除重复元素
		for p.Next.Next != nil && p.Next.Val == p.Next.Next.Val {
			tag = true
			p.Next.Next = p.Next.Next.Next
		}

		// 移除最后的重复元素
		if tag {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return newHead.Next
}
