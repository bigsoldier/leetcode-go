package code

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSym(root.Left, root.Right)
}

func isSym(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left != nil && right != nil && left.Val == right.Val {
		return isSym(left.Left, right.Right) && isSym(left.Right, right.Left)
	}
	return false
}
