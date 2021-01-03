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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	var hasLinkedToB, hasLinkedToA bool
	for l1 != nil && l2 != nil {
		if l1 == l2 {
			return l2
		}
		l1, l2 = l1.Next, l2.Next
		if l1 == nil && !hasLinkedToB {
			l1 = headB
			hasLinkedToB = true
		}
		if l2 == nil && !hasLinkedToA {
			l2 = headA
			hasLinkedToA = true
		}
	}
	return nil
}
