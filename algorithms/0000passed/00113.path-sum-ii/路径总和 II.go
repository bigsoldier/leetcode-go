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

func pathSum(root *TreeNode, sum int) [][]int {
	var nums [][]int
	dfs(&nums, nil, root, sum)
	return nums
}

func dfs(nums *[][]int, num []int, root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum -= root.Val
	var slice = make([]int, len(num))
	copy(slice, num)
	slice = append(slice, root.Val)
	if sum == 0 && root.Left == nil && root.Right == nil {
		*nums = append(*nums, slice)
		return
	}
	if root.Left != nil {
		dfs(nums, slice, root.Left, sum)
	}
	if root.Right != nil {
		dfs(nums, slice, root.Right, sum)
	}
}
