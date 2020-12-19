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

func preorderTraversal(root *TreeNode) []int {
	var ret []int
	preorder(&ret, root)
	return ret
}

func preorder(nums *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	*nums = append(*nums, node.Val)
	preorder(nums, node.Left)
	preorder(nums, node.Right)
}
