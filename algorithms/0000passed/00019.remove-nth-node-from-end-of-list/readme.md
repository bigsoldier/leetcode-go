#### 题目
<p>给定一个链表，删除链表的倒数第&nbsp;<em>n&nbsp;</em>个节点，并且返回链表的头结点。</p>

<p><strong>示例：</strong></p>

<pre>给定一个链表: <strong>1-&gt;2-&gt;3-&gt;4-&gt;5</strong>, 和 <strong><em>n</em> = 2</strong>.

当删除了倒数第二个节点后，链表变为 <strong>1-&gt;2-&gt;3-&gt;5</strong>.
</pre>

<p><strong>说明：</strong></p>

<p>给定的 <em>n</em>&nbsp;保证是有效的。</p>

<p><strong>进阶：</strong></p>

<p>你能尝试使用一趟扫描实现吗？</p>


 #### 题解
 1、两次遍历法
 获取链表的长度，然后找到需要删的节点
 ```go
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
 	var length int
 	var newList = new(ListNode)
 	newList.Next = head
 
 	first := head
 	// 遍历获取长度
 	for first != nil {
 		length ++
 		first = first.Next // 移动节点
 	}
 
 	first = newList
 	length -= n // 删除点前的长度
 	for length > 0 {
 		length--
 		first = first.Next // 移动节点
 	}
 	first.Next = first.Next.Next
 	return newList.Next
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-13/001901.png)
 时间复杂度O(n)，时间复杂度O(1)
 
 2、一次遍历法
 可以使用一个指针，第一个指针从列表开头移动 n+1 步，第二个指针从列表开头开始，知道第一个节点到达最后一个节点。
 ```go
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
 	newList := new(ListNode)
 	newList.Next = head
 	first,second := newList,newList
 	for i := 0; i <= n;i++{
 		first = first.Next
 	}
 	for first != nil {
 		first = first.Next
 		second = second.Next
 	}
 	second.Next = second.Next.Next
 	return newList.Next
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-13/001902.png)
 时间复杂度O(n)，时间复杂度O(1)