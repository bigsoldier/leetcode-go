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
 将一个大的二叉树分成一个最小单元二叉树，[1,2,3]能获取到的值为1,2,3,4,5,6.
 总的来说就是它的左子树和右子树及自身的比较
 ```go
func maxPathSum(root *TreeNode) int {
	var sum int
	pathSum(root,&sum)
	return sum
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
```
 时间复杂度O(n),空间复杂度O(1)