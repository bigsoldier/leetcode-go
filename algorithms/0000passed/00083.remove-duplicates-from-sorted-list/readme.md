#### 题目
<p>给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> 1-&gt;1-&gt;2
<strong>输出:</strong> 1-&gt;2
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> 1-&gt;1-&gt;2-&gt;3-&gt;3
<strong>输出:</strong> 1-&gt;2-&gt;3</pre>


 #### 题解
 移除重复的元素 p->next=p->next->next
 ```go
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	var newHead = head
	for newHead != nil && newHead.Next != nil {
		if newHead.Val == newHead.Next.Val {
			newHead.Next = newHead.Next.Next
		} else {
			newHead = newHead.Next
		}
	}
	return head
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-05-12/008301.png)
 时间复杂度O(n),空间复杂度O(1)