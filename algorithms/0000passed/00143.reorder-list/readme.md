#### 题目
<p>给定一个单链表&nbsp;<em>L</em>：<em>L</em><sub>0</sub>&rarr;<em>L</em><sub>1</sub>&rarr;&hellip;&rarr;<em>L</em><sub><em>n</em>-1</sub>&rarr;<em>L</em><sub>n ，</sub><br>
将其重新排列后变为： <em>L</em><sub>0</sub>&rarr;<em>L</em><sub><em>n</em></sub>&rarr;<em>L</em><sub>1</sub>&rarr;<em>L</em><sub><em>n</em>-1</sub>&rarr;<em>L</em><sub>2</sub>&rarr;<em>L</em><sub><em>n</em>-2</sub>&rarr;&hellip;</p>

<p>你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre>给定链表 1-&gt;2-&gt;3-&gt;4, 重新排列为 1-&gt;4-&gt;2-&gt;3.</pre>

<p><strong>示例 2:</strong></p>

<pre>给定链表 1-&gt;2-&gt;3-&gt;4-&gt;5, 重新排列为 1-&gt;5-&gt;2-&gt;4-&gt;3.</pre>


 #### 题解
 线性表
 ```go
func reorderList(head *ListNode) {
	if head == nil{
		return
	}
	var list = []*ListNode{}
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	i,j := 0,len(list)-1
	for i < j {
		list[i].Next = list[j]
		i++
		if i == j {
			break
		}
		list[j].Next = list[i]
		j--
	}
	list[i].Next = nil
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 寻找链表中点;后半段逆序;合并链表
 ```go
func reorderList(head *ListNode) {
	if head == nil{
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1,l2)
}

func middleNode(head *ListNode) *ListNode {
	slow,fast := head,head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev,cur *ListNode = nil,head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

func mergeList(l1, l2 *ListNode) {
	for l1 != nil && l2 != nil {
		tmp1,tmp2 := l1.Next,l2.Next
		l1.Next = l2
		l1 = tmp1
		l2.Next = l1
		l2 = tmp2
	}
}
```
 时间复杂度O(n),空间复杂度O(1)