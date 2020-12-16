#### 题目
<p>根据一棵树的前序遍历与中序遍历构造二叉树。</p>

<p><strong>注意:</strong><br>
你可以假设树中没有重复的元素。</p>

<p>例如，给出</p>

<pre>前序遍历 preorder =&nbsp;[3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]</pre>

<p>返回如下的二叉树：</p>

<pre>    3
   / \
  9  20
    /  \
   15   7</pre>


 #### 题解
 考虑一下前序遍历和中序遍历的定义
 
 前: 3  9  20 15 7
 中: 9  3  15 20 7
 
 我们很容易知道3是根节点，通过在中序遍历中找到3的位置，那么我们知道在3前面的是3的左子树，在3右面的是3的右子树。
 则在前序排列中，我们能找出9是左子树的一部分，10 15 7是右子树的一部分。
 然后再根据20是右子树的根节点，继续上面的过程
 
 递归
 ```go
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	var i int
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	//fmt.Println(len(inorder[:i]),i)
	root.Left = buildTree(preorder[1:i+1],inorder[:i])
	root.Right = buildTree(preorder[i+1:],inorder[i+1:])
	return root
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 迭代
 对于前序遍历，任意两个连续的节点u和v，根据流程，有以下关系：
 - v是u左儿子
 - u没有左儿子，并且v是u的某个祖先节点（包括u）的右儿子
 
 用一个栈`stack`维护[当前节点的所有还未考虑右儿子的祖先节点],栈顶是当前节点，同时用一个指针`index`指向中序遍历的某个位置。
 `index`对应[当前节点不断向左走达到的最终节点]。
 
 ```go
 func buildTree(preorder []int, inorder []int) *TreeNode {
 	if len(preorder) == 0 {
 		return nil
 	}
 	root := &TreeNode{Val: preorder[0]}
 
 	var stack []*TreeNode
 	var inorderIdx int
 	stack = append(stack, root)
 	for i := 1; i < len(preorder); i++ {
 		preorderVal := preorder[i]
 		node := stack[len(stack)-1]
 		if node.Val != inorder[inorderIdx] {
 			node.Left = &TreeNode{Val: preorderVal}
 			stack = append(stack, node.Left)
 		} else {
 			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIdx] {
 				node = stack[len(stack)-1]
 				stack = stack[:len(stack)-1]
 				inorderIdx++
 			}
 			node.Right = &TreeNode{Val: preorderVal}
 			stack = append(stack, node.Right)
 		}
 	}
 	return root
 }
```