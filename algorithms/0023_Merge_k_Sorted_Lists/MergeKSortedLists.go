/**
 *Created by Xie Jian on 2020/2/14 13:01
 */
package _023_Merge_k_Sorted_Lists

//Merge k sorted linked lists and return it as one sorted list. Analyze and desc
//ribe its complexity.
//
// Example:
//
//
//Input:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//Output: 1->1->2->3->4->4->5->6
//
// Related Topics Linked List Divide and Conquer Heap

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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := left + (right-left)/2
	l1 := merge(lists, left, mid)
	l2 := merge(lists, mid+1, right)
	return mergeList(l1, l2)
}

func mergeList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeList(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeList(l1, l2.Next)
		return l2
	}
}

//leetcode submit region end(Prohibit modification and deletion)
