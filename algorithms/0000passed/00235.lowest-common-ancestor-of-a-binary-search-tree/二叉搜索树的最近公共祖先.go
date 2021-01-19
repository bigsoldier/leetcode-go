package code

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	head := root
	for {
		if p.Val < head.Val && q.Val < head.Val {
			head = head.Left
		} else if p.Val > head.Val && q.Val > head.Val {
			head = head.Right
		} else {
			return head
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
