#### 题目
<p>给定一个二叉树，返回它的<em>中序&nbsp;</em>遍历。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> [1,null,2,3]
   1
    \
     2
    /
   3

<strong>输出:</strong> [1,3,2]</pre>

<p><strong>进阶:</strong>&nbsp;递归算法很简单，你可以通过迭代算法完成吗？</p>


 #### 题解
 1、递归
 ```go
 func inorderTraversal(root *TreeNode) []int {
 	var result []int
 	traversal(root,&result)
 	return result
 }
 
 func traversal(node *TreeNode, result *[]int) {
 	if node != nil {
 		if node.Left != nil {
 			traversal(node.Left,result)
 		}
 		*result = append(*result, node.Val)
 		if node.Right != nil {
 			traversal(node.Right,result)
 		}
 	}
 }

```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-06-04/009401.png)
 时间复杂度O(n),空间复杂度最坏O(n),最好O(logn)

 2、栈
 ```go
func inorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	curr := root
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, curr.Val)
		curr = curr.Right
	}
	return result
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-06-04/009402.png)
 时间复杂度O(n),空间复杂度O(n)