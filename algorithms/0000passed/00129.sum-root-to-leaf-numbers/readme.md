#### 题目
<p>给定一个二叉树，它的每个结点都存放一个&nbsp;<code>0-9</code>&nbsp;的数字，每条从根到叶子节点的路径都代表一个数字。</p>

<p>例如，从根到叶子节点路径 <code>1-&gt;2-&gt;3</code> 代表数字 <code>123</code>。</p>

<p>计算从根到叶子节点生成的所有数字之和。</p>

<p><strong>说明:</strong>&nbsp;叶子节点是指没有子节点的节点。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [1,2,3]
    1
   / \
  2   3
<strong>输出:</strong> 25
<strong>解释:</strong>
从根到叶子节点路径 <code>1-&gt;2</code> 代表数字 <code>12</code>.
从根到叶子节点路径 <code>1-&gt;3</code> 代表数字 <code>13</code>.
因此，数字总和 = 12 + 13 = <code>25</code>.</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> [4,9,0,5,1]
    4
   / \
  9   0
&nbsp;/ \
5   1
<strong>输出:</strong> 1026
<strong>解释:</strong>
从根到叶子节点路径 <code>4-&gt;9-&gt;5</code> 代表数字 495.
从根到叶子节点路径 <code>4-&gt;9-&gt;1</code> 代表数字 491.
从根到叶子节点路径 <code>4-&gt;0</code> 代表数字 40.
因此，数字总和 = 495 + 491 + 40 = <code>1026</code>.</pre>


 #### 题解
 递归
 ```go
func sumNumbers(root *TreeNode) int {
	return sumNums(root, 0)
}

func sumNums(node *TreeNode, preSum int) int {
	if node == nil {
		return 0
	}
	sum := preSum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return sum
	}

	return sumNums(node.Left, sum) + sumNums(node.Right, sum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 循环
 ```go
func sumNumbers(root *TreeNode) (sum int) {
	if root == nil {
		return 0
	}
	q := []pair{{root,root.Val}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		left,right,num := p.node.Left,p.node.Right,p.num
		if left==nil&&right==nil {
			sum+=num
		} else {
			if left != nil {
				q = append(q, pair{left,num*10+left.Val})
			}
			if right != nil {
				q = append(q, pair{right,num*10+right.Val})
			}
		}
	}
	return 
}
```
 时间复杂度O(n),空间复杂度O(n)