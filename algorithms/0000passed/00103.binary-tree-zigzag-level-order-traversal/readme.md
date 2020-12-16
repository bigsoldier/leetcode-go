#### 题目
<p>给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。</p>

<p>例如：<br>
给定二叉树&nbsp;<code>[3,9,20,null,null,15,7]</code>,</p>

<pre>    3
   / \
  9  20
    /  \
   15   7
</pre>

<p>返回锯齿形层次遍历如下：</p>

<pre>[
  [3],
  [20,9],
  [15,7]
]
</pre>


 #### 题解
 和102类似，将数据返回的顺序改变
 递归
 ```go
func zigzagLevelOrder(root *TreeNode) [][]int {
	var results [][]int
	helper(&results, root, 0)
	return results
}

func helper(levels *[][]int, node *TreeNode, level int) {
	if node == nil {
		return
	}
	if len(*levels) <= level {
		*levels = append(*levels, []int{})
	}
	if level % 2 == 0 {
		(*levels)[level] = append((*levels)[level], node.Val)
	} else {
		(*levels)[level] = append([]int{node.Val}, (*levels)[level]... )
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
func zigzagLevelOrder(root *TreeNode) [][]int {
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
			if i%2 == 0 {
				results[i] = append(results[i], node.Val)
			} else {
				results[i] = append([]int{node.Val}, results[i]...)
			}
			
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