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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var count = []int{1}
	q := []*TreeNode{root}
	for i := 0; i < len(q); i++ {
		node := q[i]
		depth := count[i]
		if node.Left == nil && node.Right == nil {
			return depth
		}
		if node.Left != nil {
			q = append(q, node.Left)
			count = append(count, depth+1)
		}
		if node.Right != nil {
			q = append(q, node.Right)
			count = append(count, depth+1)
		}
	}
	return 0
}
