package code

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	var nums []int
	inorder(root, &nums)
	swap := findTwoSwapped(nums)
	recoverT(root, 2, swap[0], swap[1])
}

func inorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inorder(root.Right, nums)
}

func findTwoSwapped(nums []int) []int {
	var x, y = -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			y = nums[i+1]
			if x == -1 {
				x = nums[i]
			}
		}
	}
	return []int{x, y}
}

func recoverT(root *TreeNode, count, x, y int) {
	if root != nil {
		if root.Val == x || root.Val == y {
			if root.Val == x {
				root.Val = y
			} else {
				root.Val = x
			}
			if count == 0 {
				return
			}
			count--
		}
		recoverT(root.Left, count, x, y)
		recoverT(root.Right, count, x, y)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
