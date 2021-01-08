#### 题目
<p>给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong>&nbsp;[1,2,3,null,5,null,4]
<strong>输出:</strong>&nbsp;[1, 3, 4]
<strong>解释:
</strong>
   1            &lt;---
 /   \
2     3         &lt;---
 \     \
  5     4       &lt;---
</pre>


 #### 题解
 1、深度优先
 ```go
func rightSideView(root *TreeNode) []int {
	var ans []int
	dfs(root,0,&ans)
	return ans
}

func dfs(node *TreeNode, depth int, ans *[]int) {
	if node == nil {
		return
	}
	if len(*ans) == depth {
		*ans = append(*ans, node.Val)
	}
	depth++
	dfs(node.Right,depth,ans)
	dfs(node.Left,depth,ans)
}
```
 时间复杂度O(n),空间复杂度O(n)

 2、广度优先
 ```go
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue)>0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			q := queue[0]
			queue = queue[1:]
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
			if i == n-1 {
				ans = append(ans, q.Val)
			}
		}
	}
	return
}
```
 时间复杂度O(n),空间复杂度O(n)
