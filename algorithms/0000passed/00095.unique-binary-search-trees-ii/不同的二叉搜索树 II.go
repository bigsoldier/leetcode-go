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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTree(1, n)
}

func generateTree(start, end int) []*TreeNode {
	var allTrees []*TreeNode
	if start > end {
		allTrees = append(allTrees, nil)
		return allTrees
	}

	for i := start; i <= end; i++ {
		left := generateTree(start, i-1)
		right := generateTree(i+1, end)
		for _, leftNode := range left {
			for _, RightNode := range right {
				currentTree := &TreeNode{
					Val:   i,
					Left:  leftNode,
					Right: RightNode,
				}
				allTrees = append(allTrees, currentTree)
			}
		}
	}
	return allTrees
}
