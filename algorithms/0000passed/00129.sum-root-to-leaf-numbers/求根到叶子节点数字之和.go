package code

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	return sumNums(root, 0)
}

func sumNums(node *TreeNode, preSum int) int {
	if node == nil {
		return 0
	}
	sum := preSum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return sum
	}

	return sumNums(node.Left, sum) + sumNums(node.Right, sum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
