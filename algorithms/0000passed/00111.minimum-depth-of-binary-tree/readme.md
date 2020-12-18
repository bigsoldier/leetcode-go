#### 题目
<p>给定一个二叉树，找出其最小深度。</p>

<p>最小深度是从根节点到最近叶子节点的最短路径上的节点数量。</p>

<p><strong>说明:</strong>&nbsp;叶子节点是指没有子节点的节点。</p>

<p><strong>示例:</strong></p>

<p>给定二叉树&nbsp;<code>[3,9,20,null,null,15,7]</code>,</p>

<pre>    3
   / \
  9  20
    /  \
   15   7</pre>

<p>返回它的最小深度 &nbsp;2.</p>


 #### 题解
 递归
 ```go
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return min(minDepth(root.Left),minDepth(root.Right))+1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

```
时间复杂度O(n),空间复杂度O(logn)

 迭代
 ```go 
 func minDepth(root *TreeNode) int {
 	if root == nil {
 		return 0
 	}
 	var count = []int{1}
 	q := []*TreeNode{root}
 	for i:=0;i < len(q);i++ {
 		node := q[i]
 		depth := count[i]
 		if node.Left == nil && node.Right == nil {
 			return depth
 		}
 		if node.Left != nil {
 			q = append(q, node.Left)
 			count = append(count, depth+1)
 		}
 		if node.Right != nil {
 			q = append(q, node.Right)
 			count = append(count, depth+1)
 		}
 	}
 	return 0
 }
```
时间复杂度O(n),空间复杂度O(n)