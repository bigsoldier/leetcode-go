package main

import "math"

// 普通二叉树结点数
func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes1(root.left) + countNodes1(root.right)
}

// 满二叉树结点数
func countNodes2(root *TreeNode) int {
	var height int
	for root != nil {
		root = root.left
		height++
	}
	return int(math.Pow(2, float64(height))) - 1
}

// 完全二叉树结点数
func countNodes3(root *TreeNode) int {
	l, r := root, root
	var hl, hr int
	for l != nil {
		l = l.left
		hl++
	}
	for r != nil {
		r = r.right
		hr++
	}
	// 如果高度相同，按满二叉树计算
	if hl == hr {
		return int(math.Pow(2, float64(hr))) - 1
	}
	// 如果高度不等，按普通二叉树计算
	return 1 + countNodes3(root.left) + countNodes3(root.right)
}
