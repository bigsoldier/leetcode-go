package main

// 判断BST（二叉搜索树）的合法性
func isValidBST(root *TreeNode) bool {
	return isValidBst(root, nil, nil)
}

// 限定以root为根的子树节点必须满足max.Val>root.Val>min.Val
func isValidBst(root *TreeNode, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.val <= min.val {
		return false
	}
	if max != nil && root.val >= max.val {
		return false
	}
	// 限定左子树的最大值为root.Val，右子树的最小值为root.Val
	return isValidBst(root.left, min, root) && isValidBst(root.right, root, max)
}

// 在BST中搜索一个数
func isInBST(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	if root.val == target {
		return true
	}
	return isInBST(root.left, target) || isInBST(root.right, target)
}

func isInBST1(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	if root.val == target {
		return true
	} else if root.val > target {
		return isInBST1(root.left, target)
	} else {
		return isInBST1(root.right, target)
	}
}

// 在BST中插一个数
func insertToBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val: val}
	}
	if root.val < val {
		root.right = insertToBST(root.right, val)
	} else if root.val > val {
		root.left = insertToBST(root.left, val)
	}
	return root
}

// 在BST删除一个数
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.val == key {
		// 处理key在叶子结点上或只有一个孩子
		if root.left == nil {
			return root.right
		}
		if root.right == nil {
			return root.left
		}
		// 如果有两个孩子，那么取左子树中最大的节点或右子树中最小的节点
		minNode := getMin(root.right)
		root.val = minNode.val // fixme: 这里应该是通过链表操作来交换数据
		root.right = deleteNode(root.right, minNode.val)
	} else if root.val > key {
		root.left = deleteNode(root.left, key)
	} else {
		root.right = deleteNode(root.right, key)
	}
	return root
}

func getMin(node *TreeNode) *TreeNode {
	for node.left != nil {
		node = node.left
	}
	return node
}
