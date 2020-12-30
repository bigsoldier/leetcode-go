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

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	length := 0
	// 计算链表长度
	for node := head; node != nil; node = node.Next {
		length++
	}
	dummyHead := &ListNode{Next: head}
	// 每两个子链表进行合并
	for subLength := 1; subLength <= length; subLength <<= 1 {
		prev, curr := dummyHead, dummyHead.Next
		for curr != nil {
			head1 := curr
			for i := 1; i < subLength && curr.Next != nil; i++ {
				curr = curr.Next
			}
			head2 := curr.Next
			curr.Next = nil
			curr = head2
			for i := 1; i < subLength && curr != nil; i++ {
				curr = curr.Next
			}

			var next *ListNode
			if curr != nil {
				next = curr.Next
				curr.Next = nil
			}

			prev.Next = merge(head1, head2)
			for prev.Next != nil {
				prev = prev.Next
			}
			curr = next
		}
	}
	return dummyHead.Next
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tmp, tmp1, tmp2 := dummyHead, head1, head2
	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val <= tmp2.Val {
			tmp.Next = tmp1
			tmp1 = tmp1.Next
		} else {
			tmp.Next = tmp2
			tmp2 = tmp2.Next
		}
		tmp = tmp.Next
	}
	if tmp1 != nil {
		tmp.Next = tmp1
	} else if tmp2 != nil {
		tmp.Next = tmp2
	}
	return dummyHead.Next
}
