#### 题目
<p>给定一个二叉树，找出其最大深度。</p>

<p>二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。</p>

<p><strong>说明:</strong>&nbsp;叶子节点是指没有子节点的节点。</p>

<p><strong>示例：</strong><br>
给定二叉树 <code>[3,9,20,null,null,15,7]</code>，</p>

<pre>    3
   / \
  9  20
    /  \
   15   7</pre>

<p>返回它的最大深度&nbsp;3 。</p>


 #### 题解
 和102、103类似
 
 递归
 ```go
func maxDepth(root *TreeNode) int {
	var results int
	helper(&results, root, 0)
	return results
}

func helper(levels *int, node *TreeNode, level int) {
	if node == nil {
		return
	}
	if *levels <= level {
		*levels++
	}

	if node.Left != nil {
		helper(levels, node.Left, level+1)
	}
	if node.Right != nil {
		helper(levels, node.Right, level+1)
	}
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 广度优先
 ```go
func maxDepth(root *TreeNode) int {
	var results int

	if root == nil {
		return 0
	}

	// 队列
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		results++
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]

			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return results
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 更简单的方法
 
 假设我们知道左子树和右子树的最大深度为l和r，那么该二叉树的最大深度为`max(l,r)+1`
 而左子树和右子树的深度可以以同样的方式进行计算
 ```go
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left),maxDepth(root.Right))+1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
```
 