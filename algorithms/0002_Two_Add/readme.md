#### 题目

给出两个 **非空** 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 **逆序** 的方式存储的，并且它们的每个节点只能存储 **一位** 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

**实例：**

```
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
```

#### 题解

![思路](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2019-12-31/000201.png)
主要考的是链表这种数据结构的理解，以及链表长度不等时怎么合并链表问题。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil || l2 == nil {
		return nil
	}

	head := &ListNode{Val:0,Next:nil}
	current := head
	carry := 0

	for l1 != nil || l2 != nil {
		var x,y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}

		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}

		current.Next = &ListNode{Val:(x+y+carry)%10,Next:nil}
		current = current.Next
		carry = (x+y+carry) / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		current.Next = &ListNode{Val:carry%10,Next:nil}
	}
	return head.Next
}
```