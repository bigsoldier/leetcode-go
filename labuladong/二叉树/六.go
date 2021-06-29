package main

// 给定正整数n，计算有多少种BST结构
func numTrees(n int) int {
	return count1(1, n)
}

// 闭区间[lo, hi]的数字能组成的BST数量
func count1(lo, hi int) int {
	if lo > hi { // 此时对应的是空节点，也是一种情况
		return 1
	}
	var res int
	for i := lo; i <= hi; i++ {
		left := count1(lo, i-1)
		right := count1(i+1, hi)
		res += left * right
	}
	return res
}

var memo1 [][]int

func numTrees1(n int) int {
	memo1 = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo1[i] = make([]int, n)
	}
	return count2(1, n)
}

func count2(lo, hi int) int {
	if lo > hi {
		return 1
	}
	if memo1[lo][hi] != 0 {
		return memo1[lo][hi]
	}
	var res int
	for i := lo; i <= hi; i++ {
		left := count2(lo, i-1)
		right := count2(i+1, hi)
		res += left * right
	}
	memo1[lo][hi] = res
	return res
}

func generateTrees(n int) []*TreeNode {
	return build(1, n)
}

func build(lo, hi int) []*TreeNode {
	var res []*TreeNode
	if lo > hi {
		res = append(res, &TreeNode{})
		return res
	}
	for i := lo; i <= hi; i++ {
		left := build(lo, i-1)
		right := build(i+1, hi)
		for _, leftNode := range left {
			for _, rightNode := range right {
				root := &TreeNode{val: i, left: leftNode, right: rightNode}
				res = append(res, root)
			}
		}
	}
	return res
}
