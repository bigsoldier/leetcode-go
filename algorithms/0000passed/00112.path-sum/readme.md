#### 题目
<p>给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。</p>

<p><strong>说明:</strong>&nbsp;叶子节点是指没有子节点的节点。</p>

<p><strong>示例:</strong>&nbsp;<br>
给定如下二叉树，以及目标和 <code>sum = 22</code>，</p>

<pre>              <strong>5</strong>
             / \
            <strong>4 </strong>  8
           /   / \
          <strong>11 </strong> 13  4
         /  \      \
        7    <strong>2</strong>      1
</pre>

<p>返回 <code>true</code>, 因为存在目标和为 22 的根节点到叶子节点的路径 <code>5-&gt;4-&gt;11-&gt;2</code>。</p>


 #### 题解
 广度优先
 ```go
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	q := []*TreeNode{root}
	var num = []int{root.Val}
	for len(q) != 0 {
		node,tmp := q[0],num[0]
		q,num = q[1:],num[1:]
		if node.Left == nil && node.Right == nil {
			if tmp == sum {
				return true
			}
			continue
		}
		if node.Left != nil {
			q = append(q, node.Left)
			num = append(num, node.Left.Val+tmp)
		}
		if node.Right != nil {
			q = append(q, node.Right)
			num = append(num, node.Right.Val+tmp)
		}
	}
	return false
}
```
 时间复杂度O(n),空间复杂度O(n) 
 
 递归
 ```go
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum==root.Val
	}
	return hasPathSum(root.Left,sum-root.Val)||hasPathSum(root.Right,sum-root.Val)
}
```
 时间复杂度O(n),空间复杂度O(n) 