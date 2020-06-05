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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	// move the two points util they reach the proper starting point
	var cur = head
	var pre *ListNode
	for m > 1 {
		pre = cur
		cur = cur.Next
		m--
		n--
	}
	// tail 是反转后链表的尾部
	con, tail := pre, cur
	var third *ListNode
	for n > 0 {
		third = cur.Next
		cur.Next = pre
		pre = cur
		cur = third
		n--
	}

	if con != nil {
		con.Next = pre
	} else {
		head = pre
	}
	tail.Next = cur
	return head
}
