#### 题目
<p>给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。</p>

<p>例如:<br>
给定二叉树:&nbsp;<code>[3,9,20,null,null,15,7]</code>,</p>

<pre>    3
   / \
  9  20
    /  \
   15   7
</pre>

<p>返回其层次遍历结果：</p>

<pre>[
  [3],
  [9,20],
  [15,7]
]
</pre>


 #### 题解
 递归
 ```go
func levelOrder(root *TreeNode) [][]int {
	var results [][]int
	helper(&results,root,0)
	return results
}

func helper(levels *[][]int,node *TreeNode, level int) {
	if node == nil {
		return 
	}
	if len(*levels) <= level {
		*levels = append(*levels, []int{})
	}
	(*levels)[level] = append((*levels)[level], node.Val)
	if node.Left != nil {
		helper(levels,node.Left,level+1)
	}
	if node.Right != nil {
		helper(levels,node.Right,level+1)
	}
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 广度优先
 ```go
func levelOrder(root *TreeNode) [][]int {
	var results [][]int

	if root == nil {
		return nil
	}
	
	// 队列
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		results = append(results, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			results[i] = append(results[i], node.Val)
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