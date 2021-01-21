#### 题目
<p>给定一个二叉树，返回所有从根节点到叶子节点的路径。</p>

<p><strong>说明:</strong>&nbsp;叶子节点是指没有子节点的节点。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong>

   1
 /   \
2     3
 \
  5

<strong>输出:</strong> [&quot;1-&gt;2-&gt;5&quot;, &quot;1-&gt;3&quot;]

<strong>解释:</strong> 所有根节点到叶子节点的路径为: 1-&gt;2-&gt;5, 1-&gt;3</pre>


 #### 题解
 1、深度优先
 ```go
func binaryTreePaths(root *TreeNode) []string {
	var result []string
	dfs(root,"",&result)
	return result
}

func dfs(node *TreeNode,curr string, result *[]string) {
	if node == nil {
		return
	}
	curr += strconv.Itoa(node.Val)
	if node.Left == nil && node.Right == nil {
		*result = append(*result, curr)
	} else {
		curr += "->"
		dfs(node.Left,curr,result)
		dfs(node.Right,curr,result)
	}
}
```
 时间复杂度O(n^2^),空间复杂度O(N^2^)
 
 2、广度优先
 ```go
func binaryTreePaths(root *TreeNode) (ans []string) {
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	var paths = []string{strconv.Itoa(root.Val)}
	for len(queue) > 0 {
		q, path := queue[0], paths[0]
		queue, paths = queue[1:], paths[1:]
		if q.Left == nil && q.Right == nil {
			ans = append(ans, path)
			continue
		}
		if q.Left != nil {
			queue = append(queue, q.Left)
			paths = append(paths, path+"->"+strconv.Itoa(q.Left.Val))
		}
		if q.Right != nil {
			queue = append(queue, q.Right)
			paths = append(paths, path+"->"+strconv.Itoa(q.Right.Val))
		}
	}
	return ans
}
```
 时间复杂度O(n^2^),空间复杂度O(N^2^)