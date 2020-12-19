#### 题目
<p>给定一个二叉树，返回它的 <em>后序&nbsp;</em>遍历。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> [1,null,2,3]  
   1
    \
     2
    /
   3 

<strong>输出:</strong> [3,2,1]</pre>

<p><strong>进阶:</strong>&nbsp;递归算法很简单，你可以通过迭代算法完成吗？</p>


 #### 题解
 递归
 ```go
func postorderTraversal(root *TreeNode) []int {
	var ret []int
	postorder(&ret,root)
	return ret
}

func postorder(nums *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	postorder(nums,node.Left)
	postorder(nums,node.Right)
	*nums = append(*nums, node.Val)
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 迭代
 ```go
func postorderTraversal(root *TreeNode) []int {
	var ret []int
	var stack []*TreeNode
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			ret = append(ret, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return ret
}
```