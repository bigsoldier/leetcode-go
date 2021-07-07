package main

// 节点的公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.left, p, q)
	right := lowestCommonAncestor(root.right, p, q)
	if left != nil && right != nil {
		return root
	} else if left == nil && right == nil {
		return nil
	} else if left == nil {
		return right
	} else {
		return left
	}
}
