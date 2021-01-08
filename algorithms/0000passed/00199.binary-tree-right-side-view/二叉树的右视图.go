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

func rightSideView(root *TreeNode) []int {
	var ans []int
	dfs(root, 0, &ans)
	return ans
}

func dfs(node *TreeNode, depth int, ans *[]int) {
	if node == nil {
		return
	}
	if len(*ans) == depth {
		*ans = append(*ans, node.Val)
	}
	depth++
	dfs(node.Right, depth, ans)
	dfs(node.Left, depth, ans)
}
