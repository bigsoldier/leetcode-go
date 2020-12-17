package code

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var global *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	global = head
	length := getLen(head)
	return buildTree(0, length-1)
}

func getLen(head *ListNode) int {
	var ret int
	for ; head != nil; head = head.Next {
		ret++
	}
	return ret
}

func buildTree(left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right + 1) / 2
	root := &TreeNode{}
	root.Left = buildTree(left, mid-1)
	root.Val = global.Val
	global = global.Next
	root.Right = buildTree(mid+1, right)
	return root
}
