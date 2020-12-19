#### 题目
<p>给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。</p>

<p><strong>说明:</strong>&nbsp;叶子节点是指没有子节点的节点。</p>

<p><strong>示例:</strong><br>
给定如下二叉树，以及目标和&nbsp;<code>sum = 22</code>，</p>

<pre>              <strong>5</strong>
             / \
            <strong>4</strong>   <strong>8</strong>
           /   / \
          <strong>11</strong>  13  <strong>4</strong>
         /  \    / \
        7    <strong>2</strong>  <strong>5</strong>   1
</pre>

<p>返回:</p>

<pre>[
   [5,4,11,2],
   [5,8,4,5]
]
</pre>


 #### 题解
 递归
 ```go
func pathSum(root *TreeNode, sum int) [][]int {
	var nums [][]int
	dfs(&nums,nil,root,sum)
	return nums
}

func dfs(nums *[][]int, num []int, root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum -= root.Val
    // 为防止指针问题
	var slice = make([]int,len(num))
	copy(slice,num)
	slice = append(slice, root.Val)
	if sum == 0 && root.Left == nil && root.Right == nil {
		*nums = append(*nums, slice)
		return
	}
	if root.Left != nil {
		dfs(nums, slice,root.Left,sum)
	}
	if root.Right != nil {
		dfs(nums,slice,root.Right,sum)
	}
}
```
 时间复杂度O(n^2^),空间复杂度O(n)