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

func postorderTraversal(root *TreeNode) []int {
	var ret []int
	postorder(&ret, root)
	return ret
}

func postorder(nums *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	postorder(nums, node.Left)
	postorder(nums, node.Right)
	*nums = append(*nums, node.Val)
}
