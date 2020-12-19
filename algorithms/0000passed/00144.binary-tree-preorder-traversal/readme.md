#### 题目
<p>给定一个二叉树，返回它的&nbsp;<em>前序&nbsp;</em>遍历。</p>

<p>&nbsp;<strong>示例:</strong></p>

<pre><strong>输入:</strong> [1,null,2,3]  
   1
    \
     2
    /
   3 

<strong>输出:</strong> [1,2,3]
</pre>

<p><strong>进阶:</strong>&nbsp;递归算法很简单，你可以通过迭代算法完成吗？</p>


 #### 题解
 递归
 ```go
func preorderTraversal(root *TreeNode) []int {
	 var ret []int
	 preorder(&ret,root)
	 return ret
}

func preorder(nums *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	*nums = append(*nums, node.Val)
	preorder(nums,node.Left)
	preorder(nums,node.Right)
}
```
 时间复杂度O(n),空间复杂度O(n)

 迭代
 ```go
func preorderTraversal(root *TreeNode) []int {
	 var vals []int
	 stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		// 迭代左节点
		for root != nil {
			vals = append(vals, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	 return vals
}

```
 时间复杂度O(n),空间复杂度O(n)