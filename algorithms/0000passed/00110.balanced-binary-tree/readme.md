#### 题目
<p>给定一个二叉树，判断它是否是高度平衡的二叉树。</p>

<p>本题中，一棵高度平衡二叉树定义为：</p>

<blockquote>
<p>一个二叉树<em>每个节点&nbsp;</em>的左右两个子树的高度差的绝对值不超过1。</p>
</blockquote>

<p><strong>示例 1:</strong></p>

<p>给定二叉树 <code>[3,9,20,null,null,15,7]</code></p>

<pre>    3
   / \
  9  20
    /  \
   15   7</pre>

<p>返回 <code>true</code> 。<br>
<br>
<strong>示例 2:</strong></p>

<p>给定二叉树 <code>[1,2,2,3,3,null,null,4,4]</code></p>

<pre>       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
</pre>

<p>返回&nbsp;<code>false</code> 。</p>

<p>&nbsp;</p>


 #### 题解
 自顶向下
 
 关注于节点的左子树是否是平衡二叉树，右子树是否是平衡二叉树，且两者高度小于等于1
 ```go
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left) - height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right) 
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height(node.Left), height(node.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
 时间复杂度O(n^2^),空间复杂度O(n)
 
 自底向上
 
 对于同一个接单，height会被重复调用。
 所以，先递归判断当前节点的子树是否平衡，如果平衡，返回其高度
 ```go
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := height(node.Left)
	rightHeight := height(node.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(rightHeight - leftHeight) > 1 {
		return -1
	}
	return max(leftHeight,rightHeight)+1
}
```