package main

import "fmt"

// 二叉搜索树中第k小的元素

var (
	smallest int // 记录结果
	rank     int // 记录当前元素的排名
)

func kthSmallest(root *TreeNode, k int) int {
	traverse1(root, k)
	return smallest
}

func traverse1(root *TreeNode, k int) {
	if root == nil {
		return
	}
	// 中序遍历
	traverse1(root.left, k)
	rank++
	if k == rank {
		smallest = root.val
		return
	}
	traverse1(root.right, k)
}

// 把二叉搜索树转换为累加树
var sum int

func convertBST(root *TreeNode) *TreeNode {
	traverse2(root)
	return root
}

func traverse2(root *TreeNode) {
	if root == nil {
		return
	}
	traverse2(root.right)
	sum += root.val
	root.val = sum
	traverse2(root.left)
}

func main() {
	root := &TreeNode{
		val: 3,
		left: &TreeNode{
			val:   1,
			left:  nil,
			right: &TreeNode{val: 2},
		},
		right: &TreeNode{val: 4},
	}
	fmt.Println(kthSmallest(root, 1))
}
