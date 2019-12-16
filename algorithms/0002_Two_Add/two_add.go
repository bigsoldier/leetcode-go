package leetcode

//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Example:
//
//
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.
//
// Related Topics Linked List Math

//leetcode submit region begin(Prohibit modification and deletion)
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	head := &ListNode{Val: 0, Next: nil}
	current := head
	carry := 0

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}

		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}

		current.Next = &ListNode{Val: (x + y + carry) % 10, Next: nil}
		current = current.Next
		carry = (x + y + carry) / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry % 10, Next: nil}
	}
	return head.Next
}

// 2 个逆序的链表，要求从低位开始相加，得出结果也逆序输出，返回值是逆序结果链表的头结点。
