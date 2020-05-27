#### 题目
<p>给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中&nbsp;<em>没有重复出现&nbsp;</em>的数字。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> 1-&gt;2-&gt;3-&gt;3-&gt;4-&gt;4-&gt;5
<strong>输出:</strong> 1-&gt;2-&gt;5
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> 1-&gt;1-&gt;1-&gt;2-&gt;3
<strong>输出:</strong> 2-&gt;3</pre>


 #### 题解
 ```go
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var newHead = new(ListNode)
	newHead.Next = head
	p := newHead
	for p.Next != nil && p.Next.Next != nil {
		var tag bool

		// 移除重复元素
		for p.Next.Next != nil && p.Next.Val == p.Next.Next.Val {
			tag = true
			p.Next.Next = p.Next.Next.Next
		}

		// 移除最后的重复元素
		if tag {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return newHead.Next
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-05-12/008201.png)
 时间复杂度O(n),空间复杂度O(1)