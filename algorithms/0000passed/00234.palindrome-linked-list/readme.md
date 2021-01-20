#### 题目
<p>请判断一个链表是否为回文链表。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> 1-&gt;2
<strong>输出:</strong> false</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> 1-&gt;2-&gt;2-&gt;1
<strong>输出:</strong> true
</pre>

<p><strong>进阶：</strong><br>
你能否用&nbsp;O(n) 时间复杂度和 O(1) 空间复杂度解决此题？</p>


 #### 题解
 1、转成数组
 ```go
func isPalindrome(head *ListNode) bool {
	lists := []int{}
	for ;head != nil; head = head.Next {
		lists = append(lists, head.Val)
	}
	for i := 0; i < len(lists)/2; i++ {
		if lists[i] != lists[len(lists)-i-1] {
			return false
		}
	}
	return true
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 2、递归
 先到达尾节点，然后利用递归的特性从后向前
 ```go
func isPalindrome(head *ListNode) bool {
	front := head
	var reverse func(node *ListNode) bool
	reverse = func(node *ListNode) bool {
		if node != nil {
			if !reverse(node.Next) {
				return false
			}
			if node.Val != front.Val {
				return false
			}
			front = front.Next
		}
		return true
	}
	return reverse(head)
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 3、反转后半链表
 ```go
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 找到前半部分链表的尾节点并反转后半部分的链表
	half := halfOfListNode(head)
	reverseList := reverse(half.Next)

	// 判断是否回文
	p1,p2 := head,reverseList
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	// 还原链表
	half.Next = reverse(reverseList)
	return result
}

func reverse(node *ListNode) *ListNode {
	var prev,curr *ListNode = nil,node
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}

func halfOfListNode(head *ListNode) *ListNode {
	slow,fast := head,head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
```
 时间复杂度O(n),空间复杂度O(1)