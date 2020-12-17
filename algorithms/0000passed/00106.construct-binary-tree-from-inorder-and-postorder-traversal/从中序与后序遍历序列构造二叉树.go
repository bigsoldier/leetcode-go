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

func buildTree(inorder []int, postorder []int) *TreeNode {
	inPos := make(map[int]int)
	for i, val := range inorder {
		inPos[val] = i
	}
	return build(postorder, 0, len(postorder)-1, 0, inPos)
}
func build(postorder []int, postStart, postEnd int, inStart int, inPos map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}
	// 根节点
	root := &TreeNode{Val: postorder[postEnd]}
	// 根节点在中序遍历的位置
	rootIdx := inPos[postorder[postEnd]]
	leftLen := rootIdx - inStart
	root.Left = build(postorder, postStart, postStart+leftLen-1, inStart, inPos)
	root.Right = build(postorder, postStart+leftLen, postEnd-1, rootIdx+1, inPos)
	return root
}
