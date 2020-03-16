/**
 *Created by Xie Jian on 2020/2/12 15:35
 */
package _019_Remove_Nth_Node_From_End_Of_List

//Given a linked list, remove the n-th node from the end of list and return its head.
//
// Example:
//
//
//Given linked list: 1->2->3->4->5, and n = 2.
//
//After removing the second node from the end, the linked list becomes 1->2->3->5.
//
//
// Note:
//
// Given n will always be valid.
//
// Follow up:
//
// Could you do this in one pass?
// Related Topics Linked List Two Pointers

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newList := new(ListNode)
	newList.Next = head
	first, second := newList, newList
	for i := 0; i <= n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return newList.Next
}

//leetcode submit region end(Prohibit modification and deletion)
