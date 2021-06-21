package main

// 构造最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructMaximumTree(nums, 0, len(nums)-1)
}

func constructMaximumTree(nums []int, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}
	// 找到数组中的最大值和对应的索引
	index, max := lo, nums[lo]
	for i := lo + 1; i <= hi; i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	root := &TreeNode{val: max}
	// 递归调用构造左右子树
	root.left = constructMaximumTree(nums, lo, index-1)
	root.right = constructMaximumTree(nums, index+1, hi)
	return root
}

// 从前序和中序遍历构造二叉树
func buildTree(preorder, inorder []int) *TreeNode {
	return buildtree(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func buildtree(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	// root节点对应的值就是前序遍历数组的第一个元素
	rootVal := preorder[preStart]
	// rootVal在中序遍历数组中的索引
	rootIdx := inStart
	for i := inStart + 1; i < inEnd; i++ {
		if inorder[i] == rootVal {
			rootIdx = i
			break
		}
	}

	leftSize := rootIdx - inStart
	// 构造当前的根结点
	root := &TreeNode{val: rootVal}
	root.left = buildtree(preorder, preStart+1, preStart+leftSize, inorder, inStart, rootIdx-1)
	root.right = buildtree(preorder, preStart+leftSize+1, preEnd, inorder, rootIdx+1, inEnd)
	return root
}

// 从后序和中序遍历构造二叉树
func buildTree2(inorder, postorder []int) *TreeNode {
	return buildtree2(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func buildtree2(inorder []int, inStart, inEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	rootVal := postorder[postEnd]
	rootIdx := inStart
	for i := inStart + 1; i < inEnd; i++ {
		if inorder[i] == rootVal {
			rootIdx = i
			break
		}
	}
	leftSize := rootIdx - inStart
	root := &TreeNode{val: rootVal}
	root.left = buildtree2(inorder, inStart, rootIdx-1, postorder, postStart, postStart+leftSize-1)
	root.right = buildtree2(inorder, rootIdx+1, inEnd, postorder, postStart+leftSize, postEnd-1)
	return root
}
