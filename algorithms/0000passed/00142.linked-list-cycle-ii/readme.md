#### 题目
<p>给定一个链表，返回链表开始入环的第一个节点。&nbsp;如果链表无环，则返回&nbsp;<code>null</code>。</p>

<p>为了表示给定链表中的环，我们使用整数 <code>pos</code> 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 <code>pos</code> 是 <code>-1</code>，则在该链表中没有环。</p>

<p><strong>说明：</strong>不允许修改给定的链表。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>head = [3,2,0,-4], pos = 1
<strong>输出：</strong>tail connects to node index 1
<strong>解释：</strong>链表中有一个环，其尾部连接到第二个节点。
</pre>

<p><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist.png" style="height: 97px; width: 300px;"></p>

<p><strong>示例&nbsp;2：</strong></p>

<pre><strong>输入：</strong>head = [1,2], pos = 0
<strong>输出：</strong>tail connects to node index 0
<strong>解释：</strong>链表中有一个环，其尾部连接到第一个节点。
</pre>

<p><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test2.png" style="height: 74px; width: 141px;"></p>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>head = [1], pos = -1
<strong>输出：</strong>no cycle
<strong>解释：</strong>链表中没有环。
</pre>

<p><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test3.png" style="height: 45px; width: 45px;"></p>

<p>&nbsp;</p>

<p><strong>进阶：</strong><br>
你是否可以不用额外空间解决此题？</p>


 #### 题解
 哈希表
 ```go
func detectCycle(head *ListNode) *ListNode {
	var nodeMap = map[*ListNode]bool{}
	for head != nil {
		if nodeMap[head] {
			return head
		}
		nodeMap[head] = true
		head = head.Next
	}
	return nil
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 快慢指针
 
 我们把链表分成3端，a是头到环处，b是环处到快慢指针相遇处，c是相遇处到入环处。
 那么我们可以得到慢指针走过a+b,
 快指针走过a+n(b+c)+b=a+(n+1)b+nc
 
 而快指针的速度是慢指针的2倍，所以我们得到等式:
 2(a+b)=a+(n+1)b+nc,简化为a=(n-1)b+nc=(n-1)(b+c)+c.
 
 也就是说当快慢指针相遇时，我们再额外用个指针从链表头开始，随后它和慢指针一起向后移动一个位置，最终会在入环处相遇。
 ```go
func detectCycle(head *ListNode) *ListNode {
	slow,fast := head,head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
```
 时间复杂度O(n),空间复杂度(1)