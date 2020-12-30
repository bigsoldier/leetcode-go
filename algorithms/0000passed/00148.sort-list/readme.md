#### 题目
<p>在&nbsp;<em>O</em>(<em>n</em>&nbsp;log&nbsp;<em>n</em>) 时间复杂度和常数级空间复杂度下，对链表进行排序。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> 4-&gt;2-&gt;1-&gt;3
<strong>输出:</strong> 1-&gt;2-&gt;3-&gt;4
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> -1-&gt;5-&gt;3-&gt;4-&gt;0
<strong>输出:</strong> -1-&gt;0-&gt;3-&gt;4-&gt;5</pre>


 #### 题解
 时间复杂度为O(nlogn)的算法有堆排序，快速排序(最差O(n^2^)),归并排序.
 
 自顶向下归并排序
 ```go
func sortList(head *ListNode) *ListNode {
	return sort(head,nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 到末尾了
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow,fast := head,head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tmp,tmp1,tmp2 := dummyHead,head1,head2
	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val <= tmp2.Val {
			tmp.Next = tmp1
			tmp1 = tmp1.Next
		} else {
			tmp.Next = tmp2
			tmp2 = tmp2.Next
		}
		tmp = tmp.Next
	}
	if tmp1 != nil {
		tmp.Next = tmp1
	} else if tmp2 != nil {
		tmp.Next = tmp2
	}
	return dummyHead.Next
}
```
 时间复杂度O(nlogn),空间复杂度O(logn)
 
 自底向上归并排序
 ```go
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	length := 0
	// 计算链表长度
	for node := head; node != nil; node = node.Next {
		length++
	}
	dummyHead := &ListNode{Next: head}
	// 每两个子链表进行合并
	for subLength := 1; subLength <= length; subLength<<=1 {
		prev,curr := dummyHead,dummyHead.Next
		for curr != nil {
			head1 := curr
			for i := 1; i < subLength && curr.Next != nil; i++ {
				curr = curr.Next
			}
			head2 := curr.Next
			curr.Next = nil
			curr = head2
			for i := 1; i < subLength && curr != nil; i++ {
				curr = curr.Next
			}

			var next *ListNode
			if curr != nil {
				next = curr.Next
				curr.Next = nil
			}

			prev.Next = merge(head1,head2)
			for prev.Next != nil {
				prev= prev.Next
			}
			curr = next
		}
	}
	return dummyHead.Next
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tmp,tmp1,tmp2 := dummyHead,head1,head2
	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val <= tmp2.Val {
			tmp.Next = tmp1
			tmp1 = tmp1.Next
		} else {
			tmp.Next = tmp2
			tmp2 = tmp2.Next
		}
		tmp = tmp.Next
	}
	if tmp1 != nil {
		tmp.Next = tmp1
	} else if tmp2 != nil {
		tmp.Next = tmp2
	}
	return dummyHead.Next
}
```
 时间复杂度O(nlogn),空间复杂度O(1)