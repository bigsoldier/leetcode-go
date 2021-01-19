package code

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 找到前半部分链表的尾节点并反转后半部分的链表
	half := halfOfListNode(head)
	reverseList := reverse(half.Next)

	// 判断是否回文
	p1, p2 := head, reverseList
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	// 还原链表
	half.Next = reverse(reverseList)
	return result
}

func reverse(node *ListNode) *ListNode {
	var prev, curr *ListNode = nil, node
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}

func halfOfListNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

type ListNode struct {
	Val  int
	Next *ListNode
}
