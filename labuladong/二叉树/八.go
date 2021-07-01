package main

import (
	"math"
)

// 二叉搜索子树的最大键值和
func maxSumBST(root *TreeNode) int {
	var maxSum int
	traverse4(root, &maxSum)
	return maxSum
}

func traverse4(root *TreeNode, maxSum *int) [4]int {
	if root == nil {
		return [4]int{1, math.MaxInt64, math.MinInt64, 0}
	}
	left := traverse4(root.left, maxSum)
	right := traverse4(root.right, maxSum)

	// res[0]:以root为根的二叉树是否是BST，若是1，则是
	// res[1]:以root为根的二叉树中所有节点的最小值
	// res[2]:以root为根的二叉树中所有节点的最大值
	// res[3]:以root为根的二叉树中所有节点值之和
	var res [4]int
	// 判断root是否是BST
	if left[0] == 1 && right[0] == 1 && root.val > left[2] && root.val < right[1] {
		res[0] = 1
		res[1] = min(left[1], root.val)
		res[2] = max(right[2], root.val)
		res[3] = left[3] + right[3] + root.val
		*maxSum = max(*maxSum, res[3])
	} else {
		res[0] = 0
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
