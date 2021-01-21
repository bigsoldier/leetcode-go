package code

import "strconv"

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

func binaryTreePaths(root *TreeNode) (ans []string) {
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	var paths = []string{strconv.Itoa(root.Val)}
	for len(queue) > 0 {
		q, path := queue[0], paths[0]
		queue, paths = queue[1:], paths[1:]
		if q.Left == nil && q.Right == nil {
			ans = append(ans, path)
			continue
		}
		if q.Left != nil {
			queue = append(queue, q.Left)
			paths = append(paths, path+"->"+strconv.Itoa(q.Left.Val))
		}
		if q.Right != nil {
			queue = append(queue, q.Right)
			paths = append(paths, path+"->"+strconv.Itoa(q.Right.Val))
		}
	}
	return ans
}
