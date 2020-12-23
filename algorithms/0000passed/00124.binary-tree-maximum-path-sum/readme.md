#### 题目
<p>给定一个<strong>非空</strong>二叉树，返回其最大路径和。</p>

<p>本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径<strong>至少包含一个</strong>节点，且不一定经过根节点。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [1,2,3]

       <strong>1</strong>
      <strong>/ \</strong>
     <strong>2</strong>   <strong>3</strong>

<strong>输出:</strong> 6
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [-10,9,20,null,null,15,7]

&nbsp;  -10
&nbsp; &nbsp;/ \
&nbsp; 9 &nbsp;<strong>20</strong>
&nbsp; &nbsp; <strong>/ &nbsp;\</strong>
&nbsp; &nbsp;<strong>15 &nbsp; 7</strong>

<strong>输出:</strong> 42</pre>


 #### 题解
 根据题意，路径至少包含一个节点，那么一个小的二叉树[1,2,3]能够获取到6个值，1,,2,3,1+2,1+3,1+2+3，
 把这个作为递归的基础，比较全局最大值
 ```go
func maxPathSum(root *TreeNode) int {
	var val = math.MinInt32
	pathSum(root,&val)
	return val
}

func pathSum(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}
	left := max(0,pathSum(node.Left,maxSum))
	right := max(0,pathSum(node.Right, maxSum))
	price := node.Val + left +right
	*maxSum = max(*maxSum,price)
	return node.Val + max(left,right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
 时间复杂度O(n),空间复杂度O(n)