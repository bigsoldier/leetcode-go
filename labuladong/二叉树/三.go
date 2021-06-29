package main

import "fmt"

// 记录所有子树以及出现的次数
var memo = map[string]int{}

// 记录重复的子树根结点
var res []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	memo = make(map[string]int)
	res = []*TreeNode{}
	traverse(root)
	return res
}

func traverse(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := traverse(root.left)
	right := traverse(root.right)
	subTree := left + "," + right + "," + string(root.val)

	freq := memo[subTree]
	if freq == 1 {
		res = append(res, root)
	}
	memo[subTree]++
	return subTree
}

func main() {
	root := &TreeNode{
		val:   1,
		left:  &TreeNode{val: 2},
		right: &TreeNode{val: 2},
	}
	ret := findDuplicateSubtrees(root)
	fmt.Println(ret)
}
