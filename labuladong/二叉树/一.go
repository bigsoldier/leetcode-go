package main

type TreeNode struct {
	val         int
	left, right *TreeNode
}

// 计算一棵二叉树共有多少个节点
func count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 自己再加上子树的节点就是整棵树的节点数
	return count(root.left) + count(root.right) + 1
}

// 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 交换root的左右子树
	tmp := root.left
	root.left = root.right
	root.right = tmp
	// 让左右子树继续翻转子节点
	invertTree(root.left)
	invertTree(root.right)
	return root
}

// 填充二叉树的右侧指针
type Node struct {
	val               int
	left, right, next *Node
}

func connect(root *Node) *Node {
	if root == nil || root.left == nil {
		return root
	}
	connectTwoNode(root.left, root.right)
	return root
}

func connectTwoNode(left, right *Node) {
	if left == nil || right == nil {
		return
	}
	left.next = right
	// 连接相同父节点的两个子节点
	connectTwoNode(left.left, left.right)
	connectTwoNode(right.left, right.right)
	// 连接跨父节点的两个子节点
	connectTwoNode(left.right, right.left)
}

// 将二叉树展开为链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.left)
	flatten(root.right)

	// 1、左右子树已经被拉平成一张链表
	left, right := root.left, root.right
	// 2、把左子树作为右子树
	root.left = nil
	root.right = left
	// 3、将原先的右子树接到当前右子树的末端
	p := root
	for p.right != nil {
		p = p.right
	}
	p.right = right
}
