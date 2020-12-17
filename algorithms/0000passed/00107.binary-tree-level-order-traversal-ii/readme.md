#### 题目
<p>给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）</p>

<p>例如：<br>
给定二叉树 <code>[3,9,20,null,null,15,7]</code>,</p>

<pre>    3
   / \
  9  20
    /  \
   15   7
</pre>

<p>返回其自底向上的层次遍历为：</p>

<pre>[
  [15,7],
  [9,20],
  [3]
]
</pre>


 #### 题解
 广度优先
 ```go
func levelOrderBottom(root *TreeNode) [][]int {
	var results [][]int

	if root == nil {
		return nil
	}

	q := []*TreeNode{root}
	for len(q)>0 {
		results = append([][]int{{}},results...)
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			results[0] = append(results[0], node.Val)
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