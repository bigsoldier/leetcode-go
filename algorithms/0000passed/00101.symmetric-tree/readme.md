#### 题目
<p>给定一个二叉树，检查它是否是镜像对称的。</p>

<p>例如，二叉树&nbsp;<code>[1,2,2,3,4,4,3]</code> 是对称的。</p>

<pre>    1
   / \
  2   2
 / \ / \
3  4 4  3
</pre>

<p>但是下面这个&nbsp;<code>[1,2,2,null,3,null,3]</code> 则不是镜像对称的:</p>

<pre>    1
   / \
  2   2
   \   \
   3    3
</pre>

<p><strong>说明:</strong></p>

<p>如果你可以运用递归和迭代两种方法解决这个问题，会很加分。</p>


 #### 题解
 递归法
 ```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSym(root.Left, root.Right)
}

func isSym(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left != nil && right != nil && left.Val == right.Val {
		return isSym(left.Left, right.Right) && isSym(left.Right, right.Left)
	}
	return false
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 迭代法
 将节点插入队列，比较相邻的两个值
 ```go
func isSymmetric(root *TreeNode) bool {
	left,right := root,root
	var q []*TreeNode
	q = append(q, left,right)
	for len(q) > 0 {
		left,right = q[0],q[1]
		q = q[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		q = append(q, left.Left,right.Right)
		q = append(q, left.Right,right.Left)
	}
	return true
}
```
时间复杂度O(n),空间复杂度O(n)