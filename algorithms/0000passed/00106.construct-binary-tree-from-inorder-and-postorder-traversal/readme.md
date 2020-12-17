#### 题目
<p>根据一棵树的中序遍历与后序遍历构造二叉树。</p>

<p><strong>注意:</strong><br>
你可以假设树中没有重复的元素。</p>

<p>例如，给出</p>

<pre>中序遍历 inorder =&nbsp;[9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]</pre>

<p>返回如下的二叉树：</p>

<pre>    3
   / \
  9  20
    /  \
   15   7
</pre>


 #### 题解
 中序遍历：左->根->右
 后序遍历：左->右->根
 
 那么我们可以大胆假设最后一个元素就是根节点，
 以后序遍历为基础，找出根节点在中序遍历的位置
 
 递归
 ```go
func buildTree(inorder []int, postorder []int) *TreeNode {
	inPos := make(map[int]int)
	for i, val := range inorder {
		inPos[val] = i
	}
	return build(postorder, 0, len(postorder)-1, 0, inPos)
}
func build(postorder []int, postStart, postEnd int, inStart int, inPos map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}
	// 根节点
	root := &TreeNode{Val: postorder[postEnd]}
	// 根节点在中序遍历的位置
	rootIdx := inPos[postorder[postEnd]]
	leftLen := rootIdx - inStart
	root.Left = build(postorder,postStart,postStart+leftLen-1,inStart,inPos)
	root.Right = build(postorder,postStart+leftLen,postEnd-1,rootIdx+1,inPos)
	return root
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 迭代
 
 - 初始栈中存放根节点(后序遍历的最后一个节点),指针指向中序遍历最后一个节点
 - 依次枚举后序遍历除第一个节点以外的每个节点，如果index恰好指向栈顶元素，那么我们不断弹出栈顶元素并将index左移，
 并将当前节点作为最后一个弹出节点的左儿子；如果index与栈顶元素不等，就将当前节点作为栈顶节点的右儿子
 ```go
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	stack := []*TreeNode{root}
	inorderIdx := len(inorder) - 1

	for i := len(postorder)-2; i >= 0; i-- {
		postorderVal := postorder[i]
		// 栈顶元素
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIdx] {
			node.Right = &TreeNode{Val: postorderVal}
			stack = append(stack, node.Right)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIdx] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIdx--
			}
			node.Left = &TreeNode{Val: postorderVal}
			stack = append(stack, node.Left)
		}
	}
	return root
}
```
 时间复杂度O(n),空间复杂度O(n)