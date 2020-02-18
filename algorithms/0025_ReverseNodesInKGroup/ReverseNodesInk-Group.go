/**
 *Created by Xie Jian on 2020/2/14 13:03
 */
package _025_ReverseNodesInKGroup

//Given a linked list, reverse the nodes of a linked list k at a time and return
// its modified list.
//
// k is a positive integer and is less than or equal to the length of the linked
// list. If the number of nodes is not a multiple of k then left-out nodes in the
//end should remain as it is.
//
//
//
//
// Example:
//
// Given this linked list: 1->2->3->4->5
//
// For k = 2, you should return: 2->1->4->3->5
//
// For k = 3, you should return: 3->2->1->4->5
//
// Note:
//
//
// Only constant extra memory is allowed.
// You may not alter the values in the list's nodes, only nodes itself may be ch
//anged.
//
// Related Topics Linked List

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

func reverseKGroup(head *ListNode, k int) *ListNode {
	var node = head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(start, end *ListNode) *ListNode {
	pre := end
	for start != end {
		tmp := start.Next
		start.Next = pre
		pre = start
		start = tmp // 移动节点
	}
	return pre
}

//leetcode submit region end(Prohibit modification and deletion)
