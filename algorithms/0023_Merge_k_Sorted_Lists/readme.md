#### 题目
合并k个排序链表，返回合并后的排序链表。请分析和描述算法复杂度。
**示例：**
```
输入：
[
    1->4->5,
    1->3->4,
    2->6
]
输出：1->1->2->3->4->4->5->6
```

#### 题解
1、暴力法
取出链表的所有元素，然后排序再组成新的链表
```go
func mergeKLists(lists []*ListNode) *ListNode {
	var vals []int
	for _, list := range lists {
		for list != nil {
			vals = append(vals, list.Val)
			list = list.Next
		}
	}
	sort.Ints(vals)

	var newList = new(ListNode)
	head := newList
	for _, val := range vals {
		newList.Next = &ListNode{Val:val,Next:nil}
		newList = newList.Next
	}
	return head.Next
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-18/002301.png)
时间复杂度O(nlogn)，空间复杂度O(n)

2、纵向比较
比较k个节点，获取最小值的节点，将选中的节点接在有序链表的后面。
时间复杂度O(kN)，空间复杂度O(n)

3、两两合并
将合并k个链表的问题转化成合并2个链表的k-1次。
时间复杂度O(kN)，空间复杂度O(1)

4、分而治之
```go
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists,0,len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := left + (right-left)/2
	l1 := merge(lists,left,mid)
	l2 := merge(lists,mid+1,right)
	return mergeList(l1,l2)
}

func mergeList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeList(l1.Next,l2)
		return l1
	} else {
		l2.Next = mergeList(l1,l2.Next)
		return l2
	}
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-18/002302.png)
时间复杂度O(nlogk)