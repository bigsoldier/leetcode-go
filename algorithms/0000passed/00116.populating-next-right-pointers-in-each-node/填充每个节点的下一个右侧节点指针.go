package code

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	// 每次循环都从该节点的最左侧节点开始
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		// 通过next遍历这一层节点，为下一层更新next节点
		for node := leftmost; node != nil; node = node.Next {
			// 左节点指向右节点
			node.Left.Next = node.Right
			// 右节点指向下一节点的左节点
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
