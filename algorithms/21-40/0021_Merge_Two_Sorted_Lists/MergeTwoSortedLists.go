/**
 *Created by Xie Jian on 2020/2/14 12:59
 */
package _021_Merge_Two_Sorted_Lists

//Merge two sorted linked lists and return it as a new list. The new list should
// be made by splicing together the nodes of the first two lists.
//
// Example:
//
//Input: 1->2->4, 1->3->4
//Output: 1->1->2->3->4->4
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var newList = new(ListNode)
	pre := newList
	for l1 != nil && l2 != nil { // 循环到链表末尾
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l1 != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}
	return newList.Next
}

//leetcode submit region end(Prohibit modification and deletion)
