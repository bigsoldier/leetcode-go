#### 题目
<p>给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。</p>

<p>本题中，一个高度平衡二叉树是指一个二叉树<em>每个节点&nbsp;</em>的左右两个子树的高度差的绝对值不超过 1。</p>

<p><strong>示例:</strong></p>

<pre>给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
</pre>


 #### 题解
 分治
 
 根据108题，如果给个列表可以得到二叉平衡树，而链表的关键在于无法直接取到重点，
 所以这里用到快慢指针，取中间作为根节点
 ```go
func sortedListToBST(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

func getMid(left, right *ListNode) *ListNode {
	start,end := left,left
	for start!=right && start.Next!=right {
		start = start.Next.Next
		end = end.Next
	}
	return end
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getMid(left,right)
	root := &TreeNode{Val: mid.Val}
	root.Left = buildTree(left,mid)
	root.Right = buildTree(mid.Next,right)
	return root
}
```
 时间复杂度O(n),空间复杂度O(logn)
 
 优化
 
 上面的方法瓶颈在于寻找中位数，由于构造出来的二叉搜索树的中序遍历就是链表，所以将分治和中序遍历结合起来
 ```go
var global *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	global = head
	length := getLen(head)
	return buildTree(0,length-1)
}

func getLen(head *ListNode) int {
	var ret int
	for ; head != nil; head = head.Next {
		ret++
	}
	return ret
}

func buildTree(left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left+right+1)/2
	root := &TreeNode{}
	root.Left = buildTree(left,mid-1)
	root.Val = global.Val
	global = global.Next
	root.Right = buildTree(mid+1,right)
	return root
}
```