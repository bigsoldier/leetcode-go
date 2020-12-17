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

func levelOrderBottom(root *TreeNode) [][]int {
	var results [][]int

	if root == nil {
		return nil
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		results = append([][]int{{}}, results...)
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			results[0] = append(results[0], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return results
}
