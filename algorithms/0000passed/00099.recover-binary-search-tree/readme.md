#### 题目
<p>二叉搜索树中的两个节点被错误地交换。</p>

<p>请在不改变其结构的情况下，恢复这棵树。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> [1,3,null,null,2]

&nbsp;  1
&nbsp; /
&nbsp;3
&nbsp; \
&nbsp;  2

<strong>输出:</strong> [3,1,null,null,2]

&nbsp;  3
&nbsp; /
&nbsp;1
&nbsp; \
&nbsp;  2
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [3,1,4,null,null,2]

  3
 / \
1   4
&nbsp;  /
&nbsp; 2

<strong>输出:</strong> [2,1,4,null,null,3]

  2
 / \
1   4
&nbsp;  /
 &nbsp;3</pre>

<p><strong>进阶:</strong></p>

<ul>
	<li>使用 O(<em>n</em>) 空间复杂度的解法很容易实现。</li>
	<li>你能想出一个只使用常数空间的解决方案吗？</li>
</ul>


 #### 题解
 1、中序遍历二叉树，然后取出不符合顺序的两个节点
 ```go
func recoverTree(root *TreeNode) {
	var nums []int
	inorder(root,&nums)
	swap := findTwoSwapped(nums)
	recoverT(root,2,swap[0],swap[1])
}

func inorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left,nums)
	*nums = append(*nums, root.Val)
	inorder(root.Right,nums)
}

func findTwoSwapped(nums []int) []int {
	var x,y = -1,-1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			y = nums[i+1]
			if x == -1 {
				x = nums[i]
			}
		}
	}
	return []int{x,y}
}

func recoverT(root *TreeNode, count, x, y int) {
	if root != nil {
		if root.Val == x || root.Val == y {
			if root.Val == x {
				root.Val = y
			} else {
				root.Val = x
			}
			if count == 0 {
				return
			}
			count--
		}
		recoverT(root.Left,count,x,y)
		recoverT(root.Right,count,x,y)
	}
}

```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-06-08/009901.png)
 时间复杂度O(n),空间复杂度O(n)
 
 2、迭代中序遍历
 ```go
func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode
	var dfs func(*TreeNode)
	// dfs 是中序遍历
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			dfs(root.Left)
		}

		if prev != nil && prev.Val > root.Val {
			// 如果要调整 [1, 2, 6, 4, 5, 3, 7] 中错位的 6 和 3
			// 其实是把 [6, 4] 中的较大值与 [5, 3] 中的较小值交换。这时，有两组错序。
			// 但是，还有可能是
			// [1, 3, 2] 中错位的 3 和 2，只有一组错序的 [3, 2]
			// 以下的两个 if 语句，可以兼容以上两种情况
			if first == nil {
				first = prev
			}
			if first != nil {
				// 当存在第二组错序的时候
				// second 的值，会被修改
				second = root
			}
		}

		prev = root

		if root.Right != nil {
			dfs(root.Right)
		}
	}

	dfs(root)

	// 题目已经保证存在被交换的节点了
	// 无需检查 first 和 second 是否为 nil
	first.Val, second.Val = second.Val, first.Val
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-06-08/009902.png)
 时间复杂度O(1),空间复杂度O(n)