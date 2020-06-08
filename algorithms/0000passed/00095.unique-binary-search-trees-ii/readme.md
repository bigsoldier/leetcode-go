#### 题目
<p>给定一个整数 <em>n</em>，生成所有由 1 ...&nbsp;<em>n</em> 为节点所组成的<strong>二叉搜索树</strong>。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 3
<strong>输出:</strong>
[
&nbsp; [1,null,3,2],
&nbsp; [3,2,null,1],
&nbsp; [3,1,null,null,2],
&nbsp; [2,1,3],
&nbsp; [1,null,2,null,3]
]
<strong>解释:</strong>
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
</pre>


 #### 题解
 1、递归
 ```go
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTree(1,n)
}

func generateTree(start,end int) []*TreeNode {
	var allTrees []*TreeNode
	if start > end {
		allTrees = append(allTrees, nil)
		return allTrees
	}

	for i := start; i <= end; i++ {
		left := generateTree(start,i-1)
		right := generateTree(i+1,end)
		for _, leftNode := range left {
			for _, RightNode := range right {
				currentTree := &TreeNode{
					Val:   i,
					Left:  leftNode,
					Right: RightNode,
				}
				allTrees = append(allTrees, currentTree)
			}
		}
	}
	return allTrees
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-06-06/009501.png)