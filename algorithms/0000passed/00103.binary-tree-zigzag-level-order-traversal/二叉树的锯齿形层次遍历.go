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

func zigzagLevelOrder(root *TreeNode) [][]int {
	var results [][]int
	helper(&results, root, 0)
	return results
}

func helper(levels *[][]int, node *TreeNode, level int) {
	if node == nil {
		return
	}
	if len(*levels) <= level {
		*levels = append(*levels, []int{})
	}
	if level%2 == 0 {
		(*levels)[level] = append((*levels)[level], node.Val)
	} else {
		(*levels)[level] = append([]int{node.Val}, (*levels)[level]...)
	}

	if node.Left != nil {
		helper(levels, node.Left, level+1)
	}
	if node.Right != nil {
		helper(levels, node.Right, level+1)
	}
}
