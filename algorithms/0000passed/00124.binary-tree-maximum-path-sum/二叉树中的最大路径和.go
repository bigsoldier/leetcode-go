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

func maxPathSum(root *TreeNode) int {
	var sum int
	pathSum(root, &sum)
	return sum
}

func pathSum(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}
	left := max(0, pathSum(node.Left, maxSum))
	right := max(0, pathSum(node.Right, maxSum))
	price := node.Val + left + right
	*maxSum = max(*maxSum, price)
	return node.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
