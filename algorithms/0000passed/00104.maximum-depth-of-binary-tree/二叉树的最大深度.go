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

func maxDepth(root *TreeNode) int {
	var results int
	helper(&results, root, 0)
	return results
}

func helper(levels *int, node *TreeNode, level int) {
	if node == nil {
		return
	}
	if *levels <= level {
		*levels++
	}

	if node.Left != nil {
		helper(levels, node.Left, level+1)
	}
	if node.Right != nil {
		helper(levels, node.Right, level+1)
	}
}
