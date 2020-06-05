#### 题目
<p>反转从位置 <em>m</em> 到 <em>n</em> 的链表。请使用一趟扫描完成反转。</p>

<p><strong>说明:</strong><br>
1 &le;&nbsp;<em>m</em>&nbsp;&le;&nbsp;<em>n</em>&nbsp;&le; 链表长度。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 1-&gt;2-&gt;3-&gt;4-&gt;5-&gt;NULL, <em>m</em> = 2, <em>n</em> = 4
<strong>输出:</strong> 1-&gt;4-&gt;3-&gt;2-&gt;5-&gt;NULL</pre>


 #### 题解
 迭代链表反转
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-06-04/009202.png)
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-06-04/009203.png)
 ```go
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	// move the two points util they reach the proper starting point
	var cur = head
	var pre *ListNode
	for m > 1 {
		pre = cur
		cur = cur.Next
		m--
		n--
	}
	// tail 是反转后链表的尾部
	con,tail := pre,cur
	var third *ListNode
	for n > 0 {
		third = cur.Next
		cur.Next = pre
		pre = cur
		cur = third
		n--
	}

	if con != nil {
		con.Next = pre
	} else {
		head = pre
	}
	tail.Next = cur
	return head
}
```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-06-04/009201.png)
 时间复杂度O(n),空间复杂度O(n)