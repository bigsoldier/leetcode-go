#### 题目
<p>给定一个二叉搜索树，编写一个函数&nbsp;<code>kthSmallest</code>&nbsp;来查找其中第&nbsp;<strong>k&nbsp;</strong>个最小的元素。</p>

<p><strong>说明：</strong><br>
你可以假设 k 总是有效的，1 &le; k &le; 二叉搜索树元素个数。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
&nbsp;  2
<strong>输出:</strong> 1</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
<strong>输出:</strong> 3</pre>

<p><strong>进阶：</strong><br>
如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化&nbsp;<code>kthSmallest</code>&nbsp;函数？</p>


 #### 题解
 1、中序遍历
 ```go
func kthSmallest(root *TreeNode, k int) int {
	var nums []int
	dfs(root,&nums)
	return nums[k-1]
}

func dfs(node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left,nums)
	*nums = append(*nums, node.Val)
	dfs(node.Right,nums)
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 2、优化递归
 ```go
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	left := getSize(root.Left)
	if left == k- 1 {
		return root.Val
	} else if left < k-1 {
		return kthSmallest(root.Right, k-left-1)
	} else {
		return kthSmallest(root.Left, k)
	}
}

func getSize(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return getSize(node.Left) + getSize(node.Right) + 1
}
```
 时间复杂度O(n),空间复杂度O(logn)
 
 3、迭代
 ```go
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return -1
}
```
 时间复杂度O(h+k),空间复杂度O(h+k)